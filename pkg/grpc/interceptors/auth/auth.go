package auth

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	pkgjwt "github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/jwt"
	accountv1 "github.com/kyfk/go-grpc-sqlc-boilerplate/protogen/proto/account/v1"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

func Func(pubKey jwk.Key) auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		auth, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		token, err := jwt.Parse(
			[]byte(auth),
			jwt.WithKey(jwa.RS256, pubKey),
			jwt.WithValidate(false),
		)
		if err != nil {
			return nil, err
		}

		claims := token.PrivateClaims()

		userID := claims[pkgjwt.UserIDClaimKey]

		ctx = context.WithValue(ctx, pkgjwt.UserIDCtxKey, userID)

		return ctx, nil
	}
}

func MatchFunc(ctx context.Context, callMeta interceptors.CallMeta) bool {
	switch {
	case healthpb.Health_ServiceDesc.ServiceName == callMeta.Service,
		accountv1.Account_ServiceDesc.ServiceName == callMeta.Service:
		return false
	}
	return true
}
