package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DB  Mongo
	MQ  Rabbit
	Ctx struct {
		Ttl time.Duration `mapstructure:"ttl"`
	} `mapstructure:"ctx"`
}

type Mongo struct {
	URI        string
	Username   string
	Password   string
	Database   string
	Collection string
}

type Rabbit struct {
	Username string
	Password string
	Host     string
	Port     int
}

func New(folder, filename string) (*Config, error) {
	cfg := new(Config)

	viper.AddConfigPath(folder)
	viper.SetConfigName(filename)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", &cfg.DB); err != nil {
		return nil, err
	}

	if err := envconfig.Process("rabbit", &cfg.MQ); err != nil {
		return nil, err
	}

	return cfg, nil
}
