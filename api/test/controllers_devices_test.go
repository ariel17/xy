package test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/api"
	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
	"gopkg.in/mgo.v2/bson"
)

func TestDeviceRegisterController(t *testing.T) {

	t.Run("POST", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {

			userID := bson.NewObjectId().Hex()
			user := domain.User{
				ID: bson.ObjectIdHex(userID),
			}

			deviceID := bson.NewObjectId().Hex()
			device := domain.Device{
				ID:     bson.ObjectIdHex(deviceID),
				UserID: user.ID,
			}

			buffer := bytes.NewBuffer([]byte(fmt.Sprintf(`{"user_id":"%s","pin":"xxx","device":{"model":"xyz"}}`, userID)))
			req, _ := http.NewRequest("POST", "/register/", buffer)
			resp := httptest.NewRecorder()

			router := api.ConfigureRouter()
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

			expectedBody := fmt.Sprintf(`{"success":true,"message":"device registered","data":{"_id":"%s"}}`, deviceID)
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}

			d, err := dao.Client.GetDevice(deviceID)
			if err != nil {
				t.Error(err)
				t.FailNow()
			} else if d == nil {
				t.Error("device not found!")
				t.FailNow()
			}

			devices, err := dao.Client.GetUserDevices(userID)
			if err != nil {
				t.Error(err)
				t.FailNow()
			} else if len(devices) == 0 {
				t.Error("devices for user not found!")
				t.FailNow()
			}

			var found bool
			for _, x := range devices {
				if x.UserID == device.UserID {
					found = true
				}
			}

			if !found {
				t.Error("device not found in user's devices!")
				t.FailNow()
			}
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
