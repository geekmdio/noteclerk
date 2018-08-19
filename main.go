package main

import (
	"fmt"
	"strings"
	"github.com/sirupsen/logrus"
)

func main() {
	configPath := fmt.Sprintf("config/config.%v.json", strings.ToLower(NoteClerkEnv))
	log.Formatter = &logrus.JSONFormatter{}
	log.Infof("Loading configuration file from %v", configPath)
	config, err := LoadConfiguration(configPath)
	if err != nil {
		log.Panicf("Failed to load configuration file %v. Error returned: %v", configPath, err)
	}
	InitializeLogger(config.LogPath)
	if NoteClerkEnv == "" {
		log.Panicf("NOTECLERK_ENVIRONMENT not set.")
	}

	log.Infof("Initializing NoteClerk v%v on the %v environment.", config.Version, NoteClerkEnv)

	log.Infof("Starting GeekMD's NoteClerk Server on %v:%v.", config.ServerIp, config.ServerPort)
	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}
}
