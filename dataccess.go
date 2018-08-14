package main

import (
	"database/sql"
	"github.com/geekmdio/ehrprotorepo/goproto"
)

// RDBMSAccessor has all methods necessary for Note transactions and, as an interface, can easily be mocked.
// Any changes to the database implementation should implement this interface, and if the new struct will take over
// as the preferred database implementation, it should be assigned to 'db' in dependencies.go.
type RDBMSAccessor interface {
	Initialize(config *Config) (*sql.DB, error)
	AddNote(note *ehrpb.Note) (id int32, err error)
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
	CreateSchema() error
}

// Find Note's with several fields to narrow search.
type NoteFindFilter struct {
	VisitGuid string
	AuthorGuid string
	PatientGuid string
	SearchTerms string
}


// Find NoteFragment's with several fields to narrow search.
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