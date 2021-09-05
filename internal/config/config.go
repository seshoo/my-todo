package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	HTTPConfig struct {
		Host         string        `mapstructure:"host"`
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"readTimeout"`
		WriteTimeout time.Duration `mapstructure:"writeTimeout"`
	}

	PostgresConfig struct {
		Username     string `mapstructure:"username"`
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		DatabaseName string `mapstructure:"dbname"`
		SslMode      string `mapstructure:"sslmode"`
	}

	Config struct {
		HTTPConfig
		PostgresConfig
	}
)

func Init(configPath string) (*Config, error) {
	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigFile(directory string) error {
	viper.AddConfigPath(directory)
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTPConfig); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.PostgresConfig); err != nil {
		return err
	}

	return nil
}
