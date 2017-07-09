package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterOk(t *testing.T) {
	pin := "abcd1234"
	reader := strings.NewReader(fmt.Sprintf("%s=%s", PIN_NAME, pin))
	req, err := http.NewRequest("POST", REGISTER_PATH, reader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(register)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status,
			http.StatusCreated)
	}

	current := strings.Trim(rr.Body.String(), "\n")
	expected := `{"success":true,"message":"Successfully created ID.","id":"`

	if !strings.Contains(current, expected) {
		t.Errorf("Handler returned unexpected body: got %v want to contain %v",
			current, expected)
	}
}
