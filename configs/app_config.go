package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN    string `mapstructure:"DSN"`
	Server struct {
		Port int `mapstructure:"Port"`
	}
	AddressApi struct {
		BaseUrl      string `mapstructure:"BaseUrl"`
		CountriesUrl string `mapstructure:"CountriesUrl"`
	} `mapstructure:"AddressApi"`
}

var AppConfig Config

func LoadConfig() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("app")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.Unmarshal(&AppConfig)
}
