package main

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-interview/api"
)

func main() {
	// Get the PORT from the environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/process-single", api.ProcessSingle)
	http.HandleFunc("/process-concurrent", api.ProcessConcurrent)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
