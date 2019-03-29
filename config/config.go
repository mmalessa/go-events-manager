package config

import (
	// "fmt"
	// "net/url"
	// "os"
	"path/filepath"

	"gopkg.in/gcfg.v1"
)

type Config struct {
	RabbitMq struct {
		Host         string
		Username     string
		Password     string
		Port         string
		Vhost        string
	}
	Logs struct {
		Error      string
		Info       string
	}
}

func LoadAndParse(location string) (*Config, error) {
	if !filepath.IsAbs(location) {
		location, err := filepath.Abs(location)
		if err != nil {
			return nil, err
		}
		location = location
	}

	cfg := &Config{}
	if err := gcfg.ReadFileInto(cfg, location); err != nil {
		return nil, err
	}

	return cfg, nil
}