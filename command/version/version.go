package version

import (
	"fmt"

	"github.com/mitchellh/cli"
)

func New(ui cli.Ui, version string) *cmd {
	return &cmd{UI: ui, version: version}
}

type cmd struct {
	UI      cli.Ui
	version string
}

func (c *cmd) Run(_ []string) int {
	c.UI.Output(fmt.Sprintf("immigrant %s", c.version))

	return 0
}

func (c *cmd) Synopsis() string {
	return "Prints the immigrant cli version"
}

func (c *cmd) Help() string {
	return ""
}
