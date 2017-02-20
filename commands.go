package main

import (
	"os"

	"github.com/chrishoffman/transit-worker/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available sr6 commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}
	Commands = map[string]cli.CommandFactory{
		"encrypt": func() (cli.Command, error) {
			return &command.EncryptCommand{
				Ui: ui,
			}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: Version,
				Ui:      ui,
			}, nil
		},
	}
}
