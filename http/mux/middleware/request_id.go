package middleware

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/atomic"

	"github.com/sknv/go-pkg/log"
	"github.com/sknv/go-pkg/rand"
)

const HeaderRequestID = "X-Request-ID"

// Global for the current process.
var (
	_prefix string
	_reqID  atomic.Uint64
)

func init() {
	_prefix = makePrefix()
}

const _fieldRequestID = "request_id"

// RequestID is a middleware that injects a request id into the context of each request.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(HeaderRequestID)
		if requestID == "" {
			requestID = newRequestID()
			w.Header().Add(HeaderRequestID, requestID) // provide request id to client to trace issues
		}
		log.AddFields(r.Context(), log.Fields{_fieldRequestID: requestID})
		next.ServeHTTP(w, r)
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

func newRequestID() string {
	curReqID := _reqID.Add(1)
	return fmt.Sprintf("%s-%d", _prefix, curReqID)
}
