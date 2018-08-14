package main

import (
	"io/ioutil"
	"encoding/json"
)

// This is the environmental variable in the OS that should be se to your preferred
// environment and matched to a configuration file. E.g. NOTECLERK_ENVIRONMENT=production
// will match to './config/config.production.json'.
const Environment = "NOTECLERK_ENVIRONMENT"

// Template for config files. The config files should be stored in the root directory
// under a folder 'config', with the naming convention 'config.<environment>.json'
// where environment is set by the os environmental variable NOTECLERK_ENVIRONMENT
// to be any string. The string the environmental variable is set to should match
// '<environment'. E.g. NOTECLERK_ENVIRONMENT=production will match to
// './config/config.production.json'

const configJson = `{
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
func LoadConfiguration(path string) (c *Config, err error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	conf := &Config{}
	unmarshalErr := json.Unmarshal(file, conf)
	if unmarshalErr != nil {
		return &Config{}, unmarshalErr
	}

	return conf, nil
}
