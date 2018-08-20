package main

import (
	"database/sql"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
)

// RDBMSAccessor has all methods necessary for Note transactions and, as an interface, can easily be mocked.
// Any changes to the database implementation should implement this interface, and if the new struct will take over
// as the preferred database implementation, it should be assigned to 'db' in dependencies.go.
type RDBMSAccessor interface {
	Initialize(config *Config) (*sql.DB, error)
	AddNote(note *ehrpb.Note) (id int64, err error)
	UpdateNote(note *ehrpb.Note) error
	DeleteNote(id int64) error
	AllNotes() ([]*ehrpb.Note, error)
	GetNoteById(id int64) (*ehrpb.Note, error)
	FindNote(filter NoteFindFilter) ([]*ehrpb.Note, error)
	AddNoteTag(noteGuid string, tag string) (id int64, err error)
	GetNoteTagsByNoteGuid(noteGuid string) (tag []string, err error)
	AddNoteFragment(note *ehrpb.NoteFragment) (id int64, guid string, err error)
	UpdateNoteFragment(note *ehrpb.NoteFragment) error
	DeleteNoteFragment(noteFragmentGuid string) error
	AllNoteFragments() ([]*ehrpb.NoteFragment, error)
	GetNoteFragmentById(id int64) (*ehrpb.NoteFragment, error)
	GetNoteFragmentsByNoteGuid(noteGuid string) ([]*ehrpb.NoteFragment, error)
	FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error)
	AddNoteFragmentTag(noteGuid string, tag string) (id int64, err error)
	GetNoteFragmentTagsByNoteFragmentGuid(noteFragGuid string) (tag []string, err error)
	createSchema() error
}

// Find Note's with several fields to narrow search.
type NoteFindFilter struct {
	VisitGuid   string
	AuthorGuid  string
	PatientGuid string
	SearchTerms string
}

// Find NoteFragment's with several fields to narrow search.
type NoteFragmentFindFilter struct {
	NoteGuid    string
	VisitGuid   string
	AuthorGuid  string
	PatientGuid string
	SearchTerms string
}
