package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWalk(t *testing.T) {
	rb, err := Walk(revisionTestDir)
	if err != nil {
		t.Fatal(err)
	}

	if len(rb) != 3 {
		t.Fatal("Failed to parse files.")
	}
}

func TestEmptyWalk(t *testing.T) {
	// Create dependent directory
	td := filepath.Join(emptyRevisionTestDir, "revisions")
	os.Mkdir(emptyRevisionTestDir, 0755)
	os.Mkdir(td, 0755)
	_, err := Walk(emptyRevisionTestDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWalkNonYamlFile(t *testing.T) {
	revisions, err := Walk(revisionTestDir)
	if err != nil {
		t.Fatal(err)
	}

	if _, p := revisions["revision_txt"]; p == true {
		t.Fatal("Walk should not parse non yaml files.")
	}
}

func TestYAMLFileExtensions(t *testing.T) {
	rb, err := Walk(fileExtTestDir)
	if err != nil {
		t.Fatal(err)
	}

	if len(rb) != 2 {
		t.Fatal("Failing to open both yaml file extensions.")
	}
}
