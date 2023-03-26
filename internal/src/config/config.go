package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Database DB  `mapstructure:"db"`
	API      API `mapstructure:"api"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Dbname   string `mapstructure:"db_name"`
}

type API struct {
	FixerKey        string `mapstructure:"fixer_api_key"`
	ExchangeRateKey string `mapstructure:"exchange_rate_api_key"`
}

func InitConfig() Config {
	viper.SetConfigType("toml")
	workingdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigFile(workingdir + "/internal/src/config/config.toml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
