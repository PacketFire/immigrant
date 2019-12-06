package core

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	revisionBuffSize int = 10 // just an arbitrary number
)

// Walk takes a path and will traverse the entire directory, dispatching the
// files to the parseRevisions function. If any fatal errors are returned an
// the returned error will be non-nil. On success a *Revisions is returned.
// on failure a *Revisions and error is returend
func Walk(path string) (*Revisions, error) {
	// parse channel for stream of revisions.
	pc := make(chan Revision)

	// map channel for syncing map back to curren routine.
	mc := make(chan map[string]Revision)

	rpath := filepath.Join(path, "revisions")

	// Provide a goroutine to protect a channel to consolidate Revisions.
	// In this source file, this will be referred to as the sync routine.
	go func(cmc chan map[string]Revision, cpc chan Revision) {
		rm := make(map[string]Revision)

		for rev := range cpc {
			id := rev.Revision
			if _, prs := rm[id]; prs == true {
				panic(fmt.Sprintf("duplicate revision ID: %s", id))
			}

			rm[id] = rev
		}

		cmc <- rm
	}(mc, pc)

	if err := filepath.Walk(rpath, parseRevisions(pc)); err != nil {
		return &Revisions{}, errors.New("directory traversal failed")
	}

	close(pc)
	revisions := <-mc

	return &Revisions{revisions: revisions}, nil
}

// parseRevisions will take a channel as an argument and will inject the
// channel into a returned closure that satisfies the WalkFunc prototype. This
// will allow the walk to safely push parsed revisions back to the calling
// method.
func parseRevisions(c chan Revision) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		var rb []Revision

		// Catch any errors passed from Walk.
		if err != nil {
			return err
		}

		// Catch Directories
		if info.IsDir() {
			return nil
		}

		// verify that the file is yaml
		if strings.HasSuffix(path, ".yml") || strings.HasSuffix(path, ".yaml") {
			yml, e := ioutil.ReadFile(path)
			if e != nil {
				return e
			}

			rb, err = unmarshalYamlRevisions(yml)
			if err != nil {
				return fmt.Errorf("unable to unmarshal %s", path)
			}
		}

		// Push each revision to to the sync routine.
		for _, rev := range rb {
			c <- rev
		}

		return nil
	}
}

func unmarshalYamlRevisions(yml []byte) ([]Revision, error) {
	rb := make([]Revision, 0, revisionBuffSize)
	if err := yaml.Unmarshal(yml, &rb); err != nil {
		return rb, errors.New("unable to unmarshal file")
	}

	return rb, nil
}
