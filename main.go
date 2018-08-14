package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var env = os.Getenv(Environment)
	var path = fmt.Sprintf("config/config.%v.json", strings.ToLower(env))

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
			os.Exit(1)
		}
	}

	printSubtext("Loading configuration file...")
	config, err := LoadConfiguration(path)
	if err != nil {
		log.Fatalf("unable to load Config file. err: %v", err)
	}

	printSubtext("Starting GeekMD's NoteClerk Server...")
	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatalf("failed to initialize server")
	}
}
