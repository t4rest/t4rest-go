package httpserver

import "net/http"

// StatusRecorder .
type StatusRecorder struct {
	http.ResponseWriter
	Status int
}

// WriteHeader overwrites default WriteHeader to record response status
func (r *StatusRecorder) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}
