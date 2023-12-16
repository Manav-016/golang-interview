package main

import (
	"log"
	"net/http"

	"github.com/golang-interview/api"
)

func main() {
	http.HandleFunc("/process-single", api.ProcessSingle)
	http.HandleFunc("/process-concurrent", api.ProcessConcurrent)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
