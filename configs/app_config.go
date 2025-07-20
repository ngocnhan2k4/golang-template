package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN string `mapstructure:"DSN"`
	Server struct{
		Port int `mapstructure:"Port"`
	}
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
