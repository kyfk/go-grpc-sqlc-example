package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/kyfk/go-grpc-sqlc-boilerplate/pkg/env"
)

type Config struct {
	once sync.Once

	Env                 env.Env `envconfig:"ENV" required:"true"`
	MySQLDataSourceName string  `envconfig:"MYSQL_DATA_SOURCE_NAME" required:"true"`
	Port                string  `envconfig:"PORT" required:"true"`
	LogLevel            string  `envconfig:"LOG_LEVEL" default:"debug"`
	RawPublicKey        string  `envconfig:"RAW_PUBLIC_KEY" required:"true"`
	RawPrivateKey       string  `envconfig:"RAW_PRIVATE_KEY" required:"true"`
}

var cfg = new(Config)

func Get() *Config {
	cfg.once.Do(func() {
		envconfig.MustProcess("", cfg)
	})
	return cfg
}
