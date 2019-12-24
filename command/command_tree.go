package command

import (
	"github.com/PacketFire/immigrant/command/context"
	"github.com/PacketFire/immigrant/command/converge"
	"github.com/PacketFire/immigrant/command/version"
	"github.com/mitchellh/cli"
)

const (
	cliVersion string = "0.0.1"
)

// Build takes a ctx as an argument and attempts to construct a cli registry,
// injecting the context into each method.
func Build(ctx context.Context) Registry {
	r := make(Registry)
	r.Register("version", func(ui cli.Ui) (cli.Command, error) { return version.New(ctx, ui, cliVersion), nil })
	r.Register("converge", func(ui cli.Ui) (cli.Command, error) { return converge.New(ctx, ui), nil })

	return r
}
