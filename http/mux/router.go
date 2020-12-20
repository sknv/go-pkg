package mux

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	muxware "github.com/sknv/go-pkg/http/mux/middleware"
	"github.com/sknv/go-pkg/log"
)

// Option configures *chi.Mux.
type Option func(*chi.Mux)

// NewRouter returns a new router.
func NewRouter(logger log.Logger, options ...Option) *chi.Mux {
	router := chi.NewRouter()

	// Prepend default middleware
	router.Use(
		middleware.RealIP,
		muxware.RequestID,
		muxware.Logger(logger),
	)

	for _, opt := range options {
		opt(router)
	}

	// Append default middleware
	router.Use(middleware.Recoverer)

	return router
}
