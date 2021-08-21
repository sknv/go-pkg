package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/sknv/go-pkg/log"
)

// WithLogger injects the provided logger into request context.
func WithLogger(logger *logrus.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ctxLog := log.ToContext(r.Context(), logger)
			next.ServeHTTP(w, r.WithContext(ctxLog))
		}
		return http.HandlerFunc(fn)
	}
}

// Log is a slightly modified version of the provided logger middleware.
func Log(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor) // save a response status
		next.ServeHTTP(ww, r)

		// Log
		log.Extract(r.Context()).WithFields(logrus.Fields{
			"op":       "http",
			"uri":      fmt.Sprintf("%s %s%s", r.Method, r.Host, r.RequestURI),
			"status":   ww.Status(),
			"ip":       r.RemoteAddr,
			"bytes_in": r.Header.Get("Content-Length"),
			"bytes_ot": ww.BytesWritten(),
			"latency":  time.Since(start),
		}).Info()
	}
	return http.HandlerFunc(fn)
}
