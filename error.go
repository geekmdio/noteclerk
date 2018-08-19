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
	ErrListenerInitFails = errors.New("ERROR 8: Initialize could not initialize server due to failure to establish a listener.")
	ErrFailToServeOnListener = errors.New("ERROR 9: Initialize could not initialize server due to a failure to serve on the provided listener.")
	
)
