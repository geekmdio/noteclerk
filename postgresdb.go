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
	log.Fatal("Not implemented.")
	return nil, nil
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

func (d *DbPostgres) CreateSchema() error {

	createNoteTable := `create table note
(
	id serial not null
		constraint note_pkey
			primary key,
	date_created integer not null,
	note_guid varchar(38) not null,
	visit_guid varchar(38) not null,
	author_guid varchar(38) not null,
	patient_guid varchar(38) not null,
	type integer not null
)
;

alter table note owner to postgres
;

create unique index note_note_guid_uindex
	on note (note_guid)
;

create unique index note_visit_guid_uindex
	on note (visit_guid)
;

create unique index note_author_guid_uindex
	on note (author_guid)
;

create unique index note_patient_guid_uindex
	on note (patient_guid)
;

`

	createNoteFragmentTable := `create table note_fragment
(
	id serial not null
		constraint note_fragment_pkey
			primary key,
	date_created integer not null,
	note_fragment_guid varchar(38) not null,
	note_guid varchar(38) not null,
	issue_guid varchar(38) not null,
	icd_10code varchar(20),
	icd_10long varchar(255),
	description varchar(255),
	status integer not null,
	priority integer not null,
	topic integer not null,
	markdown_content varchar(2000) not null
)
;

alter table note_fragment owner to postgres
;

create unique index note_fragment_note_fragment_guid_uindex
	on note_fragment (note_fragment_guid)
;

create unique index note_fragment_note_guid_uindex
	on note_fragment (note_guid)
;

create unique index note_fragment_note_guid_uindex_2
	on note_fragment (note_guid)
;

create unique index note_fragment_issue_guid_uindex
	on note_fragment (issue_guid)
;
`

	createNoteTagTable := `create table note_tag
(
	id serial not null
		constraint note_tag_pkey
			primary key,
	note_guid varchar(38) not null,
	tag varchar(55) not null
)
;

alter table note_tag owner to postgres
;

`

	createNoteFragmentTagTable := `create table note_fragment_tags
(
	id serial not null
		constraint note_fragment_tags_pkey
			primary key,
	note_fragment_guid varchar(38) not null,
	tag varchar(55) not null
)
;

alter table note_fragment_tags owner to postgres
;

`

	log.Fatalf("%v\n%v\n%v\n%v\n", createNoteTable, createNoteTagTable, createNoteFragmentTable, createNoteFragmentTagTable)
	return nil
}