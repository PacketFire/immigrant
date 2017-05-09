package main

import (
	"testing"
)

const (
	revisionTestDir string = "test/revisions/"
)

func TestWalk(t *testing.T) {
	_, err := Walk(revisionTestDir)
	if err != nil {
		t.Fatal(err)
	}
}
