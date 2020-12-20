package http

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/sknv/go-pkg/status"
)

type webError struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type webResponse struct {
	Data interface{} `json:"data,omitempty"`
	Err  webError    `json:"err,omitempty"`
}

// RenderData renders JSON data response.
func RenderData(w http.ResponseWriter, r *http.Request, data interface{}) {
	render.JSON(w, r, webResponse{Data: data})
}

// RenderError renders JSON error response.
func RenderError(w http.ResponseWriter, r *http.Request, err error) {
	// Prepare the error to render
	webErr := webError{Msg: err.Error()}

	var cause status.Error
	if errors.As(err, &cause) {
		webErr.Code = cause.Code
		webErr.Msg = cause.Message
	}

	code := http.StatusBadRequest                                     // status for unknown errors
	if statusText := http.StatusText(webErr.Code); statusText != "" { // known status
		code = webErr.Code
	}
	render.Status(r, code)
	render.JSON(w, r, webResponse{Err: webErr})
}
