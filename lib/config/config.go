package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ParseConfig takes a path to the config directory and attempts to parse the
// config.yml file in that directory. On success a map[string]string is
// returned. On failure a map[string]string and an error is returned.
func ParseConfig(path string) (map[string]string, error) {
	c := make(map[string]string)

	cp, err := configFileName(path)
	if err != nil {
		return c, err
	}

	raw, err := ioutil.ReadFile(cp)
	if err != nil {
		return c, errors.New(fmt.Sprintf("Unabled to open config file: %s", cp))
	}

	c, err = unmarshalYamlConfig(raw)
	if err != nil {
		return c, errors.New("Unable to unmarshal config.")
	}

	return c, nil
}

// configFileName takes a string representing the config directory as an arg
// and attempts to resolve the filename of the revision sequence file. On
// success, a string path to the file and nil is returned. On failure, an empty
// string and the error is returned.
func configFileName(path string) (string, error) {
	var cp string
	ymlp := filepath.Join(path, "config.yml")
	yamlp := filepath.Join(path, "config.yaml")

	if _, err := os.Stat(ymlp); err == nil {
		return ymlp, nil
	} else if _, err := os.Stat(yamlp); err == nil {
		return yamlp, nil
	}

	return cp, errors.New("Config file does not exist in config directory.")
}

// unmarshalYamlConfig takes the raw yaml []byte and unmarshals it to a
// map[string]string. On success a map[string]string and an error are returned.
// on failure, a map[string]string and error is returned.
func unmarshalYamlConfig(yml []byte) (map[string]string, error) {
	c := make(map[string]string)
	if err := yaml.Unmarshal(yml, c); err != nil {
		return c, errors.New("Unable to unmarshal config file")
	}

	return c, nil
}
