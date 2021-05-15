package http

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/pkg/errors"
)

// Start starts a web server.
func Start(server *http.Server, listener net.Listener) {
	log.Printf("starting an http server on %s", listener.Addr())
	if err := server.Serve(listener); err != nil {
		// Cannot error, because this probably is an intentional close
		log.Printf("http server stopped, reason: %s", err)
	}
}

// Stop stops a web server.
func Stop(ctx context.Context, server *http.Server) error {
	if err := server.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "shutdown server")
	}
	return nil
}
