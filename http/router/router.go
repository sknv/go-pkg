package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	pkgmware "github.com/sknv/go-pkg/http/router/middleware"
	"github.com/sknv/go-pkg/log"
)

// Option configures *chi.Mux.
type Option func(*chi.Mux)

// New returns a new router.
func New(options ...Option) *chi.Mux {
	router := chi.NewRouter()

	// Prepend default middleware, order matters
	router.Use(
		pkgmware.WithLogger(log.L()),
		middleware.RealIP,
		pkgmware.WithRequestID,
		pkgmware.Log,
	)

	for _, opt := range options {
		opt(router)
	}

	// Append default middleware
	router.Use(
		pkgmware.Recover,
	)
	return router
}
