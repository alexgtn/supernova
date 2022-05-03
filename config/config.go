package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
}

func FromLocalFile(path string, filename string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(filename)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%v", err)
	}

	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		log.Fatalf("unable to decode into Config struct, %v", err)
	}

	return conf
}
