package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
		Charset  string
	}

	JWT struct {
		Secret string
	}
}

var GlobalConfig *Config

func InitConfig() {
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	GlobalConfig = &cfg
}
