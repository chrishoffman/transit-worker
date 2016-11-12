package transitworker

import "os"

type Config struct {
	// Address is the address of the transit worker agent
	Address string

	// VaultConfig is this Agent's Vault configuration
	VaultConfig *VaultConfig
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	config := &Config{
		Address:     "127.0.0.1:8282",
		VaultConfig: DefaultVaultConfig(),
	}

	if addr := os.Getenv("TRANSIT_WORKER_ADDR"); addr != "" {
		config.Address = addr
	}

	return config
}
