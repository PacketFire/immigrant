package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ErrNoTypeSpecified is triggered when a driver type lookup occurs with no
// sprecified type field.
type ErrNoTypeSpecified struct{}

func (e *ErrNoTypeSpecified) Error() string {
	return "driver type not specified"
}

// Config represents a KV mapping of string parameters to be passed to the
// database.
type Config map[string]string

// DriverType returns the string representation of the driver to load.
func (c Config) DriverType() (string, error) {
	if v, prs := c["type"]; prs == true {
		return v, nil
	}

	return "", &ErrNoTypeSpecified{}
}

// ParseConfig takes a path to the config directory and attempts to parse the
// config.yml file in that directory. On success a Config is
// returned. On failure a Config and an error is returned.
func ParseConfig(path string) (Config, error) {
	c := make(Config)

	cp, err := configFileName(path)
	if err != nil {
		return c, err
	}

	raw, err := ioutil.ReadFile(cp)
	if err != nil {
		return c, fmt.Errorf("unabled to open config file: %s", cp)
	}

	c, err = unmarshalYamlConfig(raw)
	if err != nil {
		return c, errors.New("unable to unmarshal config")
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

	return cp, errors.New("config file does not exist in config directory")
}

// unmarshalYamlConfig takes the raw yaml []byte and unmarshals it to a
// Config. On success a Config and an error are returned.
// on failure, a Config and error is returned.
func unmarshalYamlConfig(yml []byte) (Config, error) {
	c := make(Config)
	if err := yaml.Unmarshal(yml, c); err != nil {
		return c, errors.New("Unable to unmarshal config file")
	}

	return c, nil
}
