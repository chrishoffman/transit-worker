package command

import (
	"strings"

	"github.com/mitchellh/cli"
)

type AgentCommand struct {
	Ui cli.Ui
}

func (c *AgentCommand) Help() string {
	helpText := `
Usage: transit-worker start [options] ...
  Start the transit-worker
Options:
  -vault-addr=https://127.0.0.1:8200 vault server address.
`
	return strings.TrimSpace(helpText)
}

func (c *AgentCommand) Run(args []string) int {
	return 0
}

func (c *AgentCommand) Synopsis() string {
	return "Starts the transit-worker"
}
