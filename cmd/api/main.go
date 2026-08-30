package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OXY7O/example-app-go/internal/httpapi"
)

func main() {
	server := newServer(os.Getenv("HTTP_ADDRESS"))
	log.Printf("Go Service example listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func newServer(address string) *http.Server {
	if address == "" {
		address = ":8080"
	}

	return &http.Server{
		Addr:              address,
		Handler:           httpapi.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
}
