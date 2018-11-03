package main

import "github.com/pkg/errors"

// A type created to provide enum-like functionality for errors.
type NoteClerkErrorConst int8

// Error types mapped to a constant number.
const (
	ErrDbPostgresInitializeFailsOpenConn                       = 0
	ErrDbPostgresInitializeFailsDbPing                         = 1
	ErrDbPostgresInitializeFailsSchemaCreation                 = 2
	ErrDbPostgresCreateSchemaFailsTableCreation                = 3
	ErrDbPostgresCreateTableFailsAlreadyExists                 = 4
	ErrDbPostgresCreateTableFailsDueToUnexpectedError          = 5
	ErrDbPostgresAddNoteFailsScan                              = 6
	ErrDbPostgresAddNoteFailsToAddNoteTags                     = 7
	ErrDbPostgresAddNoteFailsToAddNoteFragments                = 8
	ErrDbPostgresUpdateNoteFailsToChangeStatusToDeleted        = 9
	ErrDbPostgresUpdateNoteFailsToAddUpdatedNote               = 10
	ErrDbPostgresGetNoteFragTagByNoteGuidQueryFails            = 11
	ErrDbPostgresGetNoteFragTagByNoteGuidFailsRowScan          = 12
	ErrDbPostgresGetNoteTagsByNoteGuidQueryFails               = 13
	ErrDbPostgresGetNoteTagsByNoteGuidFailsRowScan             = 14
	ErrDbPostgresAddNoteTagFailsScan                           = 15
	ErrDbPostgresAddNoteFragmentFailsScan                      = 16
	ErrDbPostgresAddNoteFragmentFailsAddNoteTags               = 17
	ErrDbPostgresAddNoteFragmentTagFailsScan                   = 18
	ErrNoteClerkServerCreateNoteFailsAddNoteToDb               = 19
	ErrNoteClerkServerDeleteNoteFailsDeleteNoteFromDb          = 20
	ErrNoteClerkServerRetrieveNoteFailsToGetNoteFromDb         = 21
	ErrNoteClerkServerSearchNotesFailsToFindNotesInDb          = 22
	ErrNoteClerkServerUpdateNoteFailsDueToIdMismatch           = 23
	ErrNoteClerkServerUpdateNoteFailsToUpdateNoteInDb          = 24
	ErrNoteClerkServerInitializeFailsDbInitialization          = 25
	ErrNoteClerkServerInitializeFailsCreateListener            = 26
	ErrNoteClerkServerInitializeFailsInitializingRpcServer     = 27
	ErrNoteClerkServerConstructorFailsDueToNilDb               = 28
	ErrNoteClerkServerConstructorFailsDueToNilConfig           = 29
	ErrLoadConfigurationFailsReadFile                          = 30
	ErrLoadConfigurationFailsJsonMarshal                       = 31
	ErrInitializeLoggerFailsOpenLogFile                        = 32
	ErrNoteClerkServerCreateNoteRejectsNoteDueToId             = 33
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsQuery          = 34
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsScan           = 35
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsGetTags        = 36
	ErrDbPostgresAllNotesFailsQuery                            = 37
	ErrDbPostgresAllNotesFailsScan                             = 38
	ErrDbPostgresFindNotesFailsQuery                           = 39
	ErrDbPostgresFindNotesFailsScan                            = 40
	ErrDbPostgresFindNotesFailsGetTags                         = 41
	ErrDbPostgresFindNotesFailsGetNoteFragments                = 42
	ErrDbPostgresAllNoteFragmentsFailsScanRow                  = 43
	ErrDbPostgresAllNoteFragmentsQueryFails                    = 44
	ErrLoadConfigurationAbortsAfterJsonMarshalDueToEmptyConfig = 45
)

// Map NoteClerkErrorConst constants to a string messages, which can be used to produce precise error messages.
var errToMsg = map[NoteClerkErrorConst]string{
	ErrDbPostgresInitializeFailsOpenConn:                       "DbPostgres.Initialize failed to open a database connection",
	ErrDbPostgresInitializeFailsDbPing:                         "DbPostgres.Initialize failed to ping the database.",
	ErrDbPostgresInitializeFailsSchemaCreation:                 "DbPostgres.Initialize failed to create the database schema.",
	ErrDbPostgresCreateSchemaFailsTableCreation:                "DbPostgres.createSchema failed to create table.",
	ErrDbPostgresCreateTableFailsAlreadyExists:                 "DbPostgres.createTable failed to create table; requested table already exists.",
	ErrDbPostgresCreateTableFailsDueToUnexpectedError:          "DbPostgres.createTable failed to create table; error was unexpected.",
	ErrDbPostgresAddNoteFailsScan:                              "DbPostgres.AddNote failed to successfully scan query result for new Id.",
	ErrDbPostgresAddNoteFailsToAddNoteTags:                     "DbPostgres.AddNote failed to add tags to the database.",
	ErrDbPostgresAddNoteFailsToAddNoteFragments:                "DbPostgres.AddNote failed to add note fragments to the database",
	ErrDbPostgresUpdateNoteFailsToChangeStatusToDeleted:        "DbPostgres.UpdateNote failed to change note status to deleted",
	ErrDbPostgresUpdateNoteFailsToAddUpdatedNote:               "DbPostgres.UpdateNote fails to add the updated note.",
	ErrDbPostgresGetNoteFragTagByNoteGuidQueryFails:            "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid query failed.",
	ErrDbPostgresGetNoteFragTagByNoteGuidFailsRowScan:          "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid fails scan of row.",
	ErrDbPostgresGetNoteTagsByNoteGuidQueryFails:               "DbPostgres.GetNoteTagsByNoteGuid query failed.",
	ErrDbPostgresGetNoteTagsByNoteGuidFailsRowScan:             "DbPostgres.GetNoteTagsByNoteGuid fails scan of row.",
	ErrDbPostgresAddNoteTagFailsScan:                           "DbPostgres.AddNoteTag fails to scan new Id",
	ErrDbPostgresAddNoteFragmentFailsScan:                      "DbPostgres.AddNoteFragment fails to scan new Id",
	ErrDbPostgresAddNoteFragmentFailsAddNoteTags:               "DbPostgres.AddNoteFragment fails to add note fragment tags.",
	ErrDbPostgresAddNoteFragmentTagFailsScan:                   "DbPostgres.AddNoteFragmentTag fails to scan Id",
	ErrNoteClerkServerCreateNoteFailsAddNoteToDb:               "Server.CreateNote fails to add the new note to the database.",
	ErrNoteClerkServerDeleteNoteFailsDeleteNoteFromDb:          "Server.DeleteNote fails to delete the requested note from the database.",
	ErrNoteClerkServerRetrieveNoteFailsToGetNoteFromDb:         "Server.RetrieveNote fails to retrieve requested note from the database.",
	ErrNoteClerkServerSearchNotesFailsToFindNotesInDb:          "Server.SearchNotes",
	ErrNoteClerkServerUpdateNoteFailsDueToIdMismatch:           "Server.UpdateNote fails to update due to a mismatch between the Id of the presented note and the Id stated as the note to update.",
	ErrNoteClerkServerUpdateNoteFailsToUpdateNoteInDb:          "Server.UpdateNote fails to update the note in the database.",
	ErrNoteClerkServerInitializeFailsDbInitialization:          "Server.Initialize failed to initialize the database.",
	ErrNoteClerkServerInitializeFailsCreateListener:            "Server.Initialize fails to create a listener.",
	ErrNoteClerkServerInitializeFailsInitializingRpcServer:     "Server.Initialize fails to create a gRPC server on the listener.",
	ErrNoteClerkServerConstructorFailsDueToNilDb:               "Server.constructor fails to populate fields due to a nil database being received.",
	ErrNoteClerkServerConstructorFailsDueToNilConfig:           "Server.constructor fails to populate fields due to a nil config being received.",
	ErrLoadConfigurationFailsReadFile:                          "LoadConfiguration was unable to load a file at the given path.",
	ErrLoadConfigurationFailsJsonMarshal:                       "LoadConfiguration failed to unmarshal the files presumed JSON contents into the configuration structure.",
	ErrInitializeLoggerFailsOpenLogFile:                        "InitializeLogger fails to open the logging file; please check the config for that a proper path has been set.",
	ErrNoteClerkServerCreateNoteRejectsNoteDueToId:             "Server.CreateNote expects an Id of zero. Non-zero values suggest the note may exist already. Please confirm this is a new Note and Create is needed rather than an Update operation.",
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsQuery:          "DbPostgres.GetNoteFragmentsByNoteGuid failed to successfully return a query result with the given GUID string.",
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsScan:           "DbPostgres.GetNoteFragmentsByNoteGuid failed to scan results for the note fragment.",
	ErrDbPostgresGetNoteFragmentsByNoteGuidFailsGetTags:        "DbPostgres.GetNoteFragmentsByNoteGuid failed to retrieve tags.",
	ErrDbPostgresAllNotesFailsQuery:                            "DbPostgres.AllNotes failed to complete the query to retrieve all notes",
	ErrDbPostgresAllNotesFailsScan:                             "DbPostgres.AllNotes failed to scan one or more result rows from the result set.",
	ErrDbPostgresFindNotesFailsQuery:                           "DbPostgres.FindNotes fails to complete query based on data provided in search filter.",
	ErrDbPostgresFindNotesFailsScan:                            "DbPostgres.FindNotes fails to scan one or more result rows from the result set.",
	ErrDbPostgresFindNotesFailsGetTags:                         "DbPostgres.FindNotes fails to get tags.",
	ErrDbPostgresFindNotesFailsGetNoteFragments:                "DbPostgres.FindNotes fails to get note fragments.",
	ErrDbPostgresAllNoteFragmentsFailsScanRow:                  "DbPostgres.AllNoteFragments failed to scan a row.",
	ErrDbPostgresAllNoteFragmentsQueryFails:                    "DbPostgres.AllNoteFragments failed query.",
	ErrLoadConfigurationAbortsAfterJsonMarshalDueToEmptyConfig: "LoadConfiguration is aborting due to receiving an empty configuration file.",
}

func NoteClerkErrWrap(err error, nce NoteClerkErrorConst) error {
	return errors.WithMessage(err, errToMsg[nce])
}

func NoteClerkErr(nce NoteClerkErrorConst) error {
	return errors.New(errToMsg[nce])
}