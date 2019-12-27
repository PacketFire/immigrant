package command

import (
	"github.com/PacketFire/immigrant/command/converge"
	"github.com/PacketFire/immigrant/command/version"
	"github.com/PacketFire/immigrant/pkg/config"
	"github.com/mitchellh/cli"
)

const (
	cliVersion string = "0.0.1"
)

// Build takes a config as an argument and attempts to construct a cli registry,
// injecting the config into each method.
func Build(conf config.Config) Registry {
	r := make(Registry)
	r.Register("version", func(ui cli.Ui) (cli.Command, error) { return version.New(conf, ui, cliVersion), nil })
	r.Register("converge", func(ui cli.Ui) (cli.Command, error) { return converge.New(conf, ui), nil })

	return r
}
