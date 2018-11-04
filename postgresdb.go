package main

import (
	"database/sql"
	"fmt"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/geekmdio/noted"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"strings"
)

// DbPostgres implements RDBMSAccessor; purpose is to access the database via the Postgres driver.
type DbPostgres struct {
	db *sql.DB
}

// Initialize() initializes the connection to database. Ensure that the ./config/config.<environment>.json
// file has been created and properly configured with server and database values. Of note, the '<environment>'
// can be set to any value, so long as the NOTECLERK_ENVIRONMENT environmental variable's value matches.
// RETURNS: *sql.db, error
func (d *DbPostgres) Initialize(config *Config) error {

	connStr := generateConnStrFromCfg(config)

	var err error
	if d.db, err = sql.Open("postgres", connStr); err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresInitializeFailsOpenConn)
	}
	defer d.db.Close()

	if err = d.db.Ping(); err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresInitializeFailsDbPing)
	}

	if schemaErr := d.createSchema(); schemaErr != nil {
		return NoteClerkErrWrap(schemaErr, ErrDbPostgresInitializeFailsSchemaCreation)
	}

	return nil
}

func (d *DbPostgres) GetNoteFragmentsByNoteGuid(noteGuid string) ([]*ehrpb.NoteFragment, error) {
	rows, err := d.db.Query(getNoteFragmentByNoteGuidQuery, noteGuid)
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteFragmentsByNoteGuidFailsQuery)
	}
	defer rows.Close()

	noteFragments := make([]*ehrpb.NoteFragment, 0)
	for rows.Next() {
		tmp := noted.NewNoteFragment()
		if err := rows.Scan(&tmp.Id, &tmp.DateCreated.Seconds, &tmp.DateCreated.Nanos, &tmp.NoteFragmentGuid,
			&tmp.NoteGuid, &tmp.Icd_10Code, &tmp.Icd_10Long, &tmp.Description, &tmp.Status,
			&tmp.Priority, &tmp.Topic, &tmp.Content); err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteFragmentsByNoteGuidFailsScan)
		}
		if tmp.Tags, err = d.GetNoteFragmentTagsByNoteFragmentGuid(tmp.GetNoteFragmentGuid()); err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteFragmentsByNoteGuidFailsGetTags)
		}
		noteFragments = append(noteFragments, tmp)
	}
	return noteFragments, nil
}

func (d *DbPostgres) GetNoteTagsByNoteGuid(noteGuid string) (tag []string, err error) {
	rows, err := d.db.Query(getNoteTagByNoteGuidQuery, noteGuid)
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteTagsByNoteGuidQueryFails)
	}
	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var id int64
		var noteGuid string
		var tmpString string
		err := rows.Scan(&id, &noteGuid, &tmpString)
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteTagsByNoteGuidFailsRowScan)
		}
		tags = append(tags, tmpString)
	}
	return tags, nil
}

func (d *DbPostgres) GetNoteFragmentTagsByNoteFragmentGuid(noteFragGuid string) (tag []string, err error) {
	rows, err := d.db.Query(getNoteFragmentTagsByNoteFragmentGuidQuery, noteFragGuid)
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteFragTagByNoteGuidQueryFails)
	}
	defer rows.Close()

	tags := make([]string, 0)
	for rows.Next() {
		var tmpId int64
		var tmpNoteFragmentGuid string
		var tmpTag string
		err := rows.Scan(&tmpId, &tmpNoteFragmentGuid, &tmpTag)
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteFragTagByNoteGuidFailsRowScan)
		}
		tags = append(tags, tmpTag)
	}
	return tags, nil
}

func (d *DbPostgres) AddNote(n *ehrpb.Note) (id int64, guid string, err error) {

	row := d.db.QueryRow(addNoteQuery, n.DateCreated.GetSeconds(), n.DateCreated.GetNanos(),
		n.GetNoteGuid(), n.GetVisitGuid(), n.GetAuthorGuid(), n.GetPatientGuid(), n.GetType(),
		n.GetStatus())

	if err := row.Scan(&n.Id); err != nil {
		return 0, "", NoteClerkErrWrap(err, ErrDbPostgresAddNoteFailsScan)
	}

	for _, v := range n.GetTags() {
		_, err := d.AddNoteTag(n.GetNoteGuid(), v)

		if err != nil {
			return 0, "", NoteClerkErrWrap(err, ErrDbPostgresAddNoteFailsToAddNoteTags)
		}
	}

	for _, v := range n.GetFragments() {
		_, _, err := d.AddNoteFragment(v)

		if err != nil {
			return 0, "", NoteClerkErrWrap(err, ErrDbPostgresAddNoteFailsToAddNoteFragments)
		}
	}

	return n.GetId(), n.GetNoteGuid(), nil
}

func (d *DbPostgres) UpdateNote(n *ehrpb.Note) error {
	err := d.DeleteNote(n.GetNoteGuid())
	if err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresUpdateNoteFailsToChangeStatusToDeleted)
	}
	n.NoteGuid = uuid.New().String()
	for _, v := range n.GetFragments() {
		v.NoteFragmentGuid = uuid.New().String()
		v.NoteGuid = n.NoteGuid
	}
	_, _, err = d.AddNote(n)
	if err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresUpdateNoteFailsToChangeStatusToDeleted)
	}
	return nil
}

func (d *DbPostgres) DeleteNote(guid string) error {
	note, getNoteErr := d.GetNoteByGuid(guid)
	if getNoteErr != nil {
		return getNoteErr
	}
	for _, v := range note.GetFragments() {
		delErr := d.DeleteNoteFragment(v.GetNoteFragmentGuid())
		if delErr != nil {
			return delErr
		}
	}
	row := d.db.QueryRow(updateNoteStatusToStatusByNoteGuidQuery, ehrpb.RecordStatus_DELETED, guid)
	var newId int64
	scanErr := row.Scan(&newId)
	if scanErr != nil {
		return scanErr
	}
	return nil
}

func (d *DbPostgres) AllNotes() ([]*ehrpb.Note, error) {
	rows, err := d.db.Query(getAllNotesQuery)
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresAllNotesFailsQuery)
	}
	defer rows.Close()

	var notes []*ehrpb.Note
	for rows.Next() {
		tmpNote := noted.NewNote()
		err := rows.Scan(&tmpNote.Id, &tmpNote.DateCreated.Seconds, &tmpNote.DateCreated.Nanos,
			&tmpNote.NoteGuid, &tmpNote.VisitGuid, &tmpNote.AuthorGuid,
			&tmpNote.PatientGuid, &tmpNote.Type, &tmpNote.Status)
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresAllNotesFailsScan)
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
	if err := row.Scan(&newId); err != nil {
		return 0, NoteClerkErrWrap(err, ErrDbPostgresAddNoteTagFailsScan)
	}

	return newId, nil
}

func (d *DbPostgres) GetNoteByGuid(guid string) (*ehrpb.Note, error) {
	row := d.db.QueryRow(getNoteByGuidQuery, guid)

	newNote := noted.NewNote()
	err := row.Scan(&newNote.Id, &newNote.DateCreated.Seconds, &newNote.DateCreated.Nanos, &newNote.NoteGuid,
		&newNote.VisitGuid, &newNote.AuthorGuid, &newNote.PatientGuid, &newNote.Type, &newNote.Status)

	if err != nil {
		return nil, err
	}
	newNote.Fragments, err = d.GetNoteFragmentsByNoteGuid(newNote.GetNoteGuid())
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteByGuidFailsGetNote)
	}

	newNote.Tags, err = d.GetNoteTagsByNoteGuid(newNote.GetNoteGuid())
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresGetNoteByGuidFailsGetNoteFragments)
	}

	return newNote, nil
}

// TODO: Include note content and tags in search
func (d *DbPostgres) FindNotes(filter NoteFindFilter) ([]*ehrpb.Note, error) {

	notes := make([]*ehrpb.Note, 0)

	if filter.SearchTerms != "" {
		return notes, errors.New("this find notes by search terms feature is not  yet implemented.")
	}

	if err := validateNoteFormFilterFields(filter); err != nil {
		return notes, err
	}
	transEmptyFieldToWildcard(&filter)

	rows, err := d.db.Query(getNotesByFindQuery, filter.AuthorGuid, filter.VisitGuid, filter.PatientGuid)
	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresFindNotesFailsQuery)
	}
	defer rows.Close()

	for rows.Next() {
		tmpNote := noted.NewNote()
		err := rows.Scan(&tmpNote.Id, &tmpNote.DateCreated.Seconds, &tmpNote.DateCreated.Nanos,
			&tmpNote.NoteGuid, &tmpNote.VisitGuid, &tmpNote.AuthorGuid,
			&tmpNote.PatientGuid, &tmpNote.Type, &tmpNote.Status)
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresFindNotesFailsScan)
		}
		tmpNote.Tags, err = d.GetNoteTagsByNoteGuid(tmpNote.GetNoteGuid())
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresFindNotesFailsGetTags)
		}

		tmpNote.Fragments, err = d.GetNoteFragmentsByNoteGuid(tmpNote.GetNoteGuid())
		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresFindNotesFailsGetNoteFragments)
		}

		notes = append(notes, tmpNote)

	}
	return notes, nil
}

func transEmptyFieldToWildcard(filter *NoteFindFilter) {
	if filter.AuthorGuid == "" {
		filter.AuthorGuid = "%"
	}
	if filter.VisitGuid == "" {
		filter.VisitGuid = "%"
	}
	if filter.PatientGuid == "" {
		filter.PatientGuid = "%"
	}
}

func validateNoteFormFilterFields(queryFilter NoteFindFilter) error {
	_, err := uuid.Parse(queryFilter.VisitGuid)
	if err != nil && queryFilter.VisitGuid != "" {
		return err
	}
	_, err = uuid.Parse(queryFilter.PatientGuid)
	if err != nil && queryFilter.PatientGuid != "" {
		return err
	}
	_, err = uuid.Parse(queryFilter.AuthorGuid)
	if err != nil && queryFilter.AuthorGuid != "" {
		return err
	}
	return nil
}

func (d *DbPostgres) AllNoteFragments() ([]*ehrpb.NoteFragment, error) {
	rows, err := d.db.Query(getAllNoteFragmentsQuery)

	if err != nil {
		return nil, NoteClerkErrWrap(err, ErrDbPostgresAllNoteFragmentsQueryFails)
	}
	defer rows.Close()

	var notes []*ehrpb.NoteFragment
	for rows.Next() {
		tmpFrag := noted.NewNoteFragment()
		err := rows.Scan(&tmpFrag.Id, &tmpFrag.DateCreated.Seconds, &tmpFrag.DateCreated.Nanos,
			&tmpFrag.NoteFragmentGuid, &tmpFrag.NoteGuid, &tmpFrag.Icd_10Code, &tmpFrag.Icd_10Long,
			&tmpFrag.Description, &tmpFrag.Status, &tmpFrag.Priority, &tmpFrag.Topic, &tmpFrag.Content)

		if err != nil {
			return nil, NoteClerkErrWrap(err, ErrDbPostgresAllNoteFragmentsFailsScanRow)
		}

		tmpFrag.Tags, err = d.GetNoteFragmentTagsByNoteFragmentGuid(tmpFrag.GetNoteGuid())
		if err != nil {
			return nil, err
		}

		notes = append(notes, tmpFrag)

	}
	return notes, nil
}

func (d *DbPostgres) AddNoteFragment(nf *ehrpb.NoteFragment) (id int64, guid string, err error) {
	row := d.db.QueryRow(addNoteFragmentQuery, nf.DateCreated.Seconds, nf.DateCreated.Nanos,
		nf.GetNoteFragmentGuid(), nf.GetNoteGuid(), nf.GetIcd_10Code(), nf.GetIcd_10Long(),
		nf.GetDescription(), nf.GetStatus(), nf.GetPriority(), nf.GetTopic(), nf.GetContent())
	scanErr := row.Scan(&nf.Id)
	if scanErr != nil {
		return 0, "", NoteClerkErrWrap(scanErr, ErrDbPostgresAddNoteFragmentFailsScan)
	}

	for _, v := range nf.GetTags() {
		_, err := d.AddNoteFragmentTag(nf.GetNoteFragmentGuid(), v)
		if err != nil {
			return 0, "", NoteClerkErrWrap(err, ErrDbPostgresAddNoteFragmentFailsAddNoteTags)
		}
	}

	return nf.GetId(), nf.GetNoteFragmentGuid(), nil
}

func (d *DbPostgres) UpdateNoteFragment(n *ehrpb.NoteFragment) error {

	newFrag := buildNewFragmentFromOldFragment(n)

	_, _, err := d.AddNoteFragment(newFrag)
	if err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresUpdateNoteFragmentFailsAddNewNoteFragment)
	}

	err = d.DeleteNoteFragment(n.GetNoteFragmentGuid())
	if err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresUpdateNoteFragmentFailsDeletePriorNoteFragment)
	}

	return nil
}

func buildNewFragmentFromOldFragment(n *ehrpb.NoteFragment) *ehrpb.NoteFragment {
	newFrag := noted.NewNoteFragment()
	newFrag.Content = n.GetContent()
	newFrag.Topic = n.GetTopic()
	newFrag.Priority = n.GetPriority()
	newFrag.Status = n.GetStatus()
	newFrag.Description = n.GetDescription()
	newFrag.Icd_10Long = n.GetIcd_10Long()
	newFrag.Icd_10Code = n.GetIcd_10Code()
	newFrag.IssueGuid = n.GetIssueGuid()
	newFrag.Tags = n.GetTags()
	newFrag.NoteGuid = n.GetNoteGuid()
	return newFrag
}

// This is not a true delete. It changes the status of the note to DELETED. Health care
// records should not be deleted.
func (d *DbPostgres) DeleteNoteFragment(noteFragmentGuid string) error {
	row := d.db.QueryRow(updateNoteFragmentStatusToStatusByNoteFragmentGuidQuery, ehrpb.RecordStatus_DELETED, noteFragmentGuid)
	var newId int64
	scanErr := row.Scan(&newId)
	if scanErr != nil {
		return scanErr
	}
	return nil
}

//TODO: Implement feature
func (d *DbPostgres) GetNoteFragmentByGuid(guid string) (*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

//TODO: Implement feature
func (d *DbPostgres) FindNoteFragments(filter NoteFragmentFindFilter) ([]*ehrpb.NoteFragment, error) {
	log.Fatal("Not implemented.")
	return nil, nil
}

func (d *DbPostgres) AddNoteFragmentTag(noteGuid string, tag string) (id int64, err error) {
	row := d.db.QueryRow(addNoteFragmentTagQuery, noteGuid, tag)

	var newId int64
	if err := row.Scan(&newId); err != nil {
		return 0, NoteClerkErrWrap(err, ErrDbPostgresAddNoteFragmentTagFailsScan)
	}

	return newId, nil
}

func generateConnStrFromCfg(config *Config) string {
	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		config.DbUsername, config.DbPassword, config.DbIp, config.DbName, config.DbSslMode, config.DbPort)
	return connStr
}

// https://www.calhoun.io/updating-and-deleting-postgresql-records-using-gos-sql-package/
func (d *DbPostgres) createSchema() error {

	err := d.createTable(createNoteTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(err, "%v. Target Table: note", ErrDbPostgresCreateSchemaFailsTableCreation)
	}

	err = d.createTable(createNoteTagTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(err, "%v. Target Table: note_tag", ErrDbPostgresCreateSchemaFailsTableCreation)
	}

	err = d.createTable(createNoteFragmentTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(err, "%v. Target Table: note_fragment", ErrDbPostgresCreateSchemaFailsTableCreation)
	}

	err = d.createTable(createNoteFragmentTagTable)
	if notNilNotTableExists(err) {
		return errors.Wrapf(err, "%v. Target Table: note_fragment_tag", ErrDbPostgresCreateSchemaFailsTableCreation)
	}

	return nil
}

func (d *DbPostgres) createTable(query string) error {
	_, err := d.db.Exec(query)

	tableExistsError := strings.Contains(fmt.Sprintf("%v", err), "already exists")
	if tableExistsError {
		return NoteClerkErrWrap(err, ErrDbPostgresCreateTableFailsAlreadyExists)
	}
	if err != nil {
		return NoteClerkErrWrap(err, ErrDbPostgresCreateTableFailsDueToUnexpectedError)
	}
	return nil
}

func notNilNotTableExists(err error) bool {
	errMsg := fmt.Sprintf("%v", err)
	alreadyExistsError := strings.Contains(errMsg, "relation") && strings.Contains(errMsg, "already exists")
	return err != nil && !alreadyExistsError
}
