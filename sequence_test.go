package main

import (
	"fmt"
	"testing"
)

const (
	YAMLRevisionsTestDir string = "test/parse_test_dir_yaml_revisions"
)

func TestParseSequence(t *testing.T) {
	seq, err := ParseSequence(revisionTestDir)
	fmt.Println(seq)
	if err != nil {
		t.Error(err)
	}

	if len(seq.Revisions) != 2 {
		t.Errorf("Sequence fails to parse all revisions")
	}
}

func TestYamlFileParseSequence(t *testing.T) {
	seq, err := ParseSequence(revisionTestDir)
	fmt.Println(seq)
	if err != nil {
		t.Error(err)
	}

	if len(seq.Revisions) != 2 {
		t.Errorf("Sequence fails to parse all revisions")
	}
}
