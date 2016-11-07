package transitworker

import (
	"github.com/hashicorp/vault/api"
)

type vaultClient struct {
	// client is the Vault API client
	client *api.Client
}

// NewVaultClient returns a Vault client from the given config. If the client
// couldn't be made an error is returned.
func NewVaultClient() (*vaultClient, error) {
	return &vaultClient{}, nil
}
