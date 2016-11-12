package transitworker

type Server struct {
	// vault is the client for communicating with Vault.
	vault vaultClient
}
