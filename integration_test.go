package main

import (
	"database/sql"
	"fmt"
	"github.com/beevik/guid"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/geekmdio/noted"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"testing"
)

var postgresDb = &DbPostgres{}

func TestCreateTable_ReturnsError_WithImproperQuery(t *testing.T) {
	setup(t)

	if err := postgresDb.createTable("CREATE BLAH"); err == nil {
		t.Fatalf("Should not be able to create table from improper query.")
	}

	tearDown(t)
}

func TestCreateTable_WhichDoesNotExist_ReturnsNil(t *testing.T) {
	setup(t)

	dropQuery := `DROP TABLE TestCreateTable_WhichDoesNotExist_ReturnsNil;`
	postgresDb.db.Exec(dropQuery)

	createQuery := `CREATE TABLE TestCreateTable_WhichDoesNotExist_ReturnsNil (id serial NOT NULL);`
	err := postgresDb.createTable(createQuery)
	if err != nil {
		t.Fatalf("Create table should return nil, but returned error '%v'", err)
	}

	postgresDb.db.Exec(dropQuery)

	tearDown(t)
}

func TestCreateTable_WhichAlreadyExists_ReturnsProperError(t *testing.T) {
	setup(t)

	dropQuery := `DROP TABLE create_table_returns_error;`
	postgresDb.db.Exec(dropQuery)

	createQuery := `CREATE TABLE create_table_returns_error (id serial NOT NULL);`

	if err := postgresDb.createTable(createQuery); err != nil {
		t.Fatalf("Create table should not return error when given proper query. Error: %v", err)
	}

	if err := postgresDb.createTable(createQuery); err == nil {
		t.Fatalf("Create table should return an error indicating the table already exists. Actual error: %v", err)
	}

	postgresDb.db.Exec(dropQuery)

	tearDown(t)
}

func TestDbPostgres_AllNotes(t *testing.T) {
	setup(t)
	note1 := buildNote()
	note2 := buildNote()
	note3 := buildNote()

	var ns []*ehrpb.Note
	ns = append(ns, note1, note2, note3)

	for _, n := range ns {
		postgresDb.AddNote(n)
	}


	retrievedNotes, err := postgresDb.AllNotes()

	if err != nil {
		t.Fatalf("Failed to retrieve all retrievedNotes from database. Error: %v", err)
	}

	hitCount := 0
	for _, rn := range retrievedNotes {
		for _, n := range ns {
			if rn.GetNoteGuid() == n.GetNoteGuid() {
				hitCount++
			}
		}
	}
	if hitCount != len(ns) {
		t.Fatalf("Failed to find all notes. Expected %v but got %v", len(ns), hitCount)
	}

}

func TestDbPostgres_AddNote(t *testing.T) {
	setup(t)
	note := buildNote()

	id, _, err := postgresDb.AddNote(note)
	if err != nil {
		t.Fatalf("Failed to add note to datbase. Error: %v", err)
	}
	if id <= 0 {
		t.Fatalf("Expected an Id greater than zero, got %v", id)
	}
	tearDown(t)
}

func TestDbPostgres_UpdateNote(t *testing.T) {
	setup(t)
	note := buildNote()

	id, _, _ := postgresDb.AddNote(note)
	note.Id = id
	note.Fragments[0].Content = "Updated content"

	err := postgresDb.UpdateNote(note)
	if err != nil {
		t.Fatalf("Failed to add note to datbase. Error: %v", err)
	}
	tearDown(t)
}

func TestDbPostgres_DeleteNote(t *testing.T) {
	setup(t)
	note := buildNote()

	postgresDb.AddNote(note)
	err := postgresDb.DeleteNote(note.GetNoteGuid())
	if err != nil {
		t.Fatalf("Failed to delete note. Error: %v", err)
	}
	tearDown(t)
}

func TestDbPostgres_FindNotes_ByAuthorGuid(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	authorGuid := note.GetAuthorGuid()

	findQuery := NoteFindFilter{
		VisitGuid:   "",
		AuthorGuid:  authorGuid,
		PatientGuid: "",
		SearchTerms: "",
	}

	notes, err := postgresDb.FindNotes(findQuery)
	if err != nil {
		t.Fatalf("Failed to find notes. Error: %v", err)
	}

	if len(notes) < 1 {
		t.Fatalf("Should return at least one note, but did not.")
	}
	tearDown(t)
}

func TestDbPostgres_FindNotes_ByVisitGuid(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	visitGuid := note.GetVisitGuid()

	findQuery := NoteFindFilter{
		VisitGuid:   visitGuid,
		AuthorGuid:  "",
		PatientGuid: "",
		SearchTerms: "",
	}

	notes, err := postgresDb.FindNotes(findQuery)
	if err != nil {
		t.Fatalf("Failed to find notes. Error: %v", err)
	}

	if len(notes) < 1 {
		t.Fatalf("Should return at least one note, but did not.")
	}
	tearDown(t)
}

func TestDbPostgres_FindNotes_ByPatientGuid(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	patientGuid := note.GetPatientGuid()

	findQuery := NoteFindFilter{
		VisitGuid:   "",
		AuthorGuid:  "",
		PatientGuid: patientGuid,
		SearchTerms: "",
	}

	notes, err := postgresDb.FindNotes(findQuery)
	if err != nil {
		t.Fatalf("Failed to find notes. Error: %v", err)
	}

	if len(notes) < 1 {
		t.Fatalf("Should return at least one note, but did not.")
	}
	tearDown(t)
}

func TestDbPostgres_FindNotes_BySearchTermsFailsDueToNotImplemented(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	findQuery := NoteFindFilter{
		VisitGuid:   "",
		AuthorGuid:  "",
		PatientGuid: "",
		SearchTerms: "foo bar fizz buzz",
	}

	notes, err := postgresDb.FindNotes(findQuery)
	if err == nil {
		t.Fatalf("The search by terms feature should fail, it is not implemented. The error should be nil. Error: %v", err)
	}

	if len(notes) >= 1 {
		t.Fatalf("Zero notes should be returned, or an error should be thrown. The search by terms feature is not yet implemented.")
	}
	tearDown(t)
}

func TestDbPostgres_AllNoteFragments(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	frags, err := postgresDb.AllNoteFragments()
	if err != nil {
		t.Fatal(err)
	}

	if len(frags) < 1 {
		t.Fatalf("Should have at least one note fragment.")
	}

	tearDown(t)
}

func TestDbPostgres_UpdateNoteFragment(t *testing.T) {
	setup(t)

	note := buildNote()

	frag := note.GetFragments()[0]

	newFrag := noted.NewNoteFragment()
	newFrag.NoteFragmentGuid = frag.GetNoteFragmentGuid()
	newFrag.NoteGuid = frag.GetNoteGuid()
	newFrag.IssueGuid = frag.GetIssueGuid()

	newFrag.Content = "This is an updated note fragment."

	postgresDb.AddNote(note)
	err := postgresDb.UpdateNoteFragment(newFrag)

	if err != nil {
		t.Fatalf("While attempting to update the note fragment, an error occured: %v", err)
	}

	tearDown(t)
}

func TestDbPostgres_UpdateNoteFragment_FailsWhenGuidNotExist(t *testing.T) {
	setup(t)

	note := buildNote()

	frag := note.GetFragments()[0]
	frag.NoteFragmentGuid = guid.New().String()

	newFrag := noted.NewNoteFragment()
	newFrag.NoteGuid = frag.GetNoteGuid()
	newFrag.IssueGuid = frag.GetIssueGuid()

	newFrag.Content = "This is an updated note fragment."

	postgresDb.AddNote(note)
	err := postgresDb.UpdateNoteFragment(newFrag)

	if err == nil {
		t.Fatalf("While attempting to update the note fragment, an error should have occured but did not. Error: %v", err)
	}

	tearDown(t)
}

func buildNote() *ehrpb.Note {
	nb := &noted.NoteBuilder{}
	note := nb.Init().
		SetId(0).
		SetPatientGuid(uuid.New().String()).
		SetAuthorGuid(uuid.New().String()).
		SetVisitGuid(uuid.New().String()).
		SetType(ehrpb.NoteType_HISTORY_AND_PHYSICAL).
		Build()
	note.Tags = append(note.Tags, "tag1", "tag2")
	fb := &noted.NoteFragmentBuilder{}
	frag := fb.InitFromNote(note).
		SetId(0).
		SetStatus(ehrpb.RecordStatus_ACTIVE).
		SetPriority(ehrpb.RecordPriority_HIGH).
		SetTopic(ehrpb.FragmentType_MEDICAL_HISTORY).
		SetIssueGuid(uuid.New().String()).
		SetIcd10Code("ICD10").
		SetIcd10LongDescription("Long description of ICD10").
		SetDescription("My Description").
		SetContent("This is the content").
		Build()
	frag.Tags = append(frag.Tags, "fragtag1", "fragtag2")
	note.Fragments = append(note.Fragments, frag)
	return note
}

func setup(t *testing.T) {
	// Don't clutter the integration tests with logging data
	log.SetLevel(logrus.FatalLevel)

	cfg := &Config{
		Version:        "under-development",
		LogPath:        "/dev/null",
		ServerProtocol: "tcp",
		ServerIp:       "localhost",
		ServerPort:     "50051",
		DbIp:           "localhost",
		DbPort:         "5434",
		DbUsername:     "integration",
		DbPassword:     "testing",
		DbName:         "noteclerk",
		DbSslMode:      "disable",
	}

	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		cfg.DbUsername, cfg.DbPassword, cfg.DbIp, cfg.DbName, cfg.DbSslMode, cfg.DbPort)

	openDb, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to open database connection.")
	}

	postgresDb.db = openDb

	// Below ensures that the database and it's tables are setup for the integration tests.
	server := &DbPostgres{
		db: openDb,
	}

	server.createSchema()

}

func tearDown(t *testing.T) {
	err := postgresDb.db.Close()
	if err != nil {
		t.Fatalf("Failed to tear down integration testing by closing database.")
	}
}
