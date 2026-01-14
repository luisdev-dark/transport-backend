package main

import (
	"log"
	"net/http"
	"os"

	"transport-backend/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, http.HandlerFunc(api.Handler)))
}
