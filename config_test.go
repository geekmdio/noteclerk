package main

import "testing"

func TestLoadConfiguration_WherePathDoesNotPointToConfig_ReturnsError(t *testing.T) {
	_, err := LoadConfiguration("")
	if err == nil {
		t.Fatalf("Should throw error when no config file present.")
	}
}
