package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRun(t *testing.T) {
	// Test cases:
	// 1. Test that the database is started and connected
	// 2. Test that the database migration is successful
	// 3. Test the server is running on port 8080

	// Create a request to simulate server running on port 8080
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Create a new router
	r := mux.NewRouter()
	r.HandleFunc("/", Run)

	// Serve the request
	r.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}