package converge

import (
	"github.com/PacketFire/immigrant/pkg/config"
	"github.com/mitchellh/cli"
)

// New initializes the converge command.
func New(conf config.Config, ui cli.Ui) *Cmd {
	return &Cmd{UI: ui, config: conf}
}

// Cmd represents a command within mitchellh/cli and stores all the config
// necessary to execute the version command.
type Cmd struct {
	UI     cli.Ui
	config config.Config
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
