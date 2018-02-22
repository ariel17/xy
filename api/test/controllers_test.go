package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel17/xy/api/api"
	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
)

func TestControllers(t *testing.T) {

	t.Run("GetUser", func(t *testing.T) {

		t.Run("OK", func(t *testing.T) {
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

			expectedBody := `{"_id":"616263313233","nick":"ariel17"}`
			if body := strings.TrimSpace(resp.Body.String()); body != expectedBody {
				t.Errorf("body mismatch! expected %s, got %s", expectedBody, body)
				t.FailNow()
			}
		})

		t.Run("NotFound", func(t *testing.T) {
			defer dao.CleanMocks()

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

	// testCases := []struct {
	// 	name           string
	// 	handle         httprouter.Handle
	// 	path           string
	// 	userID         string
	// 	method         string
	// 	body           string
	// 	expectedStatus int
	// 	expectedBody   string
	// }{
	// 	{"GetUsersOK", controllers.GetUsers, "/users/abc123", "abc123", "GET", "", http.StatusOK, "{}"},
	// 	{"GetUsersNotFound", controllers.GetUsers, "/users/abc123", "abc123", "GET", "", http.StatusNotFound, "{}"},
	// }

	// for _, tc := range testCases {
	// 	t.Run(tc.name, func(t *testing.T) {

	// 		body := bytes.NewBufferString(tc.body)
	// 		resp, err := doRequest(tc.method, tc.path, body, tc.handle)
	// 		if err != nil {
	// 			t.Error(err)
	// 			t.FailNow()
	// 		}

	// 		if resp.Code != tc.expectedStatus {
	// 			t.Errorf("status missmatch! expected %d, got %d: %v", tc.expectedStatus, resp.Code, err)
	// 			t.FailNow()
	// 		}

	// 		if resp.Body == nil {
	// 			t.Error("body is nil")
	// 			t.FailNow()
	// 		}

	// 		if body := resp.Body.String(); body != tc.expectedBody {
	// 			t.Errorf("body missmatch! expected %s, got %s", tc.expectedBody, body)
	// 			t.FailNow()
	// 		}
	// 	})
	// }
}
