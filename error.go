package main

type NoteClerkError int8

const (
	DbPostgresInitializeFailsOpenConn                = 0
	DbPostgresInitializeFailsDbPing                  = 1
	DbPostgresInitializeFailsSchemaCreation          = 2
	DbPostgresCreateSchemaFailsTableCreation         = 3
	DbPostgresCreateTableFailsAlreadyExists          = 4
	DbPostgresCreateTableFailsDueToUnexpectedError   = 5
	DbPostgresAddNoteFailsScan                       = 6
	DbPostgresAddNoteFailsToAddNoteTags              = 7
	DbPostgresAddNoteFailsToAddNoteFragments         = 8
	DbPostgresUpdateNoteFailsToChangeStatusToDeleted = 9
	DbPostgresUpdateNoteFailsToAddUpdatedNote        = 10
	DbPostgresGetNoteFragTagByNoteGuidQueryFails     = 11
	DbPostgresGetNoteFragTagByNoteGuidFailsRowScan   = 12
	DbPostgresGetNoteTagsByNoteGuidQueryFails        = 13
	DbPostgresGetNoteTagsByNoteGuidFailsRowScan      = 14
	DbPostgresAddNoteTagFailsScan                    = 15
	DbPostgresAddNoteFragmentFailsScan               = 16
	DbPostgresAddNoteFragmentFailsAddNoteTags        = 17
	DbPostgresAddNoteFragmentTagFailsScan = 18
)

var ErrorMap = map[NoteClerkError]string{
	DbPostgresInitializeFailsOpenConn:                "DbPostgres.Initialize failed to open a database connection",
	DbPostgresInitializeFailsDbPing:                  "DbPostgres.Initialize failed to ping the database.",
	DbPostgresInitializeFailsSchemaCreation:          "DbPostgres.Initialize failed to create the database schema.",
	DbPostgresCreateSchemaFailsTableCreation:         "DbPostgres.createSchema failed to create table.",
	DbPostgresCreateTableFailsAlreadyExists:          "DbPostgres.createTable failed to create table; requested table already exists.",
	DbPostgresCreateTableFailsDueToUnexpectedError:   "DbPostgres.createTable failed to create table; error was unexpected.",
	DbPostgresAddNoteFailsScan:                       "DbPostgres.AddNote failed to successfully scan query result for new Id.",
	DbPostgresAddNoteFailsToAddNoteTags:              "DbPostgres.AddNote failed to add tags to the database.",
	DbPostgresAddNoteFailsToAddNoteFragments:         "DbPostgres.AddNote failed to add note fragments to the database",
	DbPostgresUpdateNoteFailsToChangeStatusToDeleted: "DbPostgres.UpdateNote failed to change note status to deleted",
	DbPostgresUpdateNoteFailsToAddUpdatedNote:        "DbPostgres.UpdateNote fails to add the updated note.",
	DbPostgresGetNoteFragTagByNoteGuidQueryFails:     "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid query failed.",
	DbPostgresGetNoteFragTagByNoteGuidFailsRowScan:   "DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid fails scan of row.",
	DbPostgresGetNoteTagsByNoteGuidQueryFails:        "DbPostgres.GetNoteTagsByNoteGuid query failed.",
	DbPostgresGetNoteTagsByNoteGuidFailsRowScan:      "DbPostgres.GetNoteTagsByNoteGuid fails scan of row.",
	DbPostgresAddNoteTagFailsScan:                    "DbPostgres.AddNoteTag fails to scan new Id",
	DbPostgresAddNoteFragmentFailsScan:               "DbPostgres.AddNoteFragment fails to scan new Id",
	DbPostgresAddNoteFragmentFailsAddNoteTags:        "DbPostgres.AddNoteFragment fails to add note fragment tags.",
	DbPostgresAddNoteFragmentTagFailsScan: "DbPostgres.AddNoteFragmentTag fails to scan Id",
}
