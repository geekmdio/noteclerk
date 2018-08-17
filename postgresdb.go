package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
	"log"
)

type DbPostgres struct {
	db *sql.DB
}

// Initialize() initializes the connection to database. Ensure that the ./config/config.<environment>.json
// file has been created and properly configured with server and database values. Of createNoteTable, the '<environment>'
// can be set to any value, so long as the NOTECLERK_ENVIRONMENT environmental variable's value matches.
// RETURNS: *sql.db, error
func (d *DbPostgres) Initialize(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		config.DbUsername, config.DbPassword, config.DbIp, config.DbName, config.DbSslMode, config.DbPort)

	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err = db.Ping(); err != nil {
		err = errors.New("Could not ping PostgreSQL db")
		return nil, err
	}
	log.Printf("Successfully connected to PostgreSQL db at %v:%v", config.DbIp, config.DbPort)
	d.db = db
	d.CreateSchema()

	if err != nil {
		log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return d.db, nil
}

func (d *DbPostgres) AddNote(n *ehrpb.Note) (id int32, err error) {
	log.Fatal("Not implemented.")
	return 0,nil
}

func (d *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) DeleteNote(id int32) error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	rows, err := d.db.Query("SELECT * FROM createNoteTable;")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var notes []*ehrpb.Note
	for rows.Next() {
		tmpNote := NewNote()
		err := rows.Scan(&tmpNote.Id, &tmpNote.DateCreated.Seconds, &tmpNote.DateCreated.Nanos,
			&tmpNote.NoteGuid, &tmpNote.VisitGuid, &tmpNote.AuthorGuid,
			&tmpNote.PatientGuid, &tmpNote.Type)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		notes = append(notes, tmpNote)

	}
	return notes, nil
}

func (d *DbPostgres) GetNoteById(id int32) (*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AddNoteFragment(n *ehrpb.NoteFragment)  (id int32, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (d *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) DeleteNoteFragment(id int32)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

// https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func (d *DbPostgres) CreateSchema() error {
	d.createTable(createNoteTable)
	d.createTable(createNoteTagTable)
	d.createTable(createNoteFragmentTable)
	d.createTable(createNoteFragmentTagTable)

	notes, notesErr := d.AllNotes()
	if notesErr != nil {
		log.Println(notesErr)
	}
	fmt.Println(notes)

	return nil
}

func (d *DbPostgres) createTable(query string) {
	_, err := d.db.Exec(query)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Table created successfully.")
	}
}