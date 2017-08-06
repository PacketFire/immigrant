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
	ConvergeCommand int    = 0
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
	case "converge":
		return ConvergeCommand
	}

	return InvalidCommand
}

// Shutdown triggers correct shutdown
func Shutdown(code int) {
	os.Exit(code)
}

// Converge takes a driver, sequence and synchronization channel and attempts
// to converge the remote database to local head. The ml channel is provided to
// synchronize errors handlers and report errors. On success, nil is pushed to
// the ml channel. On failure, an error is pushed to the ml channel.
func Converge(drv Driver, seq Sequence, ml chan error) {
	ml <- nil
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

	cp := ConfigPath()
	config, err := ParseConfig(cp)
	if err != nil {
		Shutdown(ExitErr)
	}

	// Ugly but will work, instantiate drive by type
	switch strings.ToLower(config["type"]) {
	case "mysql":
		drv := &MysqlDriver{}
		err = drv.Init(config)
		if err != nil {
			Shutdown(ExitErr)
		}
	default:
		Shutdown(ExitErr)
	}

	// Command router
	switch Command() {
	case ConvergeCommand:
	default:
		Shutdown(ExitErr)
	}
}
