package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/server"
)

func TestGETSubjectsOk(t *testing.T) {
	req, err := http.NewRequest("GET", server.SubjectsPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Subjects)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status,
			http.StatusOK)
	}

	current := strings.Trim(rr.Body.String(), "\n")
	expected := `[]`

	if current != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", current,
			expected)
	}
}

func TestPOSTSubjectsOk(t *testing.T) {
	req, err := http.NewRequest("POST", server.SubjectsPath, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.Subjects)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status,
			http.StatusCreated)
	}

	current := strings.Trim(rr.Body.String(), "\n")
	expected := `{"Current":{"Latitude":0,"Longitude":0,"CreatedAt":"0001-01-01T00:00:00Z"},"History":null}`

	if current != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", current,
			expected)
	}
}
