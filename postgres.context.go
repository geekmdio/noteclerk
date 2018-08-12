package main

import (
	"log"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/geekmdio/ehrprotorepo/goproto"
)

type DbPostgres struct {}

func (da *DbPostgres) Init() (*sql.DB, error) {
	connStr := "user=postgres dbname=password sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return db, nil
}

func (da *DbPostgres) AddNote(n *ehrpb.Note) (id int32, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) DeleteNote(id int32) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) GetNoteById(id int32) (*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) AddNoteFragment(n *ehrpb.NoteFragment)  (id int32, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) DeleteNoteFragment(id int32)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}