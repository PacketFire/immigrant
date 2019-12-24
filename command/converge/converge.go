package converge

import (
	"github.com/PacketFire/immigrant/command/context"
	"github.com/mitchellh/cli"
)

// New initializes the converge command.
func New(ctx context.Context, ui cli.Ui) *Cmd {
	return &Cmd{UI: ui, ctx: ctx}
}

// Cmd represents a command within mitchellh/cli and stores all the context
// necessary to execute the version command.
type Cmd struct {
	UI  cli.Ui
	ctx context.Context
}

// Run executes the converge command. An integer representing the success of
//the execution is returned. This is currently a stub.
func (c *Cmd) Run(_ []string) int {

	c.UI.Output("STUB")

	return 0
}

// Synopsis returns a short help string for the converge command.
func (c *Cmd) Synopsis() string {
	return "Converges local revision tree onto the configured store"
}

// Help returns an empty string as no additional information is needed to
// execute this command.
func (c *Cmd) Help() string {
	return ""
}
