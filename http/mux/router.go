package mux

import (
	"github.com/go-chi/chi/v5"
	chi_middleware "github.com/go-chi/chi/v5/middleware"

	"github.com/sknv/go-pkg/http/mux/middleware"
)

// Option configures *chi.Mux.
type Option func(*chi.Mux)

// NewRouter returns a new router.
func NewRouter(options ...Option) *chi.Mux {
	router := chi.NewRouter()

	// Prepend default middleware, order matters
	router.Use(
		chi_middleware.RealIP,
		middleware.Logger,
		middleware.RequestID,
	)

	for _, opt := range options {
		opt(router)
	}

	// Append default middleware
	router.Use(chi_middleware.Recoverer)

	return router
}
