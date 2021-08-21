package problem

import (
	"errors"
	"net/http"

	"github.com/sknv/go-pkg/http/render"
)

const (
	_defaultHTTPStatus = http.StatusInternalServerError
)

type jsonError struct {
	Type       string      `json:"type"`
	Title      string      `json:"title"`
	Status     int         `json:"status"`
	Detail     string      `json:"detail,omitempty"`
	Instance   string      `json:"instance,omitempty"`
	Extensions interface{} `json:"extensions,omitempty"`
}

// Render renders JSON error response.
func Render(w http.ResponseWriter, err error) {
	var problem *Problem
	if !errors.As(err, &problem) {
		problem = New(_defaultHTTPStatus, err.Error())
	}

	render.JSON(w, problem.GetStatus(), jsonError{
		Type:       problem.GetType(),
		Title:      problem.GetTitle(),
		Status:     problem.GetStatus(),
		Detail:     problem.GetDetail(),
		Instance:   problem.GetInstance(),
		Extensions: problem.GetExtensions(),
	})
}
