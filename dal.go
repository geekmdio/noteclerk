package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/geekmdio/ehrprotorepo/goproto"
)

type NotedDal interface {
	Init() (*sql.DB, error)
	AddNote(note *ehrpb.Note) (id int, guid string, err error)
	UpdateNote(note *ehrpb.Note) error
	DeleteNote(id int) error
	AllNotes() ([]*ehrpb.Note, error)
	GetNoteById(id int) (*ehrpb.Note, error)
	FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error)
	AddNoteFragment(note *ehrpb.NoteFragment) (id int, guid string, err error)
	UpdateNoteFragment(note *ehrpb.NoteFragment) error
	DeleteNoteFragment(id int) error
	AllNoteFragments() ([]*ehrpb.NoteFragment, error)
	GetNoteFragmentsById(id int) (*ehrpb.NoteFragment, error)
	FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error)
}

type NotedContext struct {}

type NoteFindFilter struct {
	Day int32
	Month int32
	Year int32
	NoteGuid string
	VisitGuid string
	AuthorGuid string
	PatientGuid string
	Type ehrpb.NoteType
	Tags []*string
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

func (da *NotedContext) Init() (*sql.DB, error) {
	connStr := "user=pqgotest dbname=pqgotest sslmode=verify-full"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to open connection to database. Error: %v", err)
	}
	return db, nil
}

func (da *NotedContext) AddNote(n *ehrpb.Note) (id int, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *NotedContext) UpdateNote(n *ehrpb.Note) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContext) DeleteNote(id int) error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContext) AllNotes() ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContext) GetNoteById(id int) (*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContext) FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContext) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContext) AddNoteFragment(n *ehrpb.NoteFragment)  (id int, guid string, err error) {
	log.Fatal("Not implemented.")
	return 0, "",nil
}

func (da *NotedContext) UpdateNoteFragment(n *ehrpb.NoteFragment)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContext) DeleteNoteFragment(id int)  error {
	log.Fatal("Not implemented.")
	return nil
}

func (da *NotedContext) GetNoteFragmentsById(id int) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (da *NotedContext) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}