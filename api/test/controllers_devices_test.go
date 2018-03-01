package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/api"
	"github.com/ariel17/xy/api/dao"
)

func TestDeviceRegisterController(t *testing.T) {

	t.Run("POST", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {

			// TODO mock user
			// TODO mock devices

			router := api.ConfigureRouter()
			req, _ := http.NewRequest("POST", "/register/", bytes.NewBuffer([]byte(`{"auth":{"user_id":"abc123","pin":"xxx"},"device":{"model":"xyz"}}`)))
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

			expectedBody := `{"success":true,"message":"device registered","data":{"_id":"616263313233"}}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}

			if d, err := dao.Devices.GetDevice("616263313233"); err != nil {
				t.Error(err)
				t.FailNow()
			} else if d == nil {
				t.Error("device not found!")
				t.FailNow()
			}
			// TODO compare if device belongs to user
		})

		t.Run("UserNotFound", func(t *testing.T) {
		})
	})
}

func TestDevicePositionController(t *testing.T) {

	t.Run("Current", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			t.Run("OK", func(t *testing.T) {
			})
		})

		t.Run("POST", func(t *testing.T) {
			t.Run("OK", func(t *testing.T) {
			})
		})
	})

	t.Run("History", func(t *testing.T) {
		t.Run("GET", func(t *testing.T) {
			t.Run("OK", func(t *testing.T) {
			})
		})
	})
}
