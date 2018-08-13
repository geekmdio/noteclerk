package main

import (
	"fmt"
)

func main() {

	fmt.Println("Initializing NoteClerk v0.1.0")

	s := &NoteClerkServer{}
	err := s.Initialize("tcp", "0.0.0.0", "50051", pdi.DB)
	if err != nil {
		pdi.Log.Fatal(err)
	}
}
