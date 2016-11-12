package transitworker

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

type vaultClient struct {
	// client is the Vault API client
	client *api.Client

	// token is the raw token used by the client
	token string
}

// NewVaultClient returns a Vault client from the given config. If the client
// couldn't be made an error is returned.
func NewVaultClient(config *api.Config, token string) (*vaultClient, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Vault API client: %v", err)
	}

	client.SetToken(token)

	v := &vaultClient{
		client: client,
		token:  token,
	}

	return v, nil
}
