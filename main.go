package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/geekmdtravis/noteclerk/gmdlog"
	"github.com/pkg/errors"
)

var log = &gmdlog.GmdLog{}

func main() {
	env := os.Getenv(Environment)
	path := fmt.Sprintf("config/config.%v.json", strings.ToLower(env))

	isProduction := strings.ToLower(env) == "production"
	log.InitializeLogger(isProduction)

	if env == "" {
		panic(fmt.Sprintf("Environmental variable %v not set.", Environment))
	}

	printHeading(fmt.Sprintf("Initializing NoteClerk v0.1.0 on the %v environment.", env))

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

	printSubtext(fmt.Sprintf("Loading configuration file from %v.", path))
	config, err := LoadConfiguration(path)
	err = errors.New("Test")
	if err != nil {
		log.Warnf("Error loading configuration file. Error: %v", err)
	}

	printSubtext(fmt.Sprintf("Starting GeekMD's NoteClerk Server on %v:%v.", config.ServerIp, config.ServerPort))
	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatalf("failed to initialize server")
	}
}
