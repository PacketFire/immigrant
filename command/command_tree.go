package command

import (
	"github.com/ncatelli/immigrant/command/converge"
	"github.com/ncatelli/immigrant/command/version"

	"github.com/mitchellh/cli"
)

const (
	cliVersion   string = "0.0.1"
	convergeTest string = "testing"
)

func init() {
	Register("version", func(ui cli.Ui) (cli.Command, error) { return version.New(ui, cliVersion), nil })
	Register("converge", func(ui cli.Ui) (cli.Command, error) { return converge.New(ui, convergeTest), nil })
}
