package main

import (
	"testing"
	"fmt"
	"database/sql"
	"github.com/geekmdio/noted"
	"github.com/google/uuid"
	"github.com/geekmdio/ehrprotorepo/v1/generated/goproto"
)

var postgresDb = &DbPostgres{}

const testingEnv = "testing"

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
	teardown(t)
}

func buildNote() *ehrpb.Note {
	nb := &noted.NoteBuilder{}
	note := nb.Init().
		SetPatientGuid(uuid.New().String()).
		SetAuthorGuid(uuid.New().String()).
		SetVisitGuid(uuid.New().String()).
		SetId(0).
		SetDateCreated(TimestampNow()).
		SetType(ehrpb.NoteType_HISTORY_AND_PHYSICAL).
		Build()
	note.Tags = append(note.Tags, "tag1", "tag2")
	fb := &noted.NoteFragmentBuilder{}
	frag := fb.InitFromNote(note).
		SetDateCreated(TimestampNow()).
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
	//TODO: Switch to this for config; build into .travis.yml a docker postgres db
	// https://docs.travis-ci.com/user/docker/
	cfg := &Config{
		Version:        "testing",
		LogPath:        "",
		ServerProtocol: "tcp",
		ServerIp:       "localhost",
		ServerPort:     "50051",
		DbIp:           "localhost",
		DbPort:         "5433",
		DbUsername:     "integration",
		DbPassword:     "testing",
		DbName:         "noteclerk",
		DbSslMode:      "disable",
	}

	cfg, err := LoadConfiguration(fmt.Sprintf("config/config.%v.json", testingEnv))
	if err != nil {
		t.Fatalf("Could not log testing configuration to perform integration test.")
	}

	connStr := fmt.Sprintf("user=%v password=%v host=%v dbname=%v sslmode=%v port=%v",
		cfg.DbUsername, cfg.DbPassword, cfg.DbIp, cfg.DbName, cfg.DbSslMode, cfg.DbPort)

	open, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatalf("Failed to open database connection.")
	}

	postgresDb.db = open

}

func teardown(t *testing.T) {
	err := postgresDb.db.Close()
	if err != nil {
		t.Fatalf("Failed to tear down integration testing by closing database.")
	}
}