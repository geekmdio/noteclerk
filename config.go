package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"os"
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

// An environmental JSON file
func LoadConfiguration() (c *Config, err error) {
	path := getEnvironmentDependentPath()

	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic("could not read configuration file")
		return nil, err
	}

	conf := &Config{}
	unmarshalErr := json.Unmarshal(file, conf)
	if unmarshalErr != nil {
		log.Fatalf("failed to unmarshal the json file, recieved err %v", unmarshalErr)
	}

	return conf, nil
}

func MakeConfig() {
	testToWrite :=
`{
	"ServerProtocol": "",
	"ServerIp": "",
	"ServerPort": "",
	"DbIp": "",
	"DbPort": "",
	"DbUsername": "",
	"DbPassword": "",
	"DbName": "",
	"DbSslMode": ""
}`
	fmt.Println("Put the following into ./config/config.<environment>.json, where <environment> is development, staging, or production:")
	fmt.Println(testToWrite)
}


func getEnvironmentDependentPath() string {
	env := os.Getenv("NOTECLERK_ENVIRONMENT")
	var path string
	switch env {
	case "development":
		path = "config/config.development.json"
		break
	case "staging":
		path = "config/config.staging.json"
		break
	case "production":
		path = "config/config.production.json"
		break
	}
	return path
}