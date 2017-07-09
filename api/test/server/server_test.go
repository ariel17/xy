package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetSubjectsOk(t *testing.T) {
	req, err := http.NewRequest("GET", SUBJECTS_PATH, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(subjects)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status,
			http.StatusOK)
	}

	current := strings.Trim(rr.Body.String(), "\n")
	expected := `[]`

	if current != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", current,
			expected)
	}
}
