package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariel17/xy/api/controllers"
	"github.com/ariel17/xy/api/dao"
	"github.com/ariel17/xy/api/domain"
	"github.com/julienschmidt/httprouter"
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

			resp, err := doRequest("GET", "/users/abc123", nil, controllers.GetUsers)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			if resp.Code != http.StatusOK {
				t.Errorf("status missmatch! expected %d, got %d: %v", http.StatusOK, resp.Code, err)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			// if body := resp.Body.String(); body != tc.expectedBody {
			// 	t.Errorf("body missmatch! expected %s, got %s", tc.expectedBody, body)
			// 	t.FailNow()
			// }
		})

		t.Run("NotFound", func(t *testing.T) {
			defer dao.CleanMocks()
			resp, err := doRequest("GET", "/users/9999", nil, controllers.GetUsers)
			if err != nil {
				t.Error(err)
				t.FailNow()
			}

			if resp.Code != http.StatusNotFound {
				t.Errorf("status missmatch! expected %d, got %d: %v", http.StatusOK, resp.Code, err)
				t.FailNow()
			}

			if resp.Body == nil {
				t.Error("body is nil")
				t.FailNow()
			}

			// if body := resp.Body.String(); body != tc.expectedBody {
			// 	t.Errorf("body missmatch! expected %s, got %s", tc.expectedBody, body)
			// 	t.FailNow()
			// }
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

func doRequest(method, uri string, body *bytes.Buffer, handle httprouter.Handle) (*httptest.ResponseRecorder, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	router := httprouter.New()
	router.Handle(method, uri, handle)
	router.ServeHTTP(resp, req)
	return resp, nil
}