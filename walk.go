package main

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
	RevisionBuffSize int = 10 //just an arbitrary number
)

// Walk takes a path and will traverse the entire directory, dispatching the
// files to the parseRevisions function. If any fatal errors are returned an
// the returned error will be non-nil. On success a populated map of Revision
// instances will be returened.
func Walk(path string) (map[string]Revision, error) {
	// parse channel for stream of revisions.
	pc := make(chan Revision)
	// map channel for syncing map back to curren routine.
	mc := make(chan map[string]Revision)

	// Provide a goroutine to protect a channel to consolidate Revisions.
	// In this source file, this will be referred to as the sync routine.
	go func(cmc chan map[string]Revision, cpc chan Revision) {
		rm := make(map[string]Revision)

		for rev := range cpc {
			rm[rev.Revision] = rev
		}

		cmc <- rm
	}(mc, pc)

	if err := filepath.Walk(path, parseRevisions(pc)); err != nil {
		return make(map[string]Revision), errors.New("Directory traversal failed")
	}

	close(pc)
	revisions := <-mc

	return revisions, nil
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
		if strings.HasSuffix(path, ".yml") {
			yml, e := ioutil.ReadFile(path)
			if e != nil {
				return e
			}

			rb, err = parseYamlRevisions(yml)
			if err != nil {
				return errors.New(fmt.Sprintf("Unable to unmarshal %s", path))
			}
		}

		// Push each revision to to the sync routine.
		for _, rev := range rb {
			c <- rev
		}

		return nil
	}
}

func parseYamlRevisions(yml []byte) ([]Revision, error) {
	rb := make([]Revision, 0, RevisionBuffSize)
	if err := yaml.Unmarshal(yml, &rb); err != nil {
		return rb, errors.New("Unable to unmarshal file")
	}

	return rb, nil
}
