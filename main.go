package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Initializing NoteClerk v0.1.0")

	s := &NoteClerkServer{}
	s.Context.Init()
	err := s.Initialize("tcp", "0.0.0.0", "50051")
	if err != nil {
		log.Fatal(err)
	}
}
