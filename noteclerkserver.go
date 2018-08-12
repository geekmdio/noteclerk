package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/geekmdio/ehrprotorepo/goproto"
	"github.com/pkg/errors"
	"context"
	"log"
)

type NoteClerkServer struct {
	db       DbAccessor
	ip       string
	port     string
	protocol string
	connAddr string
	server   *grpc.Server
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

	n.db = dependencies.DB
	_, err := n.db.Init()
	if err != nil {
		log.Fatalf("%v", err)
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


