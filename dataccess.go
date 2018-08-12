package main

import (
	"database/sql"
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