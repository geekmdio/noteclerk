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

	noteToAdd := nr.Note
	noteToAdd.NoteGuid = uuid.New().String()
	noteToAdd.DateCreated = timestampNow()

	// this is for a test commit
	id, err := n.db.AddNote(noteToAdd)
	cnr := &ehrpb.CreateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{},
	}
	if err != nil {
		pdi.Log.Fatalf("Failed to create new note. Error: %v", err)
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



func (n *NoteClerkServer) DeleteNote(context.Context, *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	panic("implement me")
}

func (n *NoteClerkServer) RetrieveNote(context.Context, *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	panic("implement me")
}

func (n *NoteClerkServer) FindNote(context.Context, *ehrpb.FindNoteRequest) (*ehrpb.FindNoteResponse, error) {
	panic("implement me")
}

func (n *NoteClerkServer) UpdateNote(context.Context, *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {
	panic("implement me")
}

func (n *NoteClerkServer) Initialize(protocol string, ip string, port string) error {
	n.constructor(protocol, ip, port)

	ehrpb.RegisterNoteServiceServer(n.server, n)

	lis, err := net.Listen(n.getProtocol(), n.getConnectionAddr())
	if err != nil {
		return errors.Errorf("Failed to listen on %v.", n.connAddr)
	}

	if err = n.server.Serve(lis); err != nil {
		return errors.Errorf("Failed to serve on the listener.")
	}

	return nil
}

func (n *NoteClerkServer) constructor(protocol string, ip string, port string) {
	n.ip = ip
	n.port = port
	n.protocol = protocol
	n.connAddr = fmt.Sprintf("%v:%v", n.getIp(), n.getPort())

	n.db = pdi.DB
	_, err := n.db.Init()
	if err != nil {
		pdi.Log.Fatalf("%v", err)
	}

	n.server = grpc.NewServer()
}


func (n *NoteClerkServer) getDb() DbAccessor {
	return n.db
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
