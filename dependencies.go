package main

import (
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"time"
)

var db = &DbPostgres{}

var mockDb = &MockDb{}

func newNote() *ehrpb.Note {
	return &ehrpb.Note{
		Id:          0,
		DateCreated: timestampNow(),
		NoteGuid:    uuid.New().String(),
		Fragments:   make([]*ehrpb.NoteFragment, 0),
		Tags:        make([]string,0),
	}
}

func newNoteFragment() *ehrpb.NoteFragment {
	return &ehrpb.NoteFragment{
		Id:               0,
		DateCreated:      timestampNow(),
		NoteFragmentGuid: uuid.New().String(),
		IssueGuid:        "IssueGuid not set",
		Icd_10Code:       "Icd_10Code not set",
		Icd_10Long:       "Icd_10Long not set",
		Description:      "Description not set",
		Status:           ehrpb.NoteFragmentStatus_INCOMPLETE,
		Priority:         ehrpb.FragmentPriority_NO_PRIORITY,
		Topic:            ehrpb.FragmentTopic_NO_TOPIC,
		MarkdownContent:  "MarkdownContent not set",
		Tags:             make([]string,0),
	}
}

func timestampNow() *timestamp.Timestamp {
	now := time.Now()
	ts := &timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.UnixNano()),
	}
	return ts
}
