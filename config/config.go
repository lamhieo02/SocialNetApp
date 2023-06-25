package config

import "github.com/spf13/viper"

type Config struct {
	Mysql MysqlConfig
 	Redis RedisConfig
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
	v.SetConfigFile("config.yml")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}