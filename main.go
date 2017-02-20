package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

// Version and BuildTime are filled in during build by the Makefile
var (
	Version   = "N/A"
	BuildTime = "N/A"
)

func main() {
	args := os.Args[1:]

	// Get the command line args. We shortcut "--version" and "-v" to
	// just show the version.
	for _, arg := range args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, 1)
			newArgs[0] = "version"
			args = newArgs
			break
		}
	}

	cli := &cli.CLI{
		Args:     args,
		Commands: Commands,
		HelpFunc: cli.BasicHelpFunc("transit-worker"),
	}
	exitCode, err := cli.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(exitCode)
}
