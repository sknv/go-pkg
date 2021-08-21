package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/sknv/go-pkg/http/problem"
	"github.com/sknv/go-pkg/log"
)

// Recover is a slightly modified version of the provided recover middleware.
func Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			//nolint:errorlint // rvr is not always an error
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				log.Extract(r.Context()).WithField("recover", rvr).Errorf("panic\n%s", debug.Stack())
				problem.Render(w, problem.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)))
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
