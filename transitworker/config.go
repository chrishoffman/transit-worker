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
		Address:     "http://127.0.0.1:8282",
		VaultConfig: DefaultVaultConfig(),
	}

	return c
}
