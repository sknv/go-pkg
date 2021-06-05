package http

import (
	"net/http"

	"github.com/go-chi/render"
)

// RenderJSON renders JSON data response with the provided status.
func RenderJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	render.Status(r, status)
	render.JSON(w, r, data)
}
