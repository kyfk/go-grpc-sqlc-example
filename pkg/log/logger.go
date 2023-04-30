package log

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level zapcore.Level, sc ServiceContext) (*zap.Logger, error) {
	logger, err := newConfig(level).Build(zap.Fields(zap.Object("serviceContext", sc)))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return logger.Named(sc.Service), nil
}

func newConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

type ServiceContext struct {
	Service   string
	Version   string
	GitCommit string
	BuildDate string
}

func (sc ServiceContext) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if sc.Service == "" {
		return errors.New("service name is mandatory")
	}
	enc.AddString("service", sc.Service)
	enc.AddString("version", sc.Version)
	enc.AddString("gitcommit", sc.GitCommit)
	enc.AddString("builddate", sc.BuildDate)
	return nil
}
