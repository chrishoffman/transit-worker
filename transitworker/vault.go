package transitworker

import (
	"fmt"
	"time"

	"github.com/hashicorp/vault/api"
)

const (
	// DefaultVaultConnectRetryIntv is the retry interval between trying to
	// connect to Vault
	DefaultVaultConnectRetryIntv = 30 * time.Second
)

// VaultConfig contains the configuration information necessary to
// communicate with Vault in order to:
type VaultConfig struct {
	// Token is the Vault token given to Nomad such that it can
	// derive child tokens. Nomad will renew this token at half its lease
	// lifetime.
	Token string

	// Addr is the address of the local Vault agent. This should be a complete
	// URL such as "http://vault.example.com"
	Addr string

	// ConnectionRetryIntv is the interval to wait before re-attempting to
	// connect to Vault.
	ConnectionRetryIntv time.Duration
}

// DefaultVaultConfig() returns the canonical defaults for the Nomad
// `vault` configuration.
func DefaultVaultConfig() *VaultConfig {
	return &VaultConfig{
		Addr:                "https://vault.service.consul:8200",
		ConnectionRetryIntv: DefaultVaultConnectRetryIntv,
	}
}

type vaultClient struct {
	// client is the Vault API client
	client *api.Client

	// token is the raw token used by the client
	token string
}

// NewVaultClient returns a Vault client from the given config. If the client
// couldn't be made an error is returned.
func NewVaultClient(config VaultConfig) (*vaultClient, error) {
	clientConfig := api.DefaultConfig()
	clientConfig.Address = config.Addr

	client, err := api.NewClient(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Vault API client: %v", err)
	}

	client.SetToken(config.Token)

	v := &vaultClient{
		client: client,
		token:  config.Token,
	}

	return v, nil
}
