package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Executables struct {
		BaseDirectory string
		Commands map [string]string
	}
	RabbitMq struct {
		Connection struct {
			Host         string
			Username     string
			Password     string
			Port         string
			Vhost        string	
		}
		Consumers map [string]string
	}
}

func loadConfig(location string) (*Config, error){
	
	yamlFile, err := ioutil.ReadFile(location)
	if err != nil {
        fmt.Printf("Load config ERROR  #%v ", err)
	}

	cfg := &Config{}
	err = yaml.Unmarshal(yamlFile, cfg)
    if err != nil {
        fmt.Printf("Unmarshal ERROR: %v", err)
    }

    return cfg, err
}