package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/sknv/go-pkg/log"
)

const _msgHTTPRequest = "http request"

// Logger is a slightly modified version of a provided logger middleware.
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor) // save a response status
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		defer func(start time.Time) {
			log.Extract(r.Context()).WithFields(log.Fields{
				"uri":     fmt.Sprintf("%s %s://%s%s", r.Method, scheme, r.Host, r.RequestURI),
				"status":  ww.Status(),
				"ip":      r.RemoteAddr,
				"latency": time.Since(start),
			}).Info(_msgHTTPRequest)
		}(time.Now())

		next.ServeHTTP(ww, r)
	}
	return http.HandlerFunc(fn)
}
