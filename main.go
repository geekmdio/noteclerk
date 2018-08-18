package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	path := fmt.Sprintf("config/config.%v.json", strings.ToLower(NoteClerkEnv))
	log.Infof("Loading configuration file from %v", path)
	config, err := LoadConfiguration(path)
	if err != nil {
		log.Panicf("Failed to load configuration file %v. Error returned: %v", path, err)
	}
	log.InitializeLogger(config.LogPath)
	if NoteClerkEnv == "" {
		log.Panicf("NOTECLERK_ENVIRONMENT not set.")
	}

	log.Infof("Initializing NoteClerk v%v on the %v environment.", config.Version, NoteClerkEnv)

	if len(os.Args) >= 2 {
		if strings.ToLower(os.Args[1]) == "config" {
			fmt.Printf("Add to './config/config.<environment>.json':\n%v\n", configJson)
			return
		} else if strings.ToLower(os.Args[1]) == "help" {
			fmt.Println("You can execute the server as ./prog, or generate a config file with ./prog config")
			return
		} else {
			fmt.Println("You've provided an incorrect argument. Run ./prog help for assistance.")
			return
		}
	}

	log.Infof("Starting GeekMD's NoteClerk Server on %v:%v.", config.ServerIp, config.ServerPort)
	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatalf("failed to initialize server")
	}
}
