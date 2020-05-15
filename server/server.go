package server

import (
	"net/http"
)

// CreateServerMux returns a new *http.ServeMux with the
// application routes registered
func CreateServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", extensoHandle)

	return mux
}
