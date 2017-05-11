package main

import (
  "os"
  "path/filepath"
)

type Sequence struct {
	Revisions []*Revision `yaml:"revisions"`
}

// Parse sequence will take a path to the config directory and attempt to open
// and parse the file to a sequence. On success a *Sequence and nil are
// returned. On failure a *Sequence and error are returned.
func ParseSequence(path string) (*Sequence, error) {
	seq := new(Sequence)

	return seq, nil
}

// revisionFileName takes a string representing the config directory as an arg
// and attempts to resolve the filename of the revision sequence file. On
// success, a string path to the file and nil is returned. On failure, an empty
// string and the error is returned.
func revisionFileName(path string) (string, error) {
  var rp string

  return rp, nil
}

// unmarshalYamlSequence takes the raw yaml string and unmarshals it to a
// *Sequence. On success, a *Sequence and an error is returned. on failure, a
// *Sequence and an error is returned.
func unmarshalYamlSequence(sf string) (*Sequence, error) {
  seq := new(Sequence)

  return seq, nil
}
