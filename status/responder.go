package status

import (
	"net/http"

	"github.com/pkg/errors"

	httputil "github.com/sknv/go-pkg/http"
)

// RenderError renders JSON error response.
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	webErr := Error{Message: err.Error()}

	var cause *Error
	if errors.As(err, &cause) {
		webErr.Code = cause.Code
		webErr.Message = cause.Message
	}

	httputil.RenderJSON(w, r, HTTPStatusFromErrorCode(webErr.Code), webErr)
}

// HTTPStatusFromErrorCode maps ErrorCode to HTTP status.
func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	default:
		return http.StatusInternalServerError
	}
}
