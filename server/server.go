package server

import (
	"fmt"
	"net/http"
)

// CreateServerMux returns a new *http.ServeMux with the
// application routes registered
func CreateServerMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", extensoHandle)

	return mux
}

func extensoHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "Path == %v", r.URL.Path)
}
