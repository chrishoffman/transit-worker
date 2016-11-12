package transitworker

type Config struct {
	// Address is the address of the transit worker agent
	Address string

	// VaultConfig is this Agent's Vault configuration
	VaultConfig *VaultConfig
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	c := &Config{
		VaultConfig: DefaultVaultConfig(),
	}

	return c
}
