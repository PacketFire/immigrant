package main

import (
	"testing"
)

const (
	revisionTestDir      string = "test/revisions/"
	emptyRevisionTestDir string = "test/empty_revisions_dir/"
)

func TestWalk(t *testing.T) {
	_, err := Walk(revisionTestDir)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmptyWalk(t *testing.T) {
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
