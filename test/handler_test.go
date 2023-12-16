package api_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-interview/api"
)

func TestProcessSingle(t *testing.T) {

	requestBody, err := json.Marshal(api.SortRequest{
		ToSort: [][]int{{3, 2, 1}, {6, 5, 4}},
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/process-single", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	log.Printf("Status code: %v, Body: %v", responseRecorder.Code, responseRecorder.Body.String())

	handler := http.HandlerFunc(api.ProcessSingle)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp api.SortResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}
}

func TestProcessConcurrent(t *testing.T) {

	requestBody, err := json.Marshal(api.SortRequest{
		ToSort: [][]int{{3, 2, 1}, {6, 5, 4}},
	})
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/process-concurrent", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	request.Header.Set("Content-Type", "application/json")

	responseRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(api.ProcessConcurrent)

	handler.ServeHTTP(responseRecorder, request)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resp api.SortResponse
	err = json.NewDecoder(responseRecorder.Body).Decode(&resp)
	if err != nil {
		t.Fatal(err)
	}
}
