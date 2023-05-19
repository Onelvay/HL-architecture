package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHTTPPort         = "8080"
	defaultHTTPReadTimeout  = 15 * time.Second
	defaultHTTPWriteTimeout = 15 * time.Second
)

type (
	Config struct {
		Http     HttpConfig
		Postgres PostgresConfig
	}

	PostgresConfig struct {
		Host     string
		Port     string
		User     string
		DbName   string
		Password string
	}

	HttpConfig struct {
		Port         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
)

func New() (cfg Config, err error) {
	postgresConfig := PostgresConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		DbName:   viper.GetString("db.name"),
		Password: viper.GetString("db.password"),
	}

	httpConfig := HttpConfig{
		Port:         defaultHTTPPort,
		ReadTimeout:  defaultHTTPReadTimeout,
		WriteTimeout: defaultHTTPWriteTimeout,
	}

	return Config{
		Http:     httpConfig,
		Postgres: postgresConfig,
	}, nil
}

func ParseYaml() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	return viper.ReadInConfig()
}
