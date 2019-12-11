package command

import (
	"github.com/PacketFire/immigrant/command/converge"
	"github.com/PacketFire/immigrant/command/version"
	"github.com/mitchellh/cli"
)

const (
	cliVersion string = "0.0.1"
)

func init() {
	Register("version", func(ui cli.Ui) (cli.Command, error) { return version.New(ui, cliVersion), nil })
	Register("converge", func(ui cli.Ui) (cli.Command, error) { return converge.New(ui), nil })
}
