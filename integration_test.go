package main

import (
	"testing"
	"fmt"
	"database/sql"
)

var postgresDb = &DbPostgres{}

const testingEnv = "testing"

func TestDbPostgres_AddNote(t *testing.T) {
	setup(t)
	note := NewNote()
	id, err := postgresDb.AddNote(note)
	if err != nil {
		t.Fatalf("Failed to add note to datbase. Error: %v", err)
	}
	if id <= 0 {
		t.Fatalf("Expected an Id greater than zero, got %v", id)
	}
	teardown()
}

func setup(t *testing.T) {
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

func teardown() {
	postgresDb.db.Close()
}