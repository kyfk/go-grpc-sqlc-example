package main

import (
	"database/sql"
	"errors"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/validator"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/config"
	grpc_auth "github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/grpc/interceptors/auth"
	grpc_zap "github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/grpc/interceptors/logging/zap"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/log"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/server"
	accountpb "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/account/v1"
	mepb "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/me/v1"
	todopb "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/todo/v1"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/server/account"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/server/health"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/server/me"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/server/todo"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/oklog/run"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	cfg := config.Get()

	lv, err := zapcore.ParseLevel(cfg.LogLevel)
	if err != nil {
		panic(err)
	}

	sc := log.ServiceContext{
		Service:   "todo",
		Version:   server.Version,
		GitCommit: server.GitCommit,
		BuildDate: server.BuildDate,
	}
	logger, err := log.NewLogger(lv, sc)
	if err != nil {
		panic(err)
	}

	_, err = sql.Open("mysql", cfg.MySQLDataSourceName)
	if err != nil {
		logger.Fatal("failed opening connection to mysql: %v", zap.Error(err))
	}

	grpcPanicRecoveryHandler := func(p any) (err error) {
		logger.Error("recovered from panic", zap.Any("paniced", p))
		return status.Errorf(codes.Internal, "%s", p)
	}

	pubKey, err := jwk.ParseKey([]byte(cfg.RawPublicKey))
	if err != nil {
		logger.Fatal("failed to parse public key", zap.Error(err))
	}

	privKey, err := jwk.ParseKey([]byte(cfg.RawPrivateKey))
	if err != nil {
		logger.Fatal("failed to parse private key", zap.Error(err))
	}
	_ = privKey

	accountServer, err := account.New()
	if err != nil {
		logger.Fatal("failed to initialize user server", zap.Error(err))
	}

	meServer, err := me.New()
	if err != nil {
		logger.Fatal("failed to initialize user server", zap.Error(err))
	}

	todoServer, err := todo.New()
	if err != nil {
		logger.Fatal("failed to initialize todo server", zap.Error(err))
	}

	healthServer := health.New()

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(grpc_zap.UnaryServerInterceptor(logger)),
			selector.UnaryServerInterceptor(auth.UnaryServerInterceptor(grpc_auth.Func(pubKey)), selector.MatchFunc(grpc_auth.MatchFunc)),
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandler(grpcPanicRecoveryHandler)),
			validator.UnaryServerInterceptor(),
		),
	)

	reflection.Register(grpcServer)

	accountpb.RegisterAccountServer(grpcServer, accountServer)
	mepb.RegisterMeServer(grpcServer, meServer)
	todopb.RegisterTodoServiceServer(grpcServer, todoServer)
	healthpb.RegisterHealthServer(grpcServer, healthServer)

	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		logger.Fatal("unexpected error", zap.Error(err))
	}

	var g run.Group

	{
		term := make(chan os.Signal, 1)
		signal.Notify(term, syscall.SIGTERM)
		g.Add(func() error {
			select {
			case sig := <-term:
				logger.Info("signal received", zap.String("signal", sig.String()))
				return errors.New(sig.String())
			}
		}, func(err error) {})
	}

	{
		logger.Debug("start grpc server at", zap.String("address", ":"+cfg.Port))
		g.Add(func() error {
			if err := grpcServer.Serve(lis); err != nil {
				logger.Error("failed to serve", zap.Error(err))
				return err
			}
			return nil
		}, func(error) {
			grpcServer.GracefulStop()
		})
	}

	if err := g.Run(); err != nil {
		os.Exit(1)
	}
}
