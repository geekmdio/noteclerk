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

	note := `create table note
(
	id serial not null
		constraint note_pkey
			primary key,
	date_created_seconds integer not null,
	date_created_nanos integer default 0 not null,
	note_guid varchar(38) not null,
	visit_guid varchar(38) not null,
	author_guid varchar(38) not null,
	patient_guid varchar(38) not null,
	type integer not null
)
;

alter table note owner to postgres
;

create unique index note_id_uindex
	on note (id)
;

create unique index note_note_guid_uindex
	on note (note_guid)
;


`

	noteFragment := `create table note_fragment
(
	id serial not null
		constraint note_fragment_pkey
			primary key,
	date_created_seconds integer not null,
	date_created_nanos integer default 0 not null,
	note_fragment_guid varchar(38) not null,
	note_guid varchar(38) not null
		constraint note_fragment_note_note_guid_fk
			references note (note_guid),
	icd_10code varchar(15) not null,
	icd_10long varchar(250) not null,
	description varchar(150) not null,
	status integer not null,
	priority integer not null,
	topic integer not null,
	markdown_content varchar(2500) not null
)
;

alter table note_fragment owner to postgres
;

create unique index note_fragment_id_uindex
	on note_fragment (id)
;

create unique index note_fragment_note_fragment_guid_uindex
	on note_fragment (note_fragment_guid)
;

`

	noteTag := `create table note_tag
(
	id serial not null
		constraint note_tag_pkey
			primary key,
	note_guid varchar(38) not null
		constraint note_tag_note_note_guid_fk
			references note (note_guid),
	tag varchar(55) not null
)
;

alter table note_tag owner to postgres
;

create unique index note_tag_id_uindex
	on note_tag (id)
;


`

	noteFragmentTag := `create table note_fragment_tag
(
	id serial not null
		constraint note_fragment_tag_pkey
			primary key,
	note_fragment_guid varchar(38) not null
		constraint note_fragment_tag_note_fragment_note_fragment_guid_fk
			references note_fragment (note_fragment_guid),
	tag varchar(55) not null
)
;

alter table note_fragment_tag owner to postgres
;

create unique index note_fragment_tag_id_uindex
	on note_fragment_tag (id)
;



`

	d.createTable(note)
	d.createTable(noteFragment)
	d.createTable(noteTag)
	d.createTable(noteFragmentTag)


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