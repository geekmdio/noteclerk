package main

import (
	"log"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/geekmdio/ehrprotorepo/goproto"
)

type DbAccessor interface {
	Init() (*sql.DB, error)
	AddNote(note *ehrpb.Note) (id int32, guid string, err error)
	UpdateNote(note *ehrpb.Note) error
	DeleteNote(id int32) error
	AllNotes() ([]*ehrpb.Note, error)
	GetNoteById(id int32) (*ehrpb.Note, error)
	FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error)
	AddNoteFragment(note *ehrpb.NoteFragment) (id int32, guid string, err error)
	UpdateNoteFragment(note *ehrpb.NoteFragment) error
	DeleteNoteFragment(id int32) error
	AllNoteFragments() ([]*ehrpb.NoteFragment, error)
	GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error)
	FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error)
}

type NotedContextPostgres struct {}

type NoteFindFilter struct {
	VisitGuid string
	AuthorGuid string
	PatientGuid string
	SearchTerms string
}

type NoteFragmentFindFilter struct {
	Day int32
	Month int32
	Year int32
	NoteGuid string
	VisitGuid string
	AuthorGuid string
	PatientGuid string
	Topic ehrpb.FragmentTopic
	Priority ehrpb.FragmentPriority
	Status ehrpb.NoteFragmentStatus
	Tags []*string
}

func (da *NotedContextPostgres) Init() (*sql.DB, error) {
	connStr := "user=postgres dbname=password sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return db, nil
}

func (da *NotedContextPostgres) AddNote(n *ehrpb.Note) (id int32, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *NotedContextPostgres) UpdateNote(n *ehrpb.Note) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContextPostgres) DeleteNote(id int32) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContextPostgres) AllNotes() ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContextPostgres) GetNoteById(id int32) (*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContextPostgres) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContextPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContextPostgres) AddNoteFragment(n *ehrpb.NoteFragment)  (id int32, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *NotedContextPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContextPostgres) DeleteNoteFragment(id int32)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContextPostgres) GetNoteFragmentsById(id int32) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContextPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}