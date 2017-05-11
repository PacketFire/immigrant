package main

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
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

	rp, err := revisionFileName(path)
	if err != nil {

		return seq, err
	}

	raw, err := ioutil.ReadFile(rp)
	if err != nil {
		return seq, errors.New(fmt.Sprintf("Unabled to open sequence: %s", rp))
	}

	seq, err = unmarshalYamlSequence(raw)
	if err != nil {
		return seq, errors.New("Unable to unmarshal sequence.")
	}

	return seq, nil
}

// revisionFileName takes a string representing the config directory as an arg
// and attempts to resolve the filename of the revision sequence file. On
// success, a string path to the file and nil is returned. On failure, an empty
// string and the error is returned.
func revisionFileName(path string) (string, error) {
	var rp string
	ymlp := filepath.Join(path, "revision.yml")
	yamlp := filepath.Join(path, "revision.yaml")

	if _, err := os.Stat(ymlp); err == nil {
		return ymlp, nil
	} else if _, err := os.Stat(yamlp); err == nil {
		return yamlp, nil
	}

	return rp, errors.New("Sequence does not exist in config directory.")
}

// unmarshalYamlSequence takes the raw yaml []byte and unmarshals it to a
// *Sequence. On success, a *Sequence and an error is returned. on failure, a
// *Sequence and an error is returned.
func unmarshalYamlSequence(yml []byte) (*Sequence, error) {
	seq := new(Sequence)
	if err := yaml.Unmarshal(yml, seq); err != nil {
		return seq, errors.New("Unable to unmarshal file")
	}

	return seq, nil
}
