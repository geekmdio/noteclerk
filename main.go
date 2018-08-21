package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

func main() {
	log.Formatter = &logrus.JSONFormatter{}
	log.Infof("Loading configuration file from %v", configPath)
	config, err := LoadConfiguration(configPath)
	if err != nil {
		log.Panic(err)
	}
	InitializeLogger(config.LogPath)
	if NoteClerkEnv == "" {
		log.Panic(ErrMainEnvironmentalVariableNotSet)
	}

	initStatement := fmt.Sprintf("NoteClerk v%v is launching in %v", config.Version, strings.ToUpper(NoteClerkEnv))
	fmt.Println(initStatement)
	log.Infof(initStatement)

	serverStartStatement := fmt.Sprintf("Starting GeekMD's NoteClerk Server on %v:%v.", config.ServerIp, config.ServerPort)
	fmt.Println(serverStartStatement)
	log.Infof(serverStartStatement)

	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatal(err)
	}
}
