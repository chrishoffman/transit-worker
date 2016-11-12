package command

import (
	"log"
	"net/http"
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

	router := transitworker.SetupRoutes()

	go http.ListenAndServe(":8080", router)

	return c.handleSignals()
}

func (c *AgentCommand) Synopsis() string {
	return "Starts the transit-worker agent"
}

// handleSignals monitors the shutdownCh channel and acts on it
func (c *AgentCommand) handleSignals() int {
	select {
	case <-c.ShutdownCh:
		if err := c.server.Shutdown(); err != nil {
			log.Println("[INFO] transit-worker: Couldn't properly shutdown the server")
			return 1
		}
		return 0
	}
}
