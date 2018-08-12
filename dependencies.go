package main

import (
	"log"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"time"
)

// Pure dependency injection.
type Dependencies struct {
	DB  DbAccessor
	MockDB DbAccessor
	Log log.Logger
	Note *ehrpb.Note
	NoteFragment *ehrpb.NoteFragment
	Timestamp *timestamp.Timestamp
}

// Pure dependency injection vector.
var pdi = Dependencies {
	DB:  &DbPostgres{},
	MockDB: &MockDb{},
	Log: log.Logger{},
	Timestamp:TimestampNow(),
	Note: NewNoteEmpty(),
	NoteFragment: NoteFragmentEmpty(),
}

func NewNoteEmpty() *ehrpb.Note {
	return &ehrpb.Note{
		Id:          0,
		DateCreated: TimestampNow(),
		NoteGuid:    uuid.New().String(),
		Fragments:   make([]*ehrpb.NoteFragment, 0),
		Tags:        make([]string,0),
	}
}

func NewNote(visitGuid string, authorGuid string, patientGuid string, noteType ehrpb.NoteType) *ehrpb.Note {
	note := NewNoteEmpty()
	note.VisitGuid = visitGuid
	note.AuthorGuid = authorGuid
	note.PatientGuid = patientGuid
	note.Type = noteType
	note.NoteGuid = uuid.New().String()
	return note
}

func NoteFragmentEmpty() *ehrpb.NoteFragment {
	return &ehrpb.NoteFragment{
		Id: 0,
		DateCreated: TimestampNow(),
		NoteFragmentGuid:     uuid.New().String(),
		IssueGuid:            "IssueGuid not set",
		Icd_10Code:           "Icd_10Code not set",
		Icd_10Long:           "Icd_10Long not set",
		Description:          "Description not set",
		Status:               ehrpb.NoteFragmentStatus_INCOMPLETE,
		Priority:             ehrpb.FragmentPriority_NO_PRIORITY,
		Topic:                ehrpb.FragmentTopic_NO_TOPIC,
		MarkdownContent:      "MarkdownContent not set",
		Tags:                 make([]string,0),
	}
}

func NoteFragment(noteGuid string, issueGuid string) *ehrpb.NoteFragment {
	noteFragment := NoteFragmentEmpty()
	noteFragment.NoteGuid = noteGuid
	noteFragment.IssueGuid = issueGuid
	return noteFragment
}

func TimestampNow() *timestamp.Timestamp {
	now := time.Now()
	ts := &timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.UnixNano()),
	}
	return ts
}
