package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	User     string
	Password string
	Port     string
	Db       string
}

type APIConfig struct {
	Port string
}

var cfg *config

type config struct {
	API APIConfig
	DB  DbConfig
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.pass"),
		Db:       viper.GetString("database.name"),
	}

	return nil
}

func GetDB() DbConfig {
	return cfg.DB
}

func GetServerPort() string {
	port := fmt.Sprintf(":%s", cfg.API.Port)
	return port
}
