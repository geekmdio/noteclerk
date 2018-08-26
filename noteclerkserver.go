package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/geekmdio/noted"
)

type NoteClerkServer struct {
	db       RDBMSAccessor
	ip       string
	port     string
	protocol string
	connAddr string
	server   *grpc.Server
}

func (n *NoteClerkServer) CreateNote(ctx context.Context, nr *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	cnr := &ehrpb.CreateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{},
	}

	noteToAdd := nr.Note
	noteToAdd.NoteGuid = uuid.New().String()
	noteToAdd.DateCreated = TimestampNow()

	id, addErr := n.db.AddNote(noteToAdd)
	if addErr != nil {
		newErr := errors.Wrapf(ErrNewNoteNotAddedToDb, "%v", addErr)
		log.Warn(newErr)
		cnr.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		cnr.Status.Message = "Failed to insert new note into database."
		cnr.Note = nil
		return nil, newErr
	}

	cnr.Note = noteToAdd
	cnr.Note.Id = id
	cnr.Status.HttpCode = ehrpb.StatusCodes_OK
	cnr.Status.Message = "Successfully submit new note."

	return cnr, nil
}

func (n *NoteClerkServer) DeleteNote(ctx context.Context, dnr *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	dnRes := &ehrpb.DeleteNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully marked note as deleted in the database.",
		},
	}

	deleteErr := n.db.DeleteNote(dnr.Id)
	if deleteErr != nil {
		newErr := errors.Wrapf(ErrDeleteNoteFailsUpdateStatus, "%v", deleteErr)
		log.Warn(newErr)
		dnRes.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		dnRes.Status.Message = "Failed to delete note from the database."
		return dnRes, newErr
	}

	return dnRes, nil
}

func (n *NoteClerkServer) RetrieveNote(ctx context.Context, rnr *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	retNoteRes := &ehrpb.RetrieveNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully retrieved note from database.",
		},
	}

	note, err := n.db.GetNoteById(rnr.Id)
	if err != nil {
		newErr := errors.Wrapf(ErrRetrieveNoteFailsRetrieveFromDb, "%v", err)
		log.Warn(newErr)
		retNoteRes.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		retNoteRes.Status.Message = "Failed to retrieve note from database."
		return retNoteRes, newErr
	}

	err = noted.OrganizeNoteFragments(note)
	if err != nil {
		log.Warn("Could not organize the note fragments by fragment priority.")
	}
	retNoteRes.Note = note

	return retNoteRes, nil
}

func (n *NoteClerkServer) SearchNotes(ctx context.Context, fnr *ehrpb.SearchNotesRequest) (*ehrpb.SearchNotesResponse, error) {
	filter := NoteFindFilter{
		VisitGuid:   fnr.VisitGuid,
		AuthorGuid:  fnr.AuthorGuid,
		PatientGuid: fnr.PatientGuid,
		SearchTerms: fnr.SearchTerms,
	}

	findNoteResponse := &ehrpb.SearchNotesResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully found one or more notes matching query.",
		},
	}

	for _, v := range findNoteResponse.Notes {
		err := noted.OrganizeNoteFragments(v)
		if err != nil {
			log.Warn("Could not organize the note fragments by fragment priority.")
		}
	}

	notes, err := n.db.FindNotes(filter)
	if err != nil {
		newErr := errors.Wrapf(ErrFindNoteFailedToFindInDb, "%v", err)
		log.Warn(newErr)
		findNoteResponse.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		findNoteResponse.Status.Message = "Failed to locate notes matching query"
		return findNoteResponse, newErr
	}

	findNoteResponse.Notes = notes
	return findNoteResponse, nil
}

func (n *NoteClerkServer) UpdateNote(ctx context.Context, unr *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {

	updateNoteResponse := &ehrpb.UpdateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully updated note.",
		},
	}

	if unr.Id != unr.Note.Id {
		newErr := ErrUpdateNoteFailedFromIdMismatch
		log.Warn(newErr)
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_CONFLICT
		updateNoteResponse.Status.Message = "Failed to update note. The id provided for the update note request does not match the id of the note."
		return updateNoteResponse, newErr
	}

	err := n.db.UpdateNote(unr.Note)
	if err != nil {
		newErr := errors.Wrapf(ErrUpdateNoteFailedToUpdateInDb, "%v", err)
		log.Warn(newErr)
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		updateNoteResponse.Status.Message = "UpdateNote failed. Unable to update note in the database."
		return updateNoteResponse, newErr
	}

	return updateNoteResponse, nil
}

func (n *NoteClerkServer) SearchNoteFragments(ctx context.Context, snf *ehrpb.SearchNoteFragmentRequest) (*ehrpb.SearchNoteFragmentResponse, error) {
	panic("implement me")
}

func (n *NoteClerkServer) Initialize(config *Config, db RDBMSAccessor) error {
	// Build up the server's fields
	conErr := n.constructor(config, db)
	if conErr != nil {
		return conErr
	}

	// Initialize server database
	err := n.db.Initialize(config)
	if err != nil {
		return errors.Wrapf(ErrDbInitFails, "%v", err)
	}
	log.Info("Successfully connected to database.")

	// Create and register gRPC server
	n.server = grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(n.server, n)
	log.Info("Assigning server a new instance of gRPC server.")

	// Create listener
	lis, err := net.Listen(n.getProtocol(), n.getConnectionAddr())
	if err != nil {
		return errors.Wrapf(ErrListenerInitFails, "Connection Address: %v. Error: %v", n.getConnectionAddr(), err)
	}
	log.Info("Successfully created a listener.")

	// Serve
	log.Info("Starting gRPC server.")
	if err = n.server.Serve(lis); err != nil {
		return errors.Wrapf(ErrFailToServeOnListener, "Connection Address: %v. Error: %v", n.getConnectionAddr(), err)
	}

	return nil
}

func (n *NoteClerkServer) constructor(config *Config, db RDBMSAccessor) error {
	if db == nil {
		return ErrServerInitFailsDbNil
	}
	if config == nil {
		return ErrServerInitFailsFromNilConfig
	}

	n.ip = config.ServerIp
	n.port = config.ServerPort
	n.protocol = config.ServerProtocol
	n.connAddr = fmt.Sprintf("%v:%v", n.getIp(), n.getPort())
	n.db = db

	return nil
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
