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

// NoteClerkServer implements the gRPC NoteServiceServer interface. It has basic CRUD functionality, plus the ability to search
// both Notes and, independently of Notes, NoteFragments. The ability to search NoteFragment's specifically gives a
// much higher degree of resolution to search findings and exclude the less relevant data.
type NoteClerkServer struct {
	db       RDBMSAccessor
	ip       string
	port     string
	protocol string
	connAddr string
	server   *grpc.Server
}

// CreateNote is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The CreateNoteRequest object carries only a Note to be added. This Note should not have an Id assigned to it, or it
// will likely generate an error when there is an attempt to add it to the database. The CreateNoteResponse contains
// a status, which includes a message and a HttpCode.
// RETURNS: CreateNoteResponse, error
func (n *NoteClerkServer) CreateNote(ctx context.Context, nr *ehrpb.CreateNoteRequest) (*ehrpb.CreateNoteResponse, error) {
	cnr := &ehrpb.CreateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{},
	}

	noteToAdd := nr.Note
	noteToAdd.NoteGuid = uuid.New().String()
	noteToAdd.DateCreated = noted.TimestampNow()

	if nr.Note.GetId() > 0 {
		return nil, errors.New(ErrMapStr[NoteClerkServerCreateNoteRejectsNoteDueToId])
	}

	id, err := n.db.AddNote(noteToAdd)
	if err != nil {
		err := errors.WithMessage(err, ErrMapStr[NoteClerkServerCreateNoteFailsAddNoteToDb])
		log.Warn(err)
		cnr.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		cnr.Status.Message = "Failed to insert new note into database."
		cnr.Note = nil
		return nil, err
	}

	cnr.Note = noteToAdd
	cnr.Note.Id = id
	cnr.Status.HttpCode = ehrpb.StatusCodes_OK
	cnr.Status.Message = "Successfully submit new note."

	return cnr, nil
}

// DeleteNote is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The DeleteNoteRequest object carries only the Id of the target note. The DeleteNoteResponse contains only
// a status, which includes a message and a HttpCode.
// RETURNS: DeleteNoteResponse, error
func (n *NoteClerkServer) DeleteNote(ctx context.Context, dnr *ehrpb.DeleteNoteRequest) (*ehrpb.DeleteNoteResponse, error) {
	dnRes := &ehrpb.DeleteNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully changed the notes status to deleted in the database.",
		},
	}

	err := n.db.DeleteNote(dnr.Id)
	if err != nil {
		err := errors.WithMessage(err, ErrMapStr[NoteClerkServerDeleteNoteFailsDeleteNoteFromDb])
		log.Warn(err)
		dnRes.Status.HttpCode = ehrpb.StatusCodes_NOT_MODIFIED
		dnRes.Status.Message = "Failed to change the notes status to deleted in the database."
		return dnRes, err
	}

	return dnRes, nil
}

// RetrieveNote is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The RetrieveNoteRequest object carries only the Id of the target note. The RetrieveNoteResponse contains a Note and
// a status, which includes a message and a HttpCode.
// RETURNS: RetrieveNoteResponse, error
func (n *NoteClerkServer) RetrieveNote(ctx context.Context, rnr *ehrpb.RetrieveNoteRequest) (*ehrpb.RetrieveNoteResponse, error) {
	res := &ehrpb.RetrieveNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully retrieved note from database.",
		},
	}

	note, err := n.db.GetNoteById(rnr.Id)
	if err != nil {
		err := errors.WithMessage(err, ErrMapStr[NoteClerkServerRetrieveNoteFailsToGetNoteFromDb])
		log.Warn(err)
		res.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		res.Status.Message = "Failed to retrieve note from database."
		return res, err
	}

	err = noted.OrganizeNoteFragments(note)
	if err != nil {
		log.Warn("Could not organize the note fragments by fragment priority.")
	}
	res.Note = note

	return res, nil
}

// SearchNotes is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The SearchNotesRequest object carries GUID's for patient, author, and visit in addition to a field for
// search terms which can scan through contents of not fragments and tags. The SearchNotesResponse contains a slice of
// Note and a status, which includes a message and a HttpCode.
// RETURNS: SearchNotesResponse, error
func (n *NoteClerkServer) SearchNotes(ctx context.Context, fnr *ehrpb.SearchNotesRequest) (*ehrpb.SearchNotesResponse, error) {
	filter := NoteFindFilter{
		VisitGuid:   fnr.VisitGuid,
		AuthorGuid:  fnr.AuthorGuid,
		PatientGuid: fnr.PatientGuid,
		SearchTerms: fnr.SearchTerms,
	}

	res := &ehrpb.SearchNotesResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully found one or more notes matching query.",
		},
	}

	notes, err := n.db.FindNotes(filter)
	if err != nil {
		err := errors.WithMessage(err, ErrMapStr[NoteClerkServerSearchNotesFailsToFindNotesInDb])
		log.Warn(err)
		res.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		res.Status.Message = "Failed to locate notes matching query"
		return res, err
	}

	for _, v := range res.Notes {
		err := noted.OrganizeNoteFragments(v)
		if err != nil {
			log.Warn("Could not organize the note fragments by fragment priority.")
		}
	}

	res.Notes = notes
	return res, nil
}

// UpdateNote is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The UpdateNoteRequest object carries a field for the Id of the target note, and an updated version of the note
// which should have an Id matching the Id field of the UpdateNoteRequest. The UpdateNoteResponse contains a
// status, which includes a message and a HttpCode.
// RETURNS: UpdateNoteResponse, error
func (n *NoteClerkServer) UpdateNote(ctx context.Context, unr *ehrpb.UpdateNoteRequest) (*ehrpb.UpdateNoteResponse, error) {

	updateNoteResponse := &ehrpb.UpdateNoteResponse{
		Status: &ehrpb.NoteServiceResponseStatus{
			HttpCode: ehrpb.StatusCodes_OK,
			Message:  "Successfully updated note.",
		},
	}

	if unr.Id != unr.Note.Id {
		newErr := errors.New(ErrMapStr[NoteClerkServerUpdateNoteFailsDueToIdMismatch])
		log.Warn(newErr)
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_CONFLICT
		updateNoteResponse.Status.Message = "Failed to update note. The id provided for the update note request does not match the id of the note."
		return updateNoteResponse, newErr
	}

	err := n.db.UpdateNote(unr.Note)
	if err != nil {
		newErr := errors.WithMessage(err, ErrMapStr[NoteClerkServerUpdateNoteFailsToUpdateNoteInDb])
		log.Warn(newErr)
		updateNoteResponse.Status.HttpCode = ehrpb.StatusCodes_NOT_FOUND
		updateNoteResponse.Status.Message = "UpdateNote failed. Unable to update note in the database."
		return updateNoteResponse, newErr
	}

	return updateNoteResponse, nil
}

// SearchNoteFragments is a method contracted by the NoteServiceServer interface. It therefore complies with gRPC conventions.
// The SearchNoteFragmentsRequest object carries fields for GUID's of patient, author, visit, and note. There is also a
// search terms field, where search terms will be evaluated against note fragment content and tags. The
// SearchNoteFragmentsResponse contains a slice of NoteFragment and a status, which includes a message and a HttpCode.
// RETURNS: SearchNoteFragmentsResponse, error
func (n *NoteClerkServer) SearchNoteFragments(ctx context.Context, snf *ehrpb.SearchNoteFragmentRequest) (*ehrpb.SearchNoteFragmentResponse, error) {
	panic("implement me")
}

// Initialize takes a configuration file and a struct which implements the RDBMSAccessor interface. That is, generally
// a SQL database using any supported driver. The configuration file carries various useful information, but in the
// context of the Initialize function it's responsible for providing important server and RDBMS connection settings.
// RETURNS: error
func (n *NoteClerkServer) Initialize(config *Config, db RDBMSAccessor) error {
	// Build up the server's fields
	conErr := n.constructor(config, db)
	if conErr != nil {
		return conErr
	}

	// Initialize server database
	err := n.db.Initialize(config)
	if err != nil {
		return errors.WithMessage(err, ErrMapStr[NoteClerkServerInitializeFailsDbInitialization])
	}
	log.Info("Successfully connected to database.")

	// Create and register gRPC server
	n.server = grpc.NewServer()
	ehrpb.RegisterNoteServiceServer(n.server, n)
	log.Info("Assigning server a new instance of gRPC server.")

	// Create listener
	lis, err := net.Listen(n.getProtocol(), n.getConnectionAddr())
	if err != nil {
		return errors.WithMessage(err, ErrMapStr[NoteClerkServerInitializeFailsCreateListener])
	}
	log.Info("Successfully created a listener.")

	// Serve
	log.Info("Starting gRPC server.")
	if err = n.server.Serve(lis); err != nil {
		return errors.WithMessage(err, ErrMapStr[NoteClerkServerInitializeFailsInitializingRpcServer])
	}

	return nil
}

// constructor populates fields belonging to the NoteClerkServer struct. It also validates the state of the
// database and configuration files.
func (n *NoteClerkServer) constructor(config *Config, db RDBMSAccessor) error {
	if db == nil {
		return errors.New(ErrMapStr[NoteClerkServerConstructorFailsDueToNilDb])
	}
	if config == nil {
		return errors.New(ErrMapStr[NoteClerkServerConstructorFailsDueToNilConfig])
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
