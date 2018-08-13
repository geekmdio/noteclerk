package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Config struct {
	ServerProtocol string
	ServerIp       string
	ServerPort     string
	DbIp           string
	DbPort         string
	DbUsername     string
	DbPassword     string
	DbName         string
	DbSslMode      string
}

type environment int

const (
	development environment = 1
	staging     environment = 2
	production  environment = 3
)

// An environmental JSON file
func LoadConfiguration() (c *Config, err error) {

	conf := &Config{}
	var path string
	switch env {
	case development:
		path = "config/config.development.json"
		break
	case staging:
		path = "config/config.staging.json"
		break
	case production:
		path = "config/config.production.json"
		break
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("could not read configuration file")
		return nil, err
	}

	unmarshalErr := json.Unmarshal(file, conf)
	if unmarshalErr != nil {
		log.Fatalf("failed to unmarshal the json file, recieved err %v", unmarshalErr)
	}

	return conf, nil
}
