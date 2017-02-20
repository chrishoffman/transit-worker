package command

import (
	"github.com/chrishoffman/transit-worker/transitworker"
	"github.com/mitchellh/cli"
)

type EncryptCommand struct {
	Ui cli.Ui
}

func (c *EncryptCommand) Help() string {
	return ""
}

func (c *EncryptCommand) Run(_ []string) int {
	transitworker.Encrypt()
	return 0
}

func (c *EncryptCommand) Synopsis() string {
	return ""
}
