package main

import (
	"fmt"
	"strings"
)

func main() {

	validateEnv()

	config := loadConfig()

	initLogger(config)

	initStatement(config)

	s := &NoteClerkServer{}
	if err := s.Initialize(config, db); err != nil {
		log.Fatal(err)
	}
}

func initStatement(config *Config) {
	initStatement := fmt.Sprintf("NoteClerk v%v is launching in %v", config.Version, strings.ToUpper(NoteClerkEnv))
	fmt.Println(initStatement)
	log.Infof(initStatement)

	serverStartStatement := fmt.Sprintf("Starting GeekMD's NoteClerk Server on %v:%v.", config.ServerIp, config.ServerPort)
	fmt.Println(serverStartStatement)
	log.Infof(serverStartStatement)
}

func loadConfig() *Config {
	fmt.Printf("Loading configuration file from %v", configPath)
	config, err := LoadConfiguration(configPath)
	if err != nil {
		log.Panic(err)
	}
	return config
}

func initLogger(config *Config) {
	if err := InitializeLogger(config.LogPath); err != nil {
		log.Fatalf("Unable to load the configuration file %v", config.LogPath)
	}
}

func validateEnv() {
	if NoteClerkEnv == "" {
		log.Panic("NOTECLERK_ENVIRONMENT environmental variable must be set.")
	}
}
