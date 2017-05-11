package main

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	c, err := ParseConfig(revisionTestDir)
	if err != nil {
		t.Errorf("Unable to parse Config.")
	}

	if drv, prs := c["Driver"]; prs != true && drv != "MySQL" {
		t.Errorf("Config unmarshal failing.")
	}
}
