package converge

import (
	"flag"
	"fmt"

	"github.com/mitchellh/cli"
)

func New(ui cli.Ui, converge string) *cmd {
	return &cmd{UI: ui, converge: converge}
}

type cmd struct {
	UI       cli.Ui
	converge string
}

func (c *cmd) Run(args []string) int {
	cmdflags := flag.NewFlagSet("converge", flag.ContinueOnError)

	if err := cmdflags.Parse(args); err != nil {
		return 0
	}
	c.UI.Output(fmt.Sprintf("Attempting to converge config: %s to database type: %s", args[0], args[1]))
	ConvergeConfig(args)
	return 0
}

func (c *cmd) Synopsis() string {
	return "Converges config to database"
}

func (c *cmd) Help() string {
	return "Runs a converge on a config to a database type"
}

func ConvergeConfig(args []string) {

}
