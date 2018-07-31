package config

import (
	"testing"
)

const (
	testDir string = "test/"
)

func TestParseConfig(t *testing.T) {
	c, err := ParseConfig(testDir)
	if err != nil {
		t.Errorf("Unable to parse Config.")
	}

	if drv, prs := c["type"]; prs != true && drv != "MySQL" {
		t.Errorf("Config unmarshal failing.")
	}
}
