package config

import (
	"testing"
)

const (
	testDir string = "test/"
	errFmt  string = "expected %v got %v"
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

func TestInvokingDriverTypeShould(t *testing.T) {
	t.Run("return a string representation of the type and nil when type key is defined.", func(t *testing.T) {
		expectedType := "mock"
		c := Config{
			"type": expectedType,
		}

		receivedType, _ := c.DriverType()
		if receivedType != expectedType {
			t.Errorf(errFmt, expectedType, receivedType)
		}
	})

	t.Run("return an error when type key is undefined.", func(t *testing.T) {
		c := Config{}

		_, e := c.DriverType()
		if e == nil {
			t.Errorf(errFmt, "error", e)
		}
	})

}
