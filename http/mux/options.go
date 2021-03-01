package mux

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// WithThrottle option.
func WithThrottle(limit int) Option {
	return func(r *chi.Mux) {
		r.Use(middleware.Throttle(limit))
	}
}

// WithTimeout option.
func WithTimeout(timeout time.Duration) Option {
	return func(r *chi.Mux) {
		r.Use(middleware.Timeout(timeout))
	}
}
