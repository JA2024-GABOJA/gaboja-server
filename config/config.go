package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}

	Db struct {
		User     string
		Pass     string
		Host     string
		Port     string
		Database string
	}

	WeatherApi struct {
		Url         string
		SideUrl     string
		AirPortCode string
		Key         string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
