package config

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Postgres PgConfig
}

type PgConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	PgDriver string
}

func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()
	fmt.Printf("Loading config file: %s\n", filename)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
