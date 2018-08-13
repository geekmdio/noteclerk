package main

import (
	"fmt"
	"log"
)

const env = development

func main() {

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
