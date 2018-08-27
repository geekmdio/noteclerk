package main

import (
	"testing"
	"fmt"
	"database/sql"
	"github.com/geekmdio/noted"
	"github.com/google/uuid"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
	"github.com/sirupsen/logrus"
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
		t.Fatalf("Create table should return an error. Actual error: %v", err)
	}

	postgresDb.db.Exec(dropQuery)

	tearDown(t)
}


func TestDbPostgres_AddNote(t *testing.T) {
	setup(t)
	note := buildNote()

	id, err := postgresDb.AddNote(note)
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

	id, _ := postgresDb.AddNote(note)
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

	id, _ := postgresDb.AddNote(note)
	err := postgresDb.DeleteNote(id)
	if err != nil {
		t.Fatalf("Failed to delete note. Error: %v", err)
	}
	tearDown(t)
}

func TestDbPostgres_FindNotes(t *testing.T) {
	setup(t)

	note := buildNote()
	postgresDb.AddNote(note)

	visitGuid := "" // note.GetVisitGuid()
	patientGuid := "" // note.GetPatientGuid()
	authorGuid := note.GetAuthorGuid()

	findQuery := NoteFindFilter{
		VisitGuid:   visitGuid,
		AuthorGuid:  authorGuid,
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

func buildNote() *ehrpb.Note {
	nb := &noted.NoteBuilder{}
	note := nb.Init().
		SetId(0).
		SetDateCreated(TimestampNow()).
		SetPatientGuid(uuid.New().String()).
		SetAuthorGuid(uuid.New().String()).
		SetVisitGuid(uuid.New().String()).
		SetType(ehrpb.NoteType_HISTORY_AND_PHYSICAL).
		Build()
	note.Tags = append(note.Tags, "tag1", "tag2")
	fb := &noted.NoteFragmentBuilder{}
	frag := fb.InitFromNote(note).
		SetId(0).
		SetDateCreated(TimestampNow()).
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

	// Below here ensures that the database and it's tables are setup for the integration tests.
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