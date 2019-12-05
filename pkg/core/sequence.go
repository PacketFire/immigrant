package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Sequence contains a list of strings representing, and mapping directly to,
// Revision IDs. This sequence is uses to determine an order for which
// revisions are applied or rolled back from.
type Sequence struct {
	Revisions []string `yaml:"revisions"`
}

// ParseSequence will take a path to the config directory and attempt to open
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
		return seq, fmt.Errorf("unabled to open sequence: %s", rp)
	}

	seq, err = unmarshalYamlSequence(raw)
	if err != nil {
		return seq, errors.New("unable to unmarshal sequence")
	}

	return seq, nil
}

// revisionFileName takes a string representing the config directory as an arg
// and attempts to resolve the filename of the revision sequence file. On
// success, a string path to the file and nil is returned. On failure, an empty
// string and the error is returned.
func revisionFileName(path string) (string, error) {
	var rp string
	ymlp := filepath.Join(path, "revisions.yml")
	yamlp := filepath.Join(path, "revisions.yaml")

	if _, err := os.Stat(ymlp); err == nil {
		return ymlp, nil
	} else if _, err := os.Stat(yamlp); err == nil {
		return yamlp, nil
	}

	return rp, errors.New("sequence does not exist in config directory")
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
