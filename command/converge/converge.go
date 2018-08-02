package converge

import (
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

func (c *cmd) Run(_ []string) int {
	c.UI.Output(fmt.Sprintf("immigrant %s", c.converge))

	return 0
}

func (c *cmd) Synopsis() string {
	return "Converges config to db"
}

func (c *cmd) Help() string {
	return ""
}

func ConsumeConfig() {

}

func ApplyConfig() {
	fmt.Println("test")
}
