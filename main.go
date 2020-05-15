package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/ibraimgm/extenso/server"
)

func main() {
	// define and parse command-line flags
	addr := flag.String("address", ":8080", "The listening address of the server")
	flag.Parse()

	// create a mux with the required routes
	mux := server.CreateServerMux()

	// Tries to run the server
	log.Printf("Starting server at address '%s'\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}
