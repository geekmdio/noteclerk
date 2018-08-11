package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
	"context"
)

type NoteClerkServer struct {
	Context DbAccessor
	server  *grpc.Server
}

func (n *NoteClerkServer) NewNote(context.Context, *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	panic("implement me")
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
	n.Context = &NotedContextPostgres{}

	connAddr := fmt.Sprintf("%v:%v", ip, port)

	lis, err := net.Listen(protocol, connAddr)
	if err != nil {
		return errors.Errorf("Failed to listen on %v. Caught error: %v", connAddr, err)
	}

	n.server = grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(n.server, n)

	if err = n.server.Serve(lis); err != nil {
		return errors.Errorf("Failed to serve: %v", err)
	}

	return nil
}

func (n *NoteClerkServer) Stop() {
	n.server.GracefulStop()
	n.server.Stop()
}
