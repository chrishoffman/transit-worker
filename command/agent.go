package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/chrishoffman/transit-worker/transitworker"
	"github.com/mitchellh/cli"
)

type AgentCommand struct {
	Ui         cli.Ui
	ShutdownCh <-chan struct{}

	args   []string
	server *transitworker.Server
}

func (c *AgentCommand) Help() string {
	helpText := `
Usage: transit-worker agent [options] ...
  Start the transit-worker agent
Options:
  -vault-addr=https://127.0.0.1:8200 vault server address.
`
	return strings.TrimSpace(helpText)
}

func (c *AgentCommand) Run(args []string) int {
	log.Println("Starting transit-worker...")
	c.args = args

	config := transitworker.DefaultConfig()
	server, err := transitworker.NewServer(config)
	if err != nil {
		fmt.Errorf("Unable to set up server: %v", err)
	}

	c.server = server
	return c.handleSignals()
}

func (c *AgentCommand) Synopsis() string {
	return "Starts the transit-worker agent"
}

// handleSignals monitors the shutdownCh channel and acts on it
func (c *AgentCommand) handleSignals() int {
	select {
	case <-c.ShutdownCh:
		c.server.Shutdown()
		return 0
	}
}
