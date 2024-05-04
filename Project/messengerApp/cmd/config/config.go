package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func LoadConfig() (*DatabaseConfig, *ServerConfig, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yaml")

	var dbConfig DatabaseConfig
	var serverConfig ServerConfig

	if err := viper.ReadInConfig(); err != nil {
		return nil, nil, fmt.Errorf("fatal error loading config file: %s", err)
	}

	if err := viper.UnmarshalKey("database", &dbConfig); err != nil {
		return nil, nil, fmt.Errorf("fatal error unmarshalling database config: %s", err)
	}

	if err := viper.UnmarshalKey("server", &serverConfig); err != nil {
		return nil, nil, fmt.Errorf("fatal error unmarshalling server config: %s", err)
	}

	return &dbConfig, &serverConfig, nil
}
