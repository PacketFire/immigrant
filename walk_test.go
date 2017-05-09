package main

import (
	"os"
	"testing"
)

const (
	revisionTestDir          string = "test/revisions"
	emptyRevisionTestDir     string = "test/empty_revisions_dir"
	duplicateRevisionTestDir string = "test/duplicate_revisions"
	fileExtTestDir           string = "test/file_ext_test"
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
	os.Mkdir(emptyRevisionTestDir, 0755)
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

func TestDuplicateRevisions(t *testing.T) {
	Walk(duplicateRevisionTestDir)
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
