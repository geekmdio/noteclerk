package main

import (
	"testing"
	"context"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/golang/protobuf/ptypes/timestamp"
	)

func TestNoteClerkServer_NewNote(t *testing.T) {
	s := &NoteClerkServer{}
	req := genNoteRequest()
	res, err := s.NewNote(context.Background(), req)
	if err != nil {
		t.Fatalf("Failed to create a new note, with error: %v", err)
	}

	valid := res.Note.GetId() == 1 && req.Note.GetId() == 1 && res.Note.GetDateCreated().Nanos == req.Note.GetDateCreated().Nanos
	if !valid {
		t.Fatalf("The returned note is invalid.")
	}
}

func TestNoteClerkServer_DeleteNote(t *testing.T) {
	s := &NoteClerkServer{}
	s.NewNote(context.Background(), genNoteRequest())
	beforeLen := len(s.mockContext)

	req := ehrpb.DeleteNoteRequest{
		Id:                   1,
	}

	_, err := s.DeleteNote(context.Background(), &req)
	if err != nil {
		t.Fatalf("Failed to delete note, with error: %v", err)
	}
	afterLen := len(s.mockContext)

	ok := beforeLen == 1 && afterLen == 0

	if !ok {
		t.Fatalf("Failed to delete the item in the mock context")
	}
}

func TestNoteClerkServer_FindNote(t *testing.T) {
	ctx := context.Background()
	s := &NoteClerkServer{}
	s.NewNote(ctx, genNoteRequest())

	findReq := ehrpb.FindNoteRequest{
		SearchTerms:          "diabetes",
		AuthorGuid:           "",
		PatientGuid:          "",
		VisitGuid:            "",
	}
	res, err := s.FindNote(ctx, &findReq)

	if err != nil {
		t.Fatal(err)
	}

	if res.Status.HttpCode != 300 {
		t.Fatalf("Expected status code 300, got %v", res.Status.HttpCode)
	}
}

func TestNoteClerkServer_Initialize(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestNoteClerkServer_RetrieveNote(t *testing.T) {
	t.Fatal("Not implemented")
}

func TestNoteClerkServer_UpdateNote(t *testing.T) {
	t.Fatal("Not implemented")
}

func genNoteRequest() *ehrpb.CreateNoteRequest {
	return &ehrpb.CreateNoteRequest{
		Note: &ehrpb.Note{
			Id: 1,
			DateCreated: &timestamp.Timestamp{
				Seconds: 1354231,
				Nanos:   324234,
			},
			NoteGuid:    "000000000",
			VisitGuid:   "000000000",
			AuthorGuid:  "000000000",
			PatientGuid: "000000000",
			Type:        ehrpb.NoteType_FOLLOW_UP,
			Fragments:   nil,
		},
	}
}