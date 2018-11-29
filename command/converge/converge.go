package converge

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/cli"
	yaml "gopkg.in/yaml.v2"
)

func New(ui cli.Ui) *cmd {
	return &cmd{UI: ui}
}

type cmd struct {
	UI       cli.Ui
	converge string
	dcfg     DbConfig
	mcfg     MigConfig
}

type DbConfig struct {
	Dtype    string `yaml:"type"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MigConfig struct {
	Rev      string `yaml:"revision,omitempty"`
	Migrate  string `yaml:"migrate,omitempty"`
	Rollback string `yaml:"rollback,omitempty"`
}

func (c *cmd) Run(args []string) int {
	cmdflags := flag.NewFlagSet("converge", flag.ContinueOnError)

	if err := cmdflags.Parse(args); err != nil {
		return 0
	}

	c.HandleConvergeArguments(args)
	c.RunMigrations()

	return 0
}

func (c *cmd) Synopsis() string {
	return "Converges config to database"
}

func (c *cmd) Help() string {
	return "Runs a converge on a config to a database type"
}

func (c *cmd) HandleConvergeArguments(args []string) {
	if len(args) == 1 {
		c.MigrationConfigReader(args[0])
	} else {
		c.DatabaseConfigReader(args[0])
		c.MigrationConfigReader(args[1])
	}
}

func (c *cmd) DatabaseConfigReader(config string) {
	configFile, err := ioutil.ReadFile("examples/" + config)
	if err != nil {
		fmt.Println(err)
	}

	if err = yaml.Unmarshal(configFile, &c.dcfg); err != nil {
		fmt.Println(err)
	}
}

func (c *cmd) MigrationConfigReader(config string) {
	migrationFile, err := ioutil.ReadFile("examples/migrations/" + config)
	if err != nil {
		fmt.Println(err)
	}

	if err := yaml.Unmarshal(migrationFile, &c.mcfg); err != nil {
		fmt.Println(err)
	}
}

func (c *cmd) RunMigrations() {
	msg := "Attempting migrations on database"
	switch c.dcfg.Dtype {
	case "mysql":
		fmt.Println(msg, c.dcfg.Dtype)

	case "postgres":
		fmt.Println(msg, c.dcfg.Dtype)

	case "sqlite":
		fmt.Println(msg, c.dcfg.Dtype)
	}
}
