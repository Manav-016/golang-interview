package api

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-interview/pkg/sorting"
)

type SortRequest struct {
	ToSort [][]int `json:"to_sort"`
}

type SortResponse struct {
	SortedArrays [][]int `json:"sorted_arrays"`
	TimeNS       int64   `json:"time_ns"`
}

func ProcessSingle(w http.ResponseWriter, r *http.Request) {
	var req SortRequest

	parseErr := json.NewDecoder(r.Body).Decode(&req)
	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusBadRequest)
		return
	}

	start := time.Now()
	sorting.SequentialSort(req.ToSort)
	elapsed := time.Since(start)

	log.Printf("Sorted arrays with signle process: %v", req.ToSort)
	log.Printf("Took %s", elapsed)

	resp := SortResponse{
		SortedArrays: req.ToSort,
		TimeNS:       elapsed.Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Error sending response: %v", parseErr)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

}

func ProcessConcurrent(w http.ResponseWriter, r *http.Request) {
	var req SortRequest

	parseErr := json.NewDecoder(r.Body).Decode(&req)

	if parseErr != nil {
		http.Error(w, parseErr.Error(), http.StatusBadRequest)
		return
	}

	start := time.Now()

	sorting.ConcurrentSort(req.ToSort)

	elapsed := time.Since(start)

	log.Printf("Sorted arrays with concurrent process: %v", req.ToSort)
	log.Printf("Took %s", elapsed)

	resp := SortResponse{
		SortedArrays: req.ToSort,
		TimeNS:       elapsed.Nanoseconds(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		log.Printf("Error sending response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}

}
