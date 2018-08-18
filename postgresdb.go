package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
)

type DbPostgres struct {
	db *sql.DB
}

// Initialize() initializes the connection to database. Ensure that the ./config/config.<environment>.json
// file has been created and properly configured with server and database values. Of note, the '<environment>'
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

func (d *DbPostgres) AddNote(n *ehrpb.Note) (id int64, err error) {

	scanErr := d.db.QueryRow(addNoteQuery, n.DateCreated.GetSeconds(), n.DateCreated.GetNanos(),
		n.GetNoteGuid(), n.GetVisitGuid(), n.GetAuthorGuid(), n.GetPatientGuid(), n.GetType(),
		n.GetStatus()).Scan(n.Id)

	if scanErr != nil {
		return 0, scanErr
	}

	for _, v := range n.GetFragments() {
		_, _, err := d.AddNoteFragment(v)
		if err != nil {
			return 0, scanErr
		}
	}

	return n.Id,nil
}

func (d *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) DeleteNote(id int64) error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	rows, err := d.db.Query("SELECT * FROM note;")
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
			&tmpNote.PatientGuid, &tmpNote.Type, &tmpNote.Status)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		notes = append(notes, tmpNote)

	}
	return notes, nil
}

func (d *DbPostgres) GetNoteById(id int64) (*ehrpb.Note, error) {
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

func (d *DbPostgres) AddNoteFragment(n *ehrpb.NoteFragment)  (id int64, guid string, err error) {
	scanErr := d.db.QueryRow(addNoteFragmentQuery, n.DateCreated.Seconds, n.DateCreated.Nanos,
		n.GetNoteFragmentGuid(), n.GetNoteGuid(), n.GetIcd_10Code(), n.GetIcd_10Long(),
		n.GetDescription(), n.GetStatus(), n.GetPriority(), n.GetTopic(), n.GetContent()).Scan(n.Id)

	if scanErr != nil {
		return 0, "", scanErr
	}
	return n.GetId(), n.GetNoteFragmentGuid(),nil
}

func (d *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) DeleteNoteFragment(id int64)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (d *DbPostgres) GetNoteFragmentsById(id int64) (*ehrpb.NoteFragment, error) {
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

	tmpNote := NewNote()
	tmpFrag := NewNoteFragment()
	tmpFrag.NoteGuid = tmpNote.GetNoteGuid()
	tmpFrag2 := NewNoteFragment()
	tmpFrag2.NoteGuid = tmpNote.GetNoteGuid()
	tmpNote.Fragments = append(tmpNote.Fragments, tmpFrag, tmpFrag2)

	d.AddNote(tmpNote)
	//TODO: remove this
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