package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Mysql MysqlConfig 
 	Redis RedisConfig
	AuthenAndPostService AuthenAndPostServiceConfig `mapstructure:"authen_and_post_service"`
}

type AuthenAndPostServiceConfig struct {
	Port int
}

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Database int
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile("./config/config.yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}