package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	printHeading("Initializing NoteClerk v0.1.0")
	// Check program arguments
	args := os.Args
	if len(args) >= 2 {
		if strings.ToLower(args[1]) == "config" {
			MakeConfig()
			return
		} else if strings.ToLower(args[1]) == "help" {
			fmt.Println("You can execute the server as ./prog, or generate a config file with ./prog config")
			return
		} else {
			fmt.Println("You've provided an incorrect argument. Run ./prog help for assistance.")
			os.Exit(1)
		}

	}

	printSubtext("Loading configuration file...")
	config, err := LoadConfiguration()
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
