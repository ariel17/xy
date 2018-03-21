package integration

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/api"
	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
)

func TestUserManagement(t *testing.T) {

	t.Run("Getting user information", func(t *testing.T) {
		t.Run("Existing user", func(t *testing.T) {
			defer dao.CleanMocks()
			u := domain.User{
				ID:   "abc123",
				Nick: "ariel17",
			}
			if err := dao.Client.InsertUser(&u); err != nil {
				t.Error(err)
				t.FailNow()
			}

			router := api.ConfigureRouter()
			req, _ := http.NewRequest("GET", "/users/abc123", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			expectedStatus := http.StatusOK
			if resp.Code != expectedStatus {
				t.Errorf("status mismatch! expected %d, got %d: %v", expectedStatus, resp.Code, resp.Body)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			expectedBody := `{"success":true,"message":"user found","data":{"_id":"616263313233","nick":"ariel17"}}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}
		})

		t.Run("Not found", func(t *testing.T) {
			router := api.ConfigureRouter()
			req, _ := http.NewRequest("GET", "/users/9999", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			expectedStatus := http.StatusNotFound
			if resp.Code != expectedStatus {
				t.Errorf("status mismatch! expected %d, got %d: %v", expectedStatus, resp.Code, resp.Body)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			expectedBody := `{"success":false,"message":"user 9999 not found"}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}
		})
	})

	t.Run("Creating a new user", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			defer dao.CleanMocks()
			u := domain.User{
				Nick: "ariel17",
			}

			router := api.ConfigureRouter()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(`{"nick":"ariel17"}`)))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			expectedStatus := http.StatusCreated
			if resp.Code != expectedStatus {
				t.Errorf("status mismatch! expected %d, got %d: %v", expectedStatus, resp.Code, resp.Body)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			expectedBody := `{"success":true,"message":"user created","data":{"_id":"","nick":"ariel17"}}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}

			if uu, err := dao.Client.GetUser(""); err != nil {
				t.Error(err)
				t.FailNow()
			} else if uu == nil {
				t.Errorf("created user not found: expected: %v", u)
				t.FailNow()
			} else if *uu != u {
				t.Errorf("created user mismatch: expected %v, got %v", u, uu)
				t.FailNow()
			}
		})

		t.Run("Missing required fields", func(t *testing.T) {
			t.FailNow()
		})
	})

	t.Run("Deleting an user", func(t *testing.T) {
		t.Run("Existing user", func(t *testing.T) {
			defer dao.CleanMocks()
			u := domain.User{
				ID:   "abc123",
				Nick: "ariel17",
			}
			if err := dao.Client.InsertUser(&u); err != nil {
				t.Error(err)
				t.FailNow()
			}

			router := api.ConfigureRouter()
			req, _ := http.NewRequest("DELETE", "/users/abc123", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			expectedStatus := http.StatusOK
			if resp.Code != expectedStatus {
				t.Errorf("status mismatch! expected %d, got %d: %v", expectedStatus, resp.Code, resp.Body)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			expectedBody := `{"success":true,"message":"user abc123 deleted","data":{"_id":"616263313233","nick":"ariel17"}}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}

			if uu, err := dao.Client.GetUser("abc123"); err != nil {
				t.Error(err)
				t.FailNow()
			} else if uu != nil {
				t.Errorf("deleted user found! %v", uu)
				t.FailNow()
			}
		})

		t.Run("User not found", func(t *testing.T) {
			t.FailNow()
		})
	})
}
