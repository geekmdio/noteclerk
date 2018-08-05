package main

import (
	"fmt"
	"github.com/geekmdio/ehrprotorepo/goproto"
	"google.golang.org/grpc"
	"net"
	"log"
	"context"
)

func main() {
	fmt.Println("Initializing NoteClerk v0.1.0")

	s := &NoteClerkServer{}

	s.Initialize("tcp", "0.0.0.0", "50051")

}


type NoteClerkServer struct {
	mockContext []ehrpb.Note
}


func (n *NoteClerkServer) NewNote(ctx context.Context, request *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	n.mockContext = append(n.mockContext, *request.Note)
	return nil, nil
}

func (n *NoteClerkServer) DeleteNote(ctx context.Context, request *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	return nil, nil
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

func (n *NoteClerkServer) Initialize(protocol string, ip string, port string) (s *grpc.Server, e error) {
	n.mockContext = make([]ehrpb.Note,0)

	connAddr := fmt.Sprintf("%v:%v", ip, port)
	lis, err := net.Listen(protocol, connAddr)
	if err != nil {
		log.Fatalf("Failed to listen on %v", connAddr)
		return nil, err
	}
	fmt.Printf("Listener successfully opened and listening on %v", connAddr)

	server := grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(server, n)

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		return nil, err
	}
	fmt.Println("Unary RPC server for NoteClerk initialized.")

	return server, nil
}

