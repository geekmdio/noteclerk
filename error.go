package main

import "github.com/pkg/errors"

var (
	ErrServerInitFailsFromNilConfig = errors.New("ERROR 0: Configuration provided to server is nil and could not be used for instantiation.")
	ErrNewNoteNotAddedToDb = errors.New("ERROR 1: NewNote could not add note to database.")
	ErrDeleteNoteFailedToDeleteFromDb = errors.New("ERROR 2: DeleteNote could not delete note from database.")
	ErrRetrieveNoteFailedToRetrieveFromDb = errors.New("ERROR 3: RetrieveNote could not retrieve note from database.")
	ErrFindNoteFailedToFindInDb = errors.New("ERROR 4: FindNote could not locate notes in the database matching the query.")
	ErrUpdateNoteFailedFromIdMismatch = errors.New("ERROR 5: UpdateNote discovered the id requested for update did not match the id of the provided note and the update request was rejected.")
	ErrUpdateNoteFailedToUpdateInDb = errors.New("ERROR 6: UpdateNote could not update the note in the database.")
	ErrDbInitFails = errors.New("ERROR 7: Initialize could not initialize server due to database initialization failure.")
	ErrListenerInitFails                         = errors.New("ERROR 8: Initialize could not initialize server due to failure to establish a listener.")
	ErrFailToServeOnListener                     = errors.New("ERROR 9: Initialize could not initialize server due to a failure to serve on the provided listener.")
	ErrPostgresDbInitFailedToOpenConn            = errors.New("ERROR 10: DbPostgres.Initialize could not open a connection to the database using the provided connection string.")
	ErrPostgresDbInitFailedToPingDb              = errors.New("ERROR 11: DbPostgres.Initialize failed to ping the database.")
	ErrPostgresDbAddNoteFailedToGetNewId         = errors.New("ERROR 12: DbPostgres.AddNote failed to retrieve a new Id; successful add to database is unlikely.")
	ErrPostgresDbAddNoteFailedToAddNoteFragments = errors.New("ERROR 13: DbPostgres.AddNote failed to add note fragments from the note to the database. Consider deleting note.")
	ErrPostgresDbAddNoteFragmentFailedToGetNewId = errors.New("ERROR 14: DbPostgres.AddNoteFragment failed to retrieve a new Id; successful add to the database is unlikely.")
	ErrPostgresDbCreateTableFails                = errors.New("ERROR 15: DbPostgres.createTable failed to create a table with the given query.")
	ErrPostgresDbCreateSchemaFails = errors.New("ERROR 16: DbPostgres.createSchema failed to create the full database schema with the given set of queries.")
	ErrPostgresDbInitFailedToCreateSchema = errors.New("ERROR 17: DbPostgres.Initialize failed to create the full database schema.")
	ErrPostgresDbInitTableAlreadyExistsErr = errors.New("ERROR 18: Table already exists and will not be replaced.")
	ErrMainEnvironmentalVariableNotSet = errors.New("ERROR 19: The NOTECLERK_ENVIRONMENT environmental variable was note set; cannot load configuration file.")
	ErrPostgresDbAddNoteTagFailedToGetNewId = errors.New("ERROR 20: DbPostgres.AddNoteTag failed to retrieve a new Id; successful add to database unlikely.")
	ErrPostgresDbAddNoteFailedToAddNoteTagToDb = errors.New("ERROR 21: DbPostgres.AddNote failed to add new note tags to database.")
	ErrPostgresDbAddNoteFragmentFailedToAddNoteFragmentTagToDb =  errors.New("ERROR 22: DbPostgres.AddNoteFragment failed to add tags to database.")
	ErrPostgresDbGetNoteFragmentTagsByNoteFragmentGuidFailsToQueryResults = errors.New("ERROR 23: DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid had an error during query.")
	ErrPostgresDbGetNoteFragmentTagsByNoteFragmentGuidFailsToScanResults = errors.New("ERROR 24: DbPostgres.GetNoteFragmentTagsByNoteFragmentGuid failed to scan row to retrieve tag.")
	ErrPostgresDbGetNoteTagsByNoteGuidFailsToQueryResults = errors.New("ERROR 23: DbPostgres.GetNoteTagsByNoteGuid had an error during query.")
	ErrPostgresDbGetNoteTagsByNoteGuidFailsToScanResults = errors.New("ERROR 24: DbPostgres.GetNoteTagsByNoteGuid failed to scan row to retrieve tag.")

)
