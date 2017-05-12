package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

const (
	ExitOk  int = 0
	ExitErr int = 1
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
