package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/sknv/go-pkg/log"
)

const (
	HeaderRequestID = "X-Request-ID"
)

const (
	_fieldRequestID = "request_id"
)

// WithRequestID is a middleware that injects a request id into the context of each request.
func WithRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		w.Header().Add(HeaderRequestID, requestID) // provide request id to client to trace issues
		log.AddFields(r.Context(), logrus.Fields{_fieldRequestID: requestID})

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
