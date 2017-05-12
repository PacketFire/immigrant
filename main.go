package main

import (
	"flag"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const (
	ExitOk          int    = 0
	ExitErr         int    = 1
	ConfigENVVar    string = "IMMIGRANT_CONFIG_DIR"
	InvalidCommand  int    = -1
	MigrateCommand  int    = 0
	RollbackCommand int    = 1
)

var (
	cd *string // flag for config directory
)

func init() {
	cd = flag.String("config-directory",
		"",
		"Specify the config directory")
	flag.Parse()
}

// Attempts to fetch the config path from the following location in the defined
// order.
func ConfigPath() string {
	if *cd != "" {
		return *cd
	}

	if env, err := os.LookupEnv(ConfigENVVar); err == true {
		return env
	}

	if fi, err := os.Stat("immigrant"); err == nil && fi.IsDir() == true {
		return "immigrant"
	}

	if fi, err := os.Stat(".immigrant"); err == nil && fi.IsDir() == true {
		return ".immigrant"
	}

	return ""
}

// Command returns the command to be run
func Command() int {
	if flag.NArg() == 0 {
		return InvalidCommand
	}

	cmd := strings.ToLower(flag.Arg(0))
	switch cmd {
	case "migrate":
		return MigrateCommand
	case "rollback":
		return RollbackCommand
	}

	return InvalidCommand
}

// Shutdown triggers correct shutdown
func Shutdown(code int) {
	os.Exit(code)
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			Shutdown(ExitErr)
		}
	}()

	// ml will function as a migration lock. This will be pushed to each
	// migration/rollback as will blcok both at the signal handler as well as
	// at the cli. This will prevent any actions from being taken until a
	// migration has hit a stable state.
	ml := make(chan error)

	// Signal Handling
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		<-ml
		Shutdown(ExitOk)
	}()
}
