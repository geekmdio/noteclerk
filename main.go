package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	args := os.Args
	if len(args) >= 2 {
		if strings.ToLower(args[1]) == "config" {
			MakeConfig()
			return
		}
	}

	fmt.Println("Initializing NoteClerk v0.1.0")

	config, err := LoadConfiguration()
	if err != nil {
		log.Fatalf("unable to load Config file. err: %v", err)
	}

	s := &NoteClerkServer{}
	err = s.Initialize(config, db)
	if err != nil {
		log.Fatalf("failed to initialize server")
	}
}
