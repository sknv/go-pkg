package render

import (
	"bytes"
	"net/http"
)

// JSON renders JSON data response with the provided status,
// automatically escaping HTML and setting the Content-Type as application/json.
func JSON(w http.ResponseWriter, status int, data any) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write(buf.Bytes())
}
