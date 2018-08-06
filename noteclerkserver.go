package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
)

type NoteClerkServer struct {
	mockContext []ehrpb.Note
}


func (n *NoteClerkServer) NewNote(ctx context.Context, request *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	n.mockContext = append(n.mockContext, *request.Note)
	res := &ehrpb.CreateNoteResponse{
		Status:               &ehrpb.NoteServiceResponseStatus {
			HttpCode:             204,
			Message:              "Created",
		},
		Note:                 request.Note,
	}
	return res, nil
}

func (n *NoteClerkServer) DeleteNote(ctx context.Context, request *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	var newContext []ehrpb.Note
	for i := range n.mockContext {
		if n.mockContext[i].Id == request.Id {
			newContext = append(newContext, n.mockContext[:i]...)
			if len(n.mockContext) > i {
				newContext = append(newContext, n.mockContext[i+1:]...)
			}
		}
		n.mockContext = newContext
	}
	res := ehrpb.DeleteNoteResponse{
		Status:               &ehrpb.NoteServiceResponseStatus{
			HttpCode:             200,
			Message:              "Success, deleted.",
		},
	}
	return &res, nil
}

func (n *NoteClerkServer) RetrieveNote(ctx context.Context, request *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	return nil, nil
}

func (n *NoteClerkServer) FindNote(ctx context.Context, request *ehrpb.FindNoteRequest) (*ehrpb.FindNoteResponse, error) {
	return nil, nil
}

func (n *NoteClerkServer) UpdateNote(context.Context, *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {
	return nil, nil
}

func (n *NoteClerkServer) Initialize(protocol string, ip string, port string) (ok bool, e error) {
	n.mockContext = make([]ehrpb.Note,0)

	connAddr := fmt.Sprintf("%v:%v", ip, port)
	lis, err := net.Listen(protocol, connAddr)
	if err != nil {
		return false, errors.Errorf("Failed to listen on %v. Caught error: %v", connAddr, err)
	}

	server := grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(server, n)

	if err = server.Serve(lis); err != nil {
		return false, errors.Errorf("Failed to serve: %v", err)
	}

	return true, nil
}

