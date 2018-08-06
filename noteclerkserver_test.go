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

	valid := res.Note.GetId() == 1 && req.Note.GetId() == 1 &&
		res.Note.GetDateCreated().Nanos == req.Note.GetDateCreated().Nanos &&
		res.Status.HttpCode == ehrpb.StatusCodes_OK

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

	if res.Status.GetHttpCode() != ehrpb.StatusCodes_OK {
		t.Fatalf("Expected status code %v, got %v", ehrpb.StatusCodes_OK, res.Status.GetHttpCode())
	}
}

func TestNoteClerkServer_RetrieveNote(t *testing.T) {
	ctx := context.Background()
	s := &NoteClerkServer{}
	s.NewNote(ctx, genNoteRequest())

	retrieveReq := &ehrpb.RetrieveNoteRequest{
		Id:                   1,
	}
	res, err := s.RetrieveNote(ctx, retrieveReq)

	if err != nil {
		t.Fatal(err)
	}

	if res.Status.GetHttpCode() != ehrpb.StatusCodes_OK {
		t.Fatalf("Expected status code %v, got %v", ehrpb.StatusCodes_OK, res.Status.GetHttpCode())
	}
}

func TestNoteClerkServer_UpdateNote(t *testing.T) {
	ctx := context.Background()
	s := &NoteClerkServer{}
	s.NewNote(ctx, genNoteRequest())
	initialType := s.mockContext[0].Type

	updateReq := &ehrpb.UpdateNoteRequest{
		Id: 1,
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
			Type:        ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION,
		},
	}
	res, err := s.UpdateNote(ctx, updateReq)
	updatedType := s.mockContext[0].Type

	if err != nil {
		t.Fatal(err)
	}

	if res.Status.GetHttpCode() != ehrpb.StatusCodes_OK {
		t.Fatalf("Expected status code %v, got %v", ehrpb.StatusCodes_OK, res.Status.GetHttpCode())
	}

	ok := initialType != updatedType && updatedType == ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION
	if !ok {
		t.Fatalf("Expected type %v and actual type %v.", ehrpb.NoteType_CONTINUED_CARE_DOCUMENTATION, updatedType)
	}

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