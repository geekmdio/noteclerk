package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
	"log"
)

type DbPostgres struct {}

// Init() initializes the connection to database. Ensure that the ./config/config.<environment>.json
// file has been created and properly configured with server and database values. Of note, the '<environment>'
// can be set to any value, so long as the NOTECLERK_ENVIRONMENT environmental variable's value matches.
// RETURNS: *sql.db, error
func (da *DbPostgres) Init(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		config.DbUsername, config.DbPassword, config.DbIp, config.DbName, config.DbSslMode, config.DbPort)

	db, err := sql.Open("postgres", connStr)

	if err = db.Ping(); err != nil {
		err = errors.New("Could not ping PostgreSQL db")
		return nil, err
	}
	log.Printf("Successfully connected to PostgreSQL db at %v:%v", config.DbIp, config.DbPort)

	if err != nil {
		log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return db, nil
}

func (da *DbPostgres) AddNote(n *ehrpb.Note) (id int32, err error) {
	log.Fatal("Not implemented.")
	return 0,nil
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