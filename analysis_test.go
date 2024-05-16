package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnalysisHandler_Status(t *testing.T) {
	// Create a request with method GET
	req, err := http.NewRequest("GET", "/analysis?duration=30s&dimension=likes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Call the AnalysisHandler function
	AnalysisHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestAnalysisHandler_ContentType(t *testing.T) {
	// Create a request with method GET
	req, err := http.NewRequest("GET", "/analysis?duration=30s&dimension=likes", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder
	rr := httptest.NewRecorder()

	// Call the AnalysisHandler function
	AnalysisHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the Content-Type header
	contentType := rr.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("handler returned wrong Content-Type header: got %v want %v", contentType, "application/json")
	}
}
