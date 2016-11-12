package transitworker

type Server struct {
	config *Config

	// vault is the client for communicating with Vault.
	vault vaultClient
}

func NewServier(config *Config) (*Server, error) {
	// Create the server
	s := &Server{
		config: config,
	}

	return s, nil
}
