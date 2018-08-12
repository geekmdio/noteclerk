package main

import (
	"testing"
	"context"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/google/uuid"
)

func TestNoteClerkServer_NewNote(t *testing.T) {
	s := &NoteClerkServer{}
	c := context.Background()
	cnr := &ehrpb.CreateNoteRequest{}
	res, err := s.NewNote(c, cnr)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if res.Status.HttpCode != ehrpb.StatusCodes_OK {
		t.Fatalf("Status could should be OK")
	}

	_, parseErr := uuid.Parse(res.Note.GetNoteGuid())
	if parseErr != nil {
		t.Fatalf("The note did not have a new GUID created for it, and is likely invalid.")
	}
}

func TestNoteClerkServer_DeleteNote(t *testing.T) {
	panic("implement me")
}

func TestNoteClerkServer_RetrieveNote(t *testing.T) {
	panic("implement me")
}

func TestNoteClerkServer_FindNote(t *testing.T) {
	panic("implement me")
}

func TestNoteClerkServer_UpdateNote(t *testing.T) {
	panic("implement me")
}

func TestNoteClerkServer_Initialize(t *testing.T) {
	panic("implement me")
}