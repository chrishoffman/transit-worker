package transitworker

import (
	"github.com/braintree/manners"
	"github.com/chrishoffman/transit-worker/handlers"
	"github.com/julienschmidt/httprouter"
)

type Server struct {
	config     *Config
	httpServer *manners.GracefulServer

	// vault is the client for communicating with Vault.
	vault vaultClient
}

func NewServer(config *Config) (*Server, error) {
	router := httprouter.New()
	router.GET("/v1/:mount/encrypt/:name", handlers.Hello)

	httpServer := manners.NewServer()
	httpServer.Addr = config.Address
	httpServer.Handler = handlers.LoggingHandler(router)

	// Create the server
	s := &Server{
		config:     config,
		httpServer: httpServer,
	}

	go s.httpServer.ListenAndServe()

	return s, nil
}

func (s *Server) Shutdown() {
	s.httpServer.BlockingClose()
}
