package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
	"context"
	"time"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
)

type NoteClerkServer struct {
	db       DbAccessor
	ip       string
	port     string
	protocol string
	connAddr string
	server   *grpc.Server
}

func (n *NoteClerkServer) NewNote(ctx context.Context, nr *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	n.verifyServerInitialized()

	noteToAdd := nr.Note
	noteToAdd.NoteGuid = uuid.New().String()
	noteToAdd.DateCreated = timestampNow()

	id, err := n.db.AddNote(noteToAdd)
	cnr := &ehrpb.CreateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{},
	}
	if err != nil {
		cnr.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		cnr.Status.Message = fmt.Sprintf("Could not add note. Error: %v", err)
		cnr.Note = nil
		return nil, err
	}

	cnr.Note = noteToAdd
	cnr.Note.Id = id
	cnr.Status.HttpCode = ehrpb.StatusCodes_OK
	cnr.Status.Message = "Successfully submit new note."

	return cnr, nil
}

func (n *NoteClerkServer) DeleteNote(ctx context.Context, dnr *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	n.verifyServerInitialized()

	dnRes := &ehrpb.DeleteNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              "Successfully deleted note.",
		},
	}

	err := n.db.DeleteNote(dnr.Id)
	if err != nil {
		dnRes.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		dnRes.Status.Message = "Failed to delete note."
		return dnRes, errors.Errorf("%v. Error: %v", dnRes.Status.Message, err)
	}

	return dnRes, nil
}

func (n *NoteClerkServer) RetrieveNote(ctx context.Context, rnr *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	n.verifyServerInitialized()

	note, err := n.db.GetNoteById(rnr.Id)
	retNotRes := &ehrpb.RetrieveNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              "Successfully retrieved note.",
		},
		Note: note,
	}
	if err != nil {
		retNotRes.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		retNotRes.Status.Message = "unable to locate note"
		return nil, fmt.Errorf("%v, error: %v", retNotRes.Status.Message, err)
	}

	return retNotRes, nil
}

func (n *NoteClerkServer) FindNote(ctx context.Context, fnr *ehrpb.FindNoteRequest) (*ehrpb.FindNoteResponse, error) {
	n.verifyServerInitialized()

	filter := NoteFindFilter{
		VisitGuid:   fnr.VisitGuid,
		AuthorGuid:  fnr.AuthorGuid,
		PatientGuid: fnr.PatientGuid,
		SearchTerms: fnr.SearchTerms,
	}

	findNoteResponse := &ehrpb.FindNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              "found one or more notes matching query",
		},
		Note:                 nil,
	}
	notes, err := n.db.FindNote(filter)
	if err != nil {
		findNoteResponse.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		findNoteResponse.Status.Message = "unable to locate notes matching that query"
		return findNoteResponse, fmt.Errorf("%v, error: %v", findNoteResponse.Status.Message, err)
	}

	findNoteResponse.Note = notes
	return findNoteResponse, nil
}

func (n *NoteClerkServer) UpdateNote(ctx context.Context, unr *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {
	n.verifyServerInitialized()

	updateNoteResponse := &ehrpb.UpdateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              "note successfully updated",
		},
	}

	if unr.Id != unr.Note.Id {
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_CONFLICT
		updateNoteResponse.Status.Message = "the id provided for the update note request does not match the id of the note"
		return updateNoteResponse, fmt.Errorf("%v", updateNoteResponse.Status.Message)
	}


	err := n.db.UpdateNote(unr.Note)
	if err != nil {
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		updateNoteResponse.Status.Message = "unable to update note"
		return updateNoteResponse, fmt.Errorf("%v, error: %v", updateNoteResponse.Status.Message, err)
	}

	return updateNoteResponse, nil
}

func (n *NoteClerkServer) Initialize(protocol string, ip string, port string, db DbAccessor) error {
	// Build up the server's fields
	n.constructor(protocol, ip, port, db)

	// Initialize server database
	_, err := n.db.Init()
	if err != nil {
		panic("Failed to initialize database.")
	}

	// Create and register GRPC server
	n.server = grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(n.server, n)


	// Create listener
	lis, err := net.Listen(n.getProtocol(), n.getConnectionAddr())
	if err != nil {
		return errors.Errorf("Failed to listen on %v.", n.connAddr)
	}

	// Serve
	if err = n.server.Serve(lis); err != nil {
		return errors.Errorf("Failed to serve on the listener.")
	}

	return nil
}

func (n *NoteClerkServer) constructor(protocol string, ip string, port string, db DbAccessor) {
	n.ip = ip
	n.port = port
	n.protocol = protocol
	n.connAddr = fmt.Sprintf("%v:%v", n.getIp(), n.getPort())
	n.db = db
}

func (n *NoteClerkServer) getIp() string {
	return n.ip
}

func (n *NoteClerkServer) getPort() string {
	return n.port
}

func (n *NoteClerkServer) getProtocol() string {
	return n.protocol
}

func (n *NoteClerkServer) getConnectionAddr() string {
	return n.connAddr
}

func timestampNow() *timestamp.Timestamp {
	now := time.Now()
	ts := &timestamp.Timestamp{
		Seconds: now.Unix(),
		Nanos:   int32(now.UnixNano()),
	}
	return ts
}
func (n *NoteClerkServer) verifyServerInitialized() {
	if n.db == nil {
		panic("NoteClerkServer's database was not initialized.")
	}
}
