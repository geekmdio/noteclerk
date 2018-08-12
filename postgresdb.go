package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
)

type DbPostgres struct {}

// Init() initializes the connection to database. Importantly, the
// PostgreSQL user, password, host, dbname, sslmode, and port should
// all be environmental variables in the OS that match the below naming schemes.
// RETURNS: *sql.DB, error
func (da *DbPostgres) Init() (*sql.DB, error) {
	user := os.Getenv("NOTECLERK_POSTGRES_USER")
	password := os.Getenv("NOTECLERK_POSTGRES_PASSWORD")
	dbName := os.Getenv("NOTECLERK_POSTGRES_DBNAME")
	sslMode := os.Getenv("NOTECLERK_POSTGRES_SSLMODE")
	port := os.Getenv("NOTECLERK_POSTGRES_PORT")
	host := os.Getenv("NOTECLERK_POSTGRES_HOST")

	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		user, password, host, dbName, sslMode, port)

	db, err := sql.Open("postgres", connStr)

	if err = db.Ping(); err != nil {
		err = errors.New("Could not ping PostgreSQL DB")
		return nil, err
	}
	pdi.Log.Printf("Successfully connected to PostgreSQL DB at %v:%v", host, port)

	if err != nil {
		pdi.Log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return db, nil
}

func (da *DbPostgres) AddNote(n *ehrpb.Note) (id int32, guid string, err error) {
	pdi.Log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	pdi.Log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) DeleteNote(id int32) error {
	pdi.Log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) GetNoteById(id int32) (*ehrpb.Note, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) AddNoteFragment(n *ehrpb.NoteFragment)  (id int32, guid string, err error) {
	pdi.Log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	pdi.Log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) DeleteNoteFragment(id int32)  error {
	pdi.Log.Fatal("Not implemented.")
	return nil
}

func (da *DbPostgres) GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}

func (da *DbPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	pdi.Log.Fatal("Not implemented.")
	return nil, nil
}