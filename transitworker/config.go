package transitworker

type Config struct {
	// VaultConfig is this Agent's Vault configuration
	VaultConfig *VaultConfig
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	config := &Config{
		VaultConfig: DefaultVaultConfig(),
	}

	return config
}
