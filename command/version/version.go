package version

import (
	"fmt"

	"github.com/mitchellh/cli"
)

// New initializes the version command.
func New(ui cli.Ui, version string) *Cmd {
	return &Cmd{UI: ui, version: version}
}

// Cmd represents a command within mitchellh/cli and stores all the context
// necessary to execute the version command.
type Cmd struct {
	UI      cli.Ui
	version string
}

// Run executes the version command, printing a version string for immigrant.
// the return value should always be 0 as this command should never fail.
func (c *Cmd) Run(_ []string) int {
	c.UI.Output(fmt.Sprintf("immigrant %s", c.version))

	return 0
}

// Synopsis returns a short help string for the version command.
func (c *Cmd) Synopsis() string {
	return "Prints the immigrant cli version"
}

// Help returns an empty string as no additional information is needed to
// execute this command.
func (c *Cmd) Help() string {
	return ""
}
