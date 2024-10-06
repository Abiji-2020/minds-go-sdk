package restapi_test

import (
	"github.com/Abiji-2020/minds-go-sdk/restapi"
	"net/http"
	"net/http/httptest"
	
)

var (
	mockServer *httptest.Server
	api        *restapi.RestAPI
)

func setupMockServer() {
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/test-get":
			if r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message": "get success"}`))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		case "/api/test-post":
			if r.Method == http.MethodPost {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"message":"post success"}`))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		case "/api/test-delete":
			if r.Method == http.MethodDelete {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"delete success"}`))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		case "/api/test-patch":
			if r.Method == http.MethodPatch {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"message":"patch success"}`))
			} else {
				w.WriteHeader(http.StatusMethodNotAllowed)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	api = restapi.NewRestAPI("1234567890abc", mockServer.URL)
}

func teardownMockServer() {
	mockServer.Close()
}
