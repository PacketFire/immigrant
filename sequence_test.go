package main

import (
	"testing"
)

func TestParseSequence(t *testing.T) {
	seq, err := ParseSequence(revisionTestDir)
	if err != nil {
		t.Error(err)
	}

	if len(seq.Revisions) != 2 {
		t.Errorf("Sequence fails to parse all revisions")
	}
}

func TestYamlFileParseSequence(t *testing.T) {
	seq, err := ParseSequence(revisionTestDir)
	if err != nil {
		t.Error(err)
	}

	if len(seq.Revisions) != 2 {
		t.Errorf("Sequence fails to parse all revisions")
	}
}
