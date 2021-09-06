package config

import (
	"time"

	"github.com/spf13/viper"
)

var defaultValues = struct {
	httpPort               string
	httpTimeoutRead        time.Duration
	httpTimeoutWrite       time.Duration
	httpMaxHeaderMegabytes int
}{
	httpPort:               "8000",
	httpTimeoutRead:        10 * time.Second,
	httpTimeoutWrite:       10 * time.Second,
	httpMaxHeaderMegabytes: 1,
}

type (
	HttpConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}

	PostgresConfig struct {
		Username     string `mapstructure:"username"`
		Host         string `mapstructure:"host"`
		Port         string `mapstructure:"port"`
		DatabaseName string `mapstructure:"dbname"`
		Password     string `mapstructure:"password"`
		SslMode      string `mapstructure:"sslmode"`
	}

	Config struct {
		Http     HttpConfig
		Postgres PostgresConfig
	}
)

func Init(configPath string) (*Config, error) {
	setDefaults()
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
	if err := viper.UnmarshalKey("http", &cfg.Http); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return err
	}

	return nil
}

func setDefaults() {
	viper.SetDefault("http.port", defaultValues.httpPort)
	viper.SetDefault("http.maxHeaderBytes", defaultValues.httpMaxHeaderMegabytes)
	viper.SetDefault("http.readTimeout", defaultValues.httpTimeoutRead)
	viper.SetDefault("http.writeTimeout", defaultValues.httpTimeoutWrite)
}
