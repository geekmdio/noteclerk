package main

type NoteClerkError int8

const (
	DbPostgresInitializeFailsOpenConn                   = 0
	DbPostgresInitializeFailsDbPing                     = 1
	DbPostgresInitializeFailsSchemaCreation             = 2
	DbPostgresCreateSchemaFailsTableCreation            = 3
	DbPostgresCreateTableFailsAlreadyExists             = 4
	DbPostgresCreateTableFailsDueToUnexpectedError      = 5
	DbPostgresAddNoteFailsScan                          = 6
	DbPostgresAddNoteFailsToAddNoteTags                 = 7
	DbPostgresAddNoteFailsToAddNoteFragments            = 8
	DbPostgresUpdateNoteFailsToChangeStatusToDeleted    = 9
	DbPostgresUpdateNoteFailsToAddUpdatedNote           = 10
	DbPostgresGetNoteFragTagByNoteGuidQueryFails        = 11
	DbPostgresGetNoteFragTagByNoteGuidFailsRowScan      = 12
	DbPostgresGetNoteTagsByNoteGuidQueryFails           = 13
	DbPostgresGetNoteTagsByNoteGuidFailsRowScan         = 14
	DbPostgresAddNoteTagFailsScan                       = 15
	DbPostgresAddNoteFragmentFailsScan                  = 16
	DbPostgresAddNoteFragmentFailsAddNoteTags           = 17
	DbPostgresAddNoteFragmentTagFailsScan               = 18
	NoteClerkServerCreateNoteFailsAddNoteToDb           = 19
	NoteClerkServerDeleteNoteFailsDeleteNoteFromDb      = 20
	NoteClerkServerRetrieveNoteFailsToGetNoteFromDb     = 21
	NoteClerkServerSearchNotesFailsToFindNotesInDb      = 22
	NoteClerkServerUpdateNoteFailsDueToIdMismatch       = 23
	NoteClerkServerUpdateNoteFailsToUpdateNoteInDb      = 24
	NoteClerkServerInitializeFailsDbInitialization      = 25
	NoteClerkServerInitializeFailsCreateListener        = 26
	NoteClerkServerInitializeFailsInitializingRpcServer = 27
	NoteClerkServerConstructorFailsDueToNilDb           = 28
	NoteClerkServerConstructorFailsDueToNilConfig       = 29
	LoadConfigurationFailsReadFile                      = 30
	LoadConfigurationFailsJsonMarshal                   = 31
	InitializeLoggerFailsOpenLogFile                    = 32
	NoteClerkServerCreateNoteRejectsNoteDueToId         = 33
)

var ErrMapStr = map[NoteClerkError]string{
	DbPostgresInitializeFailsOpenConn:                   "DbPostgres.Initialize failed to open a database connection",
	DbPostgresInitializeFailsDbPing:                     "DbPostgres.Initialize failed to ping the database.",
	DbPostgresInitializeFailsSchemaCreation:             "DbPostgres.Initialize failed to create the database schema.",
	DbPostgresCreateSchemaFailsTableCreation:            "DbPostgres.createSchema failed to create table.",
	DbPostgresCreateTableFailsAlreadyExists:             "DbPostgres.createTable failed to create table; requested table already exists.",
	DbPostgresCreateTableFailsDueToUnexpectedError:      "DbPostgres.createTable failed to create table; error was unexpected.",
	DbPostgresAddNoteFailsScan:                          "DbPostgres.AddNote failed to successfully scan query result for new Id.",
	DbPostgresAddNoteFailsToAddNoteTags:                 "DbPostgres.AddNote failed to add tags to the database.",
	DbPostgresAddNoteFailsToAddNoteFragments:            "DbPostgres.AddNote failed to add note fragments to the database",
	DbPostgresUpdateNoteFailsToChangeStatusToDeleted:    "DbPostgres.UpdateNote failed to change note status to deleted",
	DbPostgresUpdateNoteFailsToAddUpdatedNote:           "DbPostgres.UpdateNote fails to add the updated note.",
	DbPostgresGetNoteFragTagByNoteGuidQueryFails:        "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid query failed.",
	DbPostgresGetNoteFragTagByNoteGuidFailsRowScan:      "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid fails scan of row.",
	DbPostgresGetNoteTagsByNoteGuidQueryFails:           "DbPostgres.GetNoteTagsByNoteGuid query failed.",
	DbPostgresGetNoteTagsByNoteGuidFailsRowScan:         "DbPostgres.GetNoteTagsByNoteGuid fails scan of row.",
	DbPostgresAddNoteTagFailsScan:                       "DbPostgres.AddNoteTag fails to scan new Id",
	DbPostgresAddNoteFragmentFailsScan:                  "DbPostgres.AddNoteFragment fails to scan new Id",
	DbPostgresAddNoteFragmentFailsAddNoteTags:           "DbPostgres.AddNoteFragment fails to add note fragment tags.",
	DbPostgresAddNoteFragmentTagFailsScan:               "DbPostgres.AddNoteFragmentTag fails to scan Id",
	NoteClerkServerCreateNoteFailsAddNoteToDb:           "NoteClerkServer.CreateNote fails to add the new note to the database.",
	NoteClerkServerDeleteNoteFailsDeleteNoteFromDb:      "NoteClerkServer.DeleteNote fails to delete the requested note from the database.",
	NoteClerkServerRetrieveNoteFailsToGetNoteFromDb:     "NoteClerkServer.RetrieveNote fails to retrieve requested note from the database.",
	NoteClerkServerSearchNotesFailsToFindNotesInDb:      "NoteClerkServer.SearchNotes",
	NoteClerkServerUpdateNoteFailsDueToIdMismatch:       "NoteClerkServer.UpdateNote fails to update due to a mismatch between the Id of the presented note and the Id stated as the note to update.",
	NoteClerkServerUpdateNoteFailsToUpdateNoteInDb:      "NoteClerkServer.UpdateNote fails to update the note in the database.",
	NoteClerkServerInitializeFailsDbInitialization:      "NoteClerkServer.Initialize failed to initialize the database.",
	NoteClerkServerInitializeFailsCreateListener:        "NoteClerkServer.Initialize fails to create a listener.",
	NoteClerkServerInitializeFailsInitializingRpcServer: "NoteClerkServer.Initialize fails to create a gRPC server on the listener.",
	NoteClerkServerConstructorFailsDueToNilDb:           "NoteClerkServer.constructor fails to populate fields due to a nil database being received.",
	NoteClerkServerConstructorFailsDueToNilConfig:       "NoteClerkServer.constructor fails to populate fields due to a nil config being received.",
	LoadConfigurationFailsReadFile:                      "LoadConfiguration was unable to load a file at the given path.",
	LoadConfigurationFailsJsonMarshal:                   "LoadConfiguration failed to unmarshal the files presumed JSON contents into the configuration structure.",
	InitializeLoggerFailsOpenLogFile:                    "InitializeLogger fails to open the logging file; please check the config for that a proper path has been set.",
	NoteClerkServerCreateNoteRejectsNoteDueToId:         "NoteClerkServer.CreateNote expects an Id of zero. Non-zero values suggest the note may exist already. Please confirm this is a new Note and Create is needed rather than an Update operation.",
}
