package status

import (
	"net/http"

	httputil "github.com/sknv/go-pkg/http"
)

type WebError interface {
	GetCode() ErrorCode
	GetMessage() string
}

// RenderError renders JSON error response.
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	webErr, ok := err.(WebError)
	if !ok {
		webErr = NewError(Internal, err.Error())
	}

	httputil.RenderJSON(w, r, HTTPStatusFromErrorCode(webErr.GetCode()), webErr)
}

// HTTPStatusFromErrorCode maps ErrorCode to HTTP status.
func HTTPStatusFromErrorCode(code ErrorCode) int {
	switch code {
	case InvalidArgument:
		return http.StatusBadRequest
	case NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
