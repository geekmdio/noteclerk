package main

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"fmt"
	"os"
	"strings"
)

// This struct is the model for a JSON configuration file that should be located in
// ./config/config.<environment>.json, where '.' indicates the server root, and where
// <environment> can be any lowercase value so long as the NOTECLERK_ENVIRONMENT environmental variable matches.
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

// Load the configuration JSON and return the Config struct.
func LoadConfiguration() (c *Config, err error) {
	path := fmt.Sprintf("config/config.%v.json", strings.ToLower(os.Getenv("NOTECLERK_ENVIRONMENT")))

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

// Invoked on the commandline as `./noteclerk config`, this prints a copy of the bare-bones JSON configuration files
// to the terminal. This should be copied and pasted into ./config/config.<environment>.json, where '.' is the server's
// root directory, and '<environment>' can be any lowercase value so long as the NOTECLERK_ENVIRONMENT
// environmental variable matches.
func MakeConfig() {
	fmt.Println("Put the following into ./config/config.<environment>.json, where '<environment>' can be any " +
					"lowercase value so long as the NOTECLERK_ENVIRONMENT environmental variable matches:")
	fmt.Println(`{
	"ServerProtocol": "",
	"ServerIp": "",
	"ServerPort": "",
	"DbIp": "",
	"DbPort": "",
	"DbUsername": "",
	"DbPassword": "",
	"DbName": "",
	"DbSslMode": ""
}`)
}