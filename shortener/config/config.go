package config

import (
	"flag"
	"github.com/spf13/viper"
)

var config Config

func Get() *Config {
	return &config
}

func init() {
	configPath := flag.String("config", "./config/shortener.local.yaml", "config path")
	flag.Parse()
	viper.SetConfigType("yaml")
	viper.SetConfigFile(*configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic("(viper.ReadInConfig) " + err.Error())
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic("(viper.Unmarshal) " + err.Error())
	}

}

type Config struct {
	Service  Service  `mapstructure:"service"`
	Http     HTTP     `mapstructure:"http"`
	Storage  string   `mapstructure:"storage"`
	Postgres Postgres `mapstructure:"postgres"`
	Memory   Memory   `mapstructure:"memory"`
}

type Service struct {
	Name              string `mapstructure:"service"`
	URLEncodingLength int    `mapstructure:"urlEncodingLength"`
}

type HTTP struct {
	Port                  int `mapstructure:"port"`
	RequestTimeoutSeconds int `mapstructure:"requestTimeoutSeconds"`
}

type Postgres struct {
	RequestTimeoutSeconds int    `mapstructure:"requestTimeoutSeconds"`
	Port                  int    `mapstructure:"port"`
	User                  string `mapstructure:"user"`
	Password              string `mapstructure:"password"`
	Database              string `mapstructure:"database"`
}

type Memory struct {
	RequestTimeoutSeconds int `mapstructure:"requestTimeoutSeconds"`
}
