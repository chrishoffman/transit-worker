package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type StartCommand struct {
	Ui cli.Ui
}

func (c *StartCommand) Help() string {
	helpText := `
Usage: transit-worker start [options] ...
  Start the transit-worker
Options:
  -vault-addr=https://127.0.0.1:8200 vault server address.
`
	return strings.TrimSpace(helpText)
}

func (c *StartCommand) Run(args []string) int {
	return 0
}

func (c *StartCommand) Synopsis() string {
	return "Starts the transit-worker"
}
