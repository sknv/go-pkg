package middleware

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/atomic"

	"github.com/sknv/go-pkg/log"
	"github.com/sknv/go-pkg/rand"
)

const (
	HeaderRequestID = "X-Request-Id"
)

// Global for the current process.
var (
	_prefix string
	_reqID  atomic.Uint64
)

func init() {
	_prefix = makePrefix()
}

// RequestID is a middleware that injects a request id into the context of each request.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(HeaderRequestID)
		if requestID == "" {
			curReqID := _reqID.Add(1)
			requestID = fmt.Sprintf("%s-%d", _prefix, curReqID)
		}
		ctx := log.PutRequestID(r.Context(), requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}

func makePrefix() string {
	hostname, err := os.Hostname()
	if hostname == "" || err != nil {
		hostname = "localhost"
	}
	return fmt.Sprintf("%s/%s", hostname, rand.String(8))
}
