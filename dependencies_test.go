package main

import (
	"testing"
)

func TestInitializeLogger_WithInvalidPath_CallsFatalLog(t *testing.T) {
	if err := InitializeLogger(""); err == nil {
		t.Fatalf("Expected an error, but got nil. Should not be able to open a file that does not have a valid path.")
	}
}
