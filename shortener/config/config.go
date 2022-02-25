package config

import (
	"flag"
	"github.com/spf13/viper"
)

var config Config

func Get() *Config {
	return &config
}

var configPath = flag.String("config", "./config/shortener.local.yaml", "config path")

// Init вместо init, потому что с последним слишком трудно тестить.
// Легче один раз вызвать в main
func Init() {
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
}

type Service struct {
	Name                 string `mapstructure:"name"`
	MinURLEncodingLength int    `mapstructure:"minUrlEncodingLength"`
	ShortURLDomainName   string `mapstructure:"shortUrlDomainName"`
}

type HTTP struct {
	Port                  int `mapstructure:"port"`
	RequestTimeoutSeconds int `mapstructure:"requestTimeoutSeconds"`
}

type Postgres struct {
	URL string `mapstructure:"url"`
}
