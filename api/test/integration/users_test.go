package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/api"
	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
)

func checkExpectedStatus(t *testing.T, response *httptest.ResponseRecorder, expectedStatus int) {
	if response.Code != expectedStatus {
		t.Errorf("status mismatch! expected %d, got %d: %v", expectedStatus, response.Code, response.Body)
		t.FailNow()
	}
}

func checkBody(t *testing.T, response *httptest.ResponseRecorder, expectedBody string) {
	if response.Body == nil {
		t.Error("body is nil")
		t.FailNow()
	}

	if body := strings.TrimSpace(response.Body.String()); body != expectedBody {
		t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
		t.FailNow()
	}
}

func TestUserManagement(t *testing.T) {

	t.Run("Getting user information", func(t *testing.T) {

		t.Run("Existing user", func(t *testing.T) {
			defer dao.CleanMocks()
			u := domain.NewUser("ariel17")
			if err := dao.Client.InsertUser(u); err != nil {
				t.Error(err)
				t.FailNow()
			}

			router := api.ConfigureRouter()
			url := fmt.Sprintf("/users/%s", string(u.ID))
			req, _ := http.NewRequest("GET", url, nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			checkExpectedStatus(t, resp, http.StatusOK)
			checkBody(t, resp, fmt.Sprintf(`{"success":true,"message":"user found","data":{"_id":"%s","nick":"%s"}}`, u.ID, u.Nick))
		})

		t.Run("Not found", func(t *testing.T) {
			router := api.ConfigureRouter()
			req, _ := http.NewRequest("GET", "/users/9999", nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			checkExpectedStatus(t, resp, http.StatusNotFound)
			checkBody(t, resp, `{"success":false,"message":"user 9999 not found"}`)
		})
	})

	t.Run("Creating a new user", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			defer dao.CleanMocks()
			u := domain.NewUser("ariel17")
			if bytes, err := json.Marshal(u); err != nil {
				t.Error(err)
				t.FailNow()
			}

			router := api.ConfigureRouter()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(bytes))
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			checkExpectedStatus(t, resp, http.StatusCreated)
			checkBody(t, resp, fmt.Sprintf(`{"success":true,"message":"user created","data":{"_id":"%s","nick":"%s"}}`, u.ID, u.Nick))

			if uu, err := dao.Client.GetUser(u.ID); err != nil {
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
