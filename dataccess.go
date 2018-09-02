package main

import (
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"context"
)

// NoteClerkServer interface implements the gRPC server NoteServiceServer interface and adds an initialize feature.
// Any structures implementing this interface can be injected into the server global singleton variable in the
// dependencies.
type NoteClerkServer interface {
	CreateNote(context.Context, *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error)
	RetrieveNote(context.Context, *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error)
	UpdateNote(context.Context, *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error)
	DeleteNote(context.Context, *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error)
	SearchNotes(context.Context, *ehrpb.SearchNotesRequest) (*ehrpb.SearchNotesResponse, error)
	SearchNoteFragments(context.Context, *ehrpb.SearchNoteFragmentRequest) (*ehrpb.SearchNoteFragmentResponse, error)
	Initialize(config *Config, db RDBMSAccessor) error
}

// RDBMSAccessor has all methods necessary for Note transactions and, as an interface, can easily be mocked.
// Any changes to the database implementation should implement this interface, and if the new struct will take over
// as the preferred database implementation, it should be assigned to 'db' in dependencies.go.
type RDBMSAccessor interface {
	Initialize(config *Config) error
	AddNote(note *ehrpb.Note) (id int64, guid string, err error)
	UpdateNote(note *ehrpb.Note) error
	DeleteNote(guid string) error
	AllNotes() ([]*ehrpb.Note, error)
	GetNoteByGuid(guid string) (*ehrpb.Note, error)
	FindNotes(filter NoteFindFilter) ([]*ehrpb.Note, error)
	AddNoteTag(noteGuid string, tag string) (id int64, err error)
	GetNoteTagsByNoteGuid(noteGuid string) (tag []string, err error)
	AddNoteFragment(note *ehrpb.NoteFragment) (id int64, guid string, err error)
	UpdateNoteFragment(note *ehrpb.NoteFragment) error
	DeleteNoteFragment(noteFragmentGuid string) error
	AllNoteFragments() ([]*ehrpb.NoteFragment, error)
	GetNoteFragmentByGuid(guid string) (*ehrpb.NoteFragment, error)
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
