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

type User struct {
	Host string
	Port string
	DB   DB
}

type Auth struct {
	Host string
	Port string
}

type Gateway struct {
	Port string
}

type Config struct {
	JWT     JWT
	User    User
	Auth    Auth
	Gateway Gateway
}

type JWT struct {
	Secret string
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
