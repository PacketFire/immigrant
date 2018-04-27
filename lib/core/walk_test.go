package core

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

	if rb.Length() != 3 {
		t.Fatal("Failing to parse revisions.")
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
	rev, err := Walk(revisionTestDir)
	if err != nil {
		t.Fatal(err)
	}

	if _, p := rev.Revision("revision_txt"); p == true {
		t.Fatal("Walk should not parse non yaml files.")
	}
}

func TestYAMLFileExtensions(t *testing.T) {
	rb, err := Walk(fileExtTestDir)
	if err != nil {
		t.Fatal(err)
	}

	if rb.Length() != 2 {
		t.Fatal("Failing to open both yaml file extensions.")
	}
}
