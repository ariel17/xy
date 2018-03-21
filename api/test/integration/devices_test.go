package integration

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

func TestDeviceRegistration(t *testing.T) {

	t.Run("Registering a new device", func(t *testing.T) {
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

		t.Run("Failed when user is not found", func(t *testing.T) {
			t.FailNow()
		})
	})

	t.Run("Getting the user's devices", func(t *testing.T) {
		t.Run("Existing user", func(t *testing.T) {
			`{
				"status":"success",
				"message":"User's devices fetch successful.",
				"data": [
					{"_id":"1","model":"Model X","created_at":"2018-01-01T00:00:00Z"},
					{"_id":"2","model":"Model Y","created_at":"2018-01-02T00:00:00Z"}
				]
			}`

			t.FailNow()
		})

		t.Run("User not found", func(t *testing.T) {
			`{
				"status":"failed",
				"message":"User not exists.",
			}`

			t.FailNow()
		})
	})
}

func TestDevicePosition(t *testing.T) {

	t.Run("Getting current device position", func(t *testing.T) {
		`{
			"status":"success",
			"message":"Current position fetch successful.",
			"data":{
				"datetime":"2018-01-01T00:00:00Z",
				"latitude":0.0,
				"longitude":0.0
			}
		}`

		t.FailNow()
	})

	t.Run("Updating the current device position", func(t *testing.T) {
		`{
			"status":"success",
			"message":"Device position updated.",
			"data":{
				"datetime":"2018-01-01T00:00:00Z",
				"latitude":0.0,
				"longitude":0.0
			}
		}`
		t.FailNow()
	})

	t.Run("Getting the device's position history", func(t *testing.T) {
		`{
			"status":"success",
			"message":"History for device fetched successful.",
			"data": [
				{"datetime":"2018-01-01T00:00:00Z","latitude":0.0,"longitude":0.0},
				{"datetime":"2018-01-01T01:00:00Z","latitude":1.0,"longitude":-1.0},
				{"datetime":"2018-01-01T02:00:00Z","latitude":2.0,"longitude":-3.0},
				{"datetime":"2018-01-01T05:00:00Z","latitude":3.0,"longitude":-10.0},
				{"datetime":"2018-01-01T07:00:00Z","latitude":4.0,"longitude":-20.0},
			]
		}`
		t.FailNow()
	})
}
