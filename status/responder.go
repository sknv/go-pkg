package status

import (
	"net/http"

	httputil "github.com/sknv/go-pkg/http"
)

type DescriptiveError interface {
	GetCode() ErrorCode
	GetMessage() string
}

type jsonError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// RenderError renders JSON error response.
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	descErr, ok := err.(DescriptiveError)
	if !ok {
		descErr = NewError(Internal, err.Error())
	}

	status := HTTPStatusFromErrorCode(descErr.GetCode())
	jsErr := jsonError{
		Code:    string(descErr.GetCode()),
		Message: descErr.GetMessage(),
	}
	httputil.RenderJSON(w, r, status, jsErr)
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
