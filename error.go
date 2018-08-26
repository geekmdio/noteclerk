package main

import "github.com/pkg/errors"

var (
	ErrServerInitFailsFromNilConfig    = errors.New("ERROR 0: Configuration provided to server is nil and could not be used for instantiation.")
	ErrNewNoteNotAddedToDb             = errors.New("ERROR 1: NewNote could not add note to database.")
	ErrDeleteNoteFailsUpdateStatus     = errors.New("ERROR 2: DeleteNote could not set the status to deleted in the database.")
	ErrRetrieveNoteFailsRetrieveFromDb = errors.New("ERROR 3: RetrieveNote could not retrieve note from database.")
	ErrFindNoteFailedToFindInDb        = errors.New("ERROR 4: FindNotes could not locate notes in the database matching the query.")
	ErrUpdateNoteFailedFromIdMismatch  = errors.New("ERROR 5: UpdateNote discovered the id requested for update did not match the id of the provided note and the update request was rejected.")
	ErrUpdateNoteFailedToUpdateInDb    = errors.New("ERROR 6: UpdateNote could not update the note in the database.")
	ErrDbInitFails                        = errors.New("ERROR 7: Initialize could not initialize server due to database initialization failure.")
	ErrListenerInitFails                  = errors.New("ERROR 8: Initialize could not initialize server due to failure to establish a listener.")
	ErrFailToServeOnListener              = errors.New("ERROR 9: Initialize could not initialize server due to a failure to serve on the provided listener.")
	ErrInitFailsOpenConn                  = errors.New("ERROR 10: Initialize could not open a connection to the database using the provided connection string.")
	ErrInitFailsPingDb                    = errors.New("ERROR 11: Initialize failed to ping the database. Successful database connection is unlikely.")
	ErrAddNoteFailsScan                   = errors.New("ERROR 12: AddNote failed to scan a new Id; successful add to database is unlikely.")
	ErrAddNoteFailsAdd                    = errors.New("ERROR 13: AddNote failed to add note fragments from the note to the database. Consider deleting note.")
	ErrAddNoteFragmentFailsScan           = errors.New("ERROR 14: AddNoteFragment failed to scan a new Id; successful add to the database is unlikely.")
	ErrCreateTableFails                   = errors.New("ERROR 15: createTable failed to create a table with the given query.")
	ErrCreateSchemaFails                  = errors.New("ERROR 16: createSchema failed to create the full database schema with the given set of queries.")
	//ErrCreateSchemaFails 			   = errors.New("ERROR 17: Initialize failed to create the database schema.") // Deprecated
	ErrTableAlreadyExists              = errors.New("ERROR 18: Table already exists and will not be replaced.")
	ErrEnvironmentNotSet               = errors.New("ERROR 19: The NOTECLERK_ENVIRONMENT environmental variable has not been set.")
	ErrAddNoteTagFailsScan             = errors.New("ERROR 20: AddNoteTag failed to scan a new Id; successful add to database unlikely.")
	ErrAddNoteFailsAddTagToDb          = errors.New("ERROR 21: AddNote failed to add new note tags to database.")
	ErrAddNoteFragmentFailsAddTagToDb  = errors.New("ERROR 22: AddNoteFragment failed to add tags to database.")
	ErrGetFragTagsByFragGuidFailsQuery = errors.New("ERROR 23: GetNoteFragmentTagsByNoteFragmentGuid had an error during query.")
	ErrGetFragTagsByFragGuidFailsScan  = errors.New("ERROR 24: GetNoteFragmentTagsByNoteFragmentGuid failed to scan row to retrieve tag.")
	ErrGetNoteTagsByNoteGuidFailsQuery = errors.New("ERROR 23: GetNoteTagsByNoteGuid had an error during query.")
	ErrGetNoteTagsByNoteGuidFailsScan  = errors.New("ERROR 24: GetNoteTagsByNoteGuid failed to scan row to retrieve tag.")
	ErrServerInitFailsDbNil            = errors.New("ERROR 25: Initialize function has received a nil database and cannot initialize.")
)
