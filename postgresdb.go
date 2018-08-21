package main

import (
	"database/sql"
	"fmt"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
)

type DbPostgres struct {
	db *sql.DB
}

// Initialize() initializes the connection to database. Ensure that the ./config/config.<environment>.json
// file has been created and properly configured with server and database values. Of note, the '<environment>'
// can be set to any value, so long as the NOTECLERK_ENVIRONMENT environmental variable's value matches.
// RETURNS: *sql.db, error
func (d *DbPostgres) Initialize(config *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		config.DbUsername, config.DbPassword, config.DbIp, config.DbName, config.DbSslMode, config.DbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, errors.Wrapf(ErrPostgresDbInitFailedToOpenConn, "%v", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		return nil, errors.Wrapf(ErrPostgresDbInitFailedToPingDb, "%v", err)
	}

	d.db = db
	schemaErr := d.createSchema()
	if schemaErr != nil {
		return nil, errors.Wrapf(ErrPostgresDbInitFailedToCreateSchema, "%v", schemaErr)
	}

	return d.db, nil
}

func (d *DbPostgres) GetNoteFragmentsByNoteGuid(noteGuid string) ([]*ehrpb.NoteFragment, error) {
	rows, err := d.db.Query(getNoteFragmentByNoteGuid, noteGuid)
	if err != nil {
		//TODO: Custom error
		return nil, err
	}
	defer rows.Close()

	noteFragments := make([]*ehrpb.NoteFragment, 0)
	for rows.Next() {
		tmp := NewNoteFragment()
		err := rows.Scan(&tmp.Id, &tmp.DateCreated.Seconds, &tmp.DateCreated.Nanos, &tmp.NoteFragmentGuid,
			&tmp.NoteGuid, &tmp.Icd_10Code, &tmp.Icd_10Long, &tmp.Description, &tmp.Status,
			&tmp.Priority, &tmp.Topic, &tmp.Content)
		if err != nil {
			//TODO: Custom error
			return nil, err
		}
		tmp.Tags, err = d.GetNoteFragmentTagsByNoteFragmentGuid(tmp.GetNoteFragmentGuid())
		if err != nil {
			//TODO: Custom error
			return nil, err
		}
		noteFragments = append(noteFragments, tmp)
	}
	return noteFragments, nil
}

func (d *DbPostgres) GetNoteTagsByNoteGuid(noteGuid string) (tag []string, err error) {
	rows, err := d.db.Query(getNoteTagByNoteGuid, noteGuid)
	if err != nil {
		return nil, errors.Wrapf(ErrPostgresDbGetNoteTagsByNoteGuidFailsToQueryResults, "%v", err)
	}
	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var id int64
		var noteGuid string
		var tmpString string
		err := rows.Scan(&id, &noteGuid, &tmpString)
		if err != nil {
			return nil, errors.Wrapf(ErrPostgresDbGetNoteTagsByNoteGuidFailsToScanResults, "%v", err)
		}
		tags = append(tags, tmpString)
	}
	return tags, nil
}

func (d *DbPostgres) GetNoteFragmentTagsByNoteFragmentGuid(noteFragGuid string) (tag []string, err error) {
	rows, err := d.db.Query(getNoteFragmentTagsByNoteFragmentGuid, noteFragGuid)
	if err != nil {
		return nil, errors.Wrapf(ErrPostgresDbGetNoteFragmentTagsByNoteFragmentGuidFailsToQueryResults, "%v", err)
	}
	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var tmpId int64
		var tmpNoteFragmentGuid string
		var tmpTag string
		err := rows.Scan(&tmpId, &tmpNoteFragmentGuid, &tmpTag)
		if err != nil {
			return nil, errors.Wrapf(ErrPostgresDbGetNoteFragmentTagsByNoteFragmentGuidFailsToScanResults, "%v", err)
		}
		tags = append(tags, tmpTag)
	}
	return tags, nil
}

func (d *DbPostgres) AddNote(n *ehrpb.Note) (id int64, err error) {

	row := d.db.QueryRow(addNoteQuery, n.DateCreated.GetSeconds(), n.DateCreated.GetNanos(),
		n.GetNoteGuid(), n.GetVisitGuid(), n.GetAuthorGuid(), n.GetPatientGuid(), n.GetType(),
		n.GetStatus())

	if scanErr := row.Scan(&n.Id); scanErr != nil && scanErr != sql.ErrNoRows {
		return 0, errors.Wrapf(ErrPostgresDbAddNoteFailedToGetNewId, "%v", scanErr)
	}

	for _, v := range n.GetTags() {
		_, err := d.AddNoteTag(n.GetNoteGuid(), v)
		if err != nil && err != sql.ErrNoRows {
			return 0, errors.Wrapf(ErrPostgresDbAddNoteFailedToAddNoteTagToDb, "%v", err)
		}
	}

	for _, v := range n.GetFragments() {
		_, _, err := d.AddNoteFragment(v)
		if err != nil && err != sql.ErrNoRows {
			return 0, errors.Wrapf(ErrPostgresDbAddNoteFailedToAddNoteFragments, "%v", err)
		}
	}

	return n.Id, nil
}

func (d *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	delErr := d.DeleteNote(n.GetId())
	if delErr != nil && delErr != sql.ErrNoRows {
		return delErr
	}
	n.NoteGuid = uuid.New().String()
	for _, v := range n.GetFragments() {
		v.NoteFragmentGuid = uuid.New().String()
		v.NoteGuid = n.NoteGuid
	}
	_, addErr := d.AddNote(n)
	if addErr != nil {
		return addErr
	}
	return nil
}

func (d *DbPostgres) DeleteNote(id int64) error {
	note, getNoteErr := d.GetNoteById(id)
	if getNoteErr != nil {
		return getNoteErr
	}
	for _, v := range note.GetFragments() {
		delErr := d.DeleteNoteFragment(v.GetNoteFragmentGuid())
		if delErr != nil {
			return delErr
		}
	}
	row := d.db.QueryRow(updateNoteStatusToStatusByNoteIdQuery, ehrpb.RecordStatus_DELETED, id)
	var newId int64
	scanErr := row.Scan(newId)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return scanErr
	}
	return nil
}

func (d *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	rows, err := d.db.Query("SELECT * FROM note;")
	if err != nil {
		// TODO: Custom error giving more context.
		return nil, err
	}
	defer rows.Close()

	var notes []*ehrpb.Note
	for rows.Next() {
		tmpNote := NewNote()
		err := rows.Scan(&tmpNote.Id, &tmpNote.DateCreated.Seconds, &tmpNote.DateCreated.Nanos,
			&tmpNote.NoteGuid, &tmpNote.VisitGuid, &tmpNote.AuthorGuid,
			&tmpNote.PatientGuid, &tmpNote.Type, &tmpNote.Status)
		if err != nil {
			//TODO: Custom error giving more context.
			return nil, err
		}
		tmpNote.Tags, err = d.GetNoteTagsByNoteGuid(tmpNote.GetNoteGuid())
		if err != nil {
			return nil, err
		}

		tmpNote.Fragments, err = d.GetNoteFragmentsByNoteGuid(tmpNote.GetNoteGuid())
		if err != nil {
			return nil, err
		}

		notes = append(notes, tmpNote)

	}
	return notes, nil
}

func (d *DbPostgres) AddNoteTag(noteGuid string, tag string) (id int64, err error) {
	row := d.db.QueryRow(addNoteTagQuery, noteGuid, tag)

	var newId int64
	if scanErr := row.Scan(&newId); scanErr != nil && scanErr != sql.ErrNoRows {
		return 0, errors.Wrapf(ErrPostgresDbAddNoteTagFailedToGetNewId, "%v", scanErr)
	}

	return newId, nil
}

func (d *DbPostgres) GetNoteById(id int64) (*ehrpb.Note, error) {
	row := d.db.QueryRow(getNoteByIdQuery, id)

	newNote := NewNote()
	err := row.Scan(&newNote.Id, &newNote.DateCreated.Seconds, &newNote.DateCreated.Nanos, &newNote.NoteGuid,
		&newNote.VisitGuid, &newNote.AuthorGuid, &newNote.PatientGuid, &newNote.Type, &newNote.Status)

	//TODO: Custom errors
	if err != nil {
		return nil, err
	}
	newNote.Fragments, err = d.GetNoteFragmentsByNoteGuid(newNote.GetNoteGuid())
	if err != nil {
		return nil, err
	}

	newNote.Tags, err = d.GetNoteTagsByNoteGuid(newNote.GetNoteGuid())
	if err != nil {
		return nil, err
	}

	return newNote, nil
}

func (d *DbPostgres) FindNotes(filter NoteFindFilter) ([]*ehrpb.Note, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AddNoteFragment(nf *ehrpb.NoteFragment) (id int64, guid string, err error) {
	row := d.db.QueryRow(addNoteFragmentQuery, nf.DateCreated.Seconds, nf.DateCreated.Nanos,
		nf.GetNoteFragmentGuid(), nf.GetNoteGuid(), nf.GetIcd_10Code(), nf.GetIcd_10Long(),
		nf.GetDescription(), nf.GetStatus(), nf.GetPriority(), nf.GetTopic(), nf.GetContent())
	scanErr := row.Scan(&nf.Id)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return 0, "", errors.Wrapf(ErrPostgresDbAddNoteFragmentFailedToGetNewId, "%v", scanErr)
	}

	for _, v := range nf.GetTags() {
		_, err := d.AddNoteFragmentTag(nf.GetNoteFragmentGuid(), v)
		if err != nil && err != sql.ErrNoRows {
			return 0, nf.NoteFragmentGuid, errors.Wrapf(ErrPostgresDbAddNoteFragmentFailedToAddNoteFragmentTagToDb, "%v", err)
		}
	}

	return nf.GetId(), nf.GetNoteFragmentGuid(), nil
}

func (d *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment) error {
	err := d.DeleteNoteFragment(n.GetNoteFragmentGuid())
	if err != nil {
		return err
	}
	log.Fatal("Not implemented.")
	return nil
}

// This is not a true delete. It changes the status of the note to DELETED. Health care
// records should not be deleted.
func (d *DbPostgres) DeleteNoteFragment(noteFragmentGuid string) error {
	row := d.db.QueryRow(updateNoteFragmentStatusToStatusByNoteFragmentGuidQuery, ehrpb.RecordStatus_DELETED, noteFragmentGuid)
	var newId int64
	scanErr := row.Scan(newId)
	if scanErr != nil && scanErr != sql.ErrNoRows {
		return scanErr
	}
	return nil
}

func (d *DbPostgres) GetNoteFragmentById(id int64) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AddNoteFragmentTag(noteGuid string, tag string) (id int64, err error) {
	row := d.db.QueryRow(addNoteFragmentTagQuery, noteGuid, tag)

	var newId int64
	if scanErr := row.Scan(&newId); scanErr != nil && scanErr != sql.ErrNoRows {
		return 0, errors.Wrapf(ErrPostgresDbAddNoteTagFailedToGetNewId, "%v", scanErr)
	}

	return newId, nil
}

// https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func (d *DbPostgres) createSchema() error {
	err := d.createTable(createNoteTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(ErrPostgresDbCreateSchemaFails, "Target table: note. Error: %v", err)
	}
	if err == ErrPostgresDbInitTableAlreadyExistsErr {
		log.Warn("Table 'note' already exists.")
	}
	err = d.createTable(createNoteTagTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(ErrPostgresDbCreateSchemaFails, "Target table: note_tag. Error: %v", err)
	}
	if err == ErrPostgresDbInitTableAlreadyExistsErr {
		log.Warn("Table 'note_tag' already exists.")
	}
	err = d.createTable(createNoteFragmentTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(ErrPostgresDbCreateSchemaFails, "Target table: note_fragment. Error: %v", err)
	}
	if err == ErrPostgresDbInitTableAlreadyExistsErr {
		log.Warn("Table 'note_fragment' already exists.")
	}
	err = d.createTable(createNoteFragmentTagTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(ErrPostgresDbCreateSchemaFails, "Target table: note_fragment_tag. Error: %v", err)
	}
	if err == ErrPostgresDbInitTableAlreadyExistsErr {
		log.Warn("Table 'note_fragment_tag' already exists.")
	}

	//TODO: Remove this.
	tmpNote := NewNote()
	tmpNote.Tags = append(tmpNote.Tags, "note1Tag1", "note1Tag2")

	tmpFrag := NewNoteFragment()
	tmpFrag.NoteGuid = tmpNote.GetNoteGuid()
	tmpFrag.Tags = append(tmpFrag.Tags, "frag1Tag1", "frag1Tag2")
	tmpFrag.Content = "This is my content for frag1"
	err = d.DeleteNoteFragment(tmpFrag.GetNoteFragmentGuid())
	if err != nil {
		log.Warn(err)
	}

	tmpFrag2 := NewNoteFragment()
	tmpFrag2.NoteGuid = tmpNote.GetNoteGuid()
	tmpFrag2.Tags = append(tmpFrag2.Tags, "frag2Tag1", "frag2Tag2")
	tmpFrag2.Content = "This is my content for frag2"

	tmpNote.Fragments = append(tmpNote.Fragments, tmpFrag, tmpFrag2)

	_, err = d.AddNote(tmpNote)
	if err != nil {
		log.Warn(err)
	}

	notes, notesErr := d.AllNotes()
	if notesErr != nil {
		log.Warn(notesErr)
	}

	tmpNote.Tags = append(tmpNote.Tags, "updatedTag")
	updateErr := d.UpdateNote(tmpNote)
	if updateErr != nil {
		log.Warn(updateErr)
	}
	fmt.Println(notes)
	//TODO: End remove

	return nil
}

func (d *DbPostgres) createTable(query string) error {
	_, err := d.db.Exec(query)

	tableExistsError := strings.Contains(fmt.Sprintf("%v", err), "already exists")
	if tableExistsError {
		return ErrPostgresDbInitTableAlreadyExistsErr
	}
	if err != nil {
		return errors.Wrapf(ErrPostgresDbCreateTableFails, "%v", err)
	}
	return nil
}

func notNilNotTableExists(err error) bool {
	return err != nil && err != ErrPostgresDbInitTableAlreadyExistsErr
}
