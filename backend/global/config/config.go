package config

import (
	"time"

	"github.com/spf13/viper"
)

type DB struct {
	User    string
	Passwd  string
	DSN     string
	Name    string
	Options string
}

type Account struct {
	Port string
	DB   DB
}

type Gateway struct {
	Port string
}

type Config struct {
	JWT     JWT
	Account Account
	Gateway Gateway
}

type JWT struct {
	Secret string
	Issuer string
	Expire time.Duration
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	conf := &Config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}
	conf.JWT.Expire = conf.JWT.Expire * time.Minute
	return conf, nil
}
