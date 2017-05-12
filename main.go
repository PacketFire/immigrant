package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	ExitOk       int    = 0
	ExitErr      int    = 1
	ConfigENVVar string = "IMMIGRANT_CONFIG_DIR"
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

	// Signal Handling
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		Shutdown(ExitOk)
	}()
}
