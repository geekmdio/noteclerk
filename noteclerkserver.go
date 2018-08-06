package main

import (
	"context"
	"fmt"
		"google.golang.org/grpc"

	"github.com/geekmdio/ehrprotorepo/goproto"
	"net"
	"github.com/pkg/errors"
)

type NoteClerkServer struct {
	mockContext []ehrpb.Note
	server *grpc.Server
}

func (n *NoteClerkServer) Initialize(protocol string, ip string, port string) error {
	n.mockContext = make([]ehrpb.Note,0)

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

func (n *NoteClerkServer) NewNote(ctx context.Context, request *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	n.mockContext = append(n.mockContext, *request.Note)
	res := &ehrpb.CreateNoteResponse{
		Status:               &ehrpb.NoteServiceResponseStatus {
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              ehrpb.StatusCodes_OK.String(),
		},
		Note:                 request.Note,
	}
	return res, nil
}

func (n *NoteClerkServer) DeleteNote(ctx context.Context, request *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	var newMockContext []ehrpb.Note
	var statusCode ehrpb.StatusCodes
	for i := range n.mockContext {
		if n.mockContext[i].Id == request.Id {
			newMockContext = append(newMockContext, n.mockContext[:i]...)
			if len(n.mockContext) > i {
				newMockContext = append(newMockContext, n.mockContext[i+1:]...)
			}
			statusCode = ehrpb.StatusCodes_OK
		}
		n.mockContext = newMockContext
	}
	res := ehrpb.DeleteNoteResponse{
		Status:               &ehrpb.NoteServiceResponseStatus{
			HttpCode:             statusCode,
			Message:              statusCode.String(),
		},
	}
	return &res, nil
}

func (n *NoteClerkServer) RetrieveNote(ctx context.Context, request *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	var retNote *ehrpb.Note = nil
	var statusCode ehrpb.StatusCodes
	for _, note := range n.mockContext {
		if note.Id == request.Id {
			retNote = &note
			statusCode = ehrpb.StatusCodes_OK
		}
	}
	res := &ehrpb.RetrieveNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             statusCode,
			Message:              statusCode.String(),
		},
		Note: retNote,
	}
	return res, nil
}

func (n *NoteClerkServer) FindNote(ctx context.Context, request *ehrpb.FindNoteRequest) (*ehrpb.FindNoteResponse, error) {
	var found []*ehrpb.Note
	for _, note := range n.mockContext {
		if note.AuthorGuid == request.AuthorGuid ||
			note.VisitGuid == request.VisitGuid ||
			note.PatientGuid == request.PatientGuid {
			found = append(found, &note)
		}
	}
	res := &ehrpb.FindNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              ehrpb.StatusCodes_OK.String(),
		},
		Note:                 found,
	}
	return res, nil
}

func (n *NoteClerkServer) UpdateNote(ctx context.Context, request *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {
	for _, note := range n.mockContext {
		if note.Id == request.Note.Id {
			note = *request.Note
		}
	}

	res := &ehrpb.UpdateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode:             ehrpb.StatusCodes_OK,
			Message:              ehrpb.StatusCodes_OK.String(),
		},
	}
	return res, nil
}

