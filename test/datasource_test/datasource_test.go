package datasource_test

import (
	"github.com/Abiji-2020/minds-go-sdk/datasource"
	"github.com/Abiji-2020/minds-go-sdk/restapi"

	"net/http"
	"net/http/httptest"
)

var (
	mockServer *httptest.Server
	api        *restapi.RestAPI
	ds         *datasource.Datasource
)

func setupMockServer() {
	mockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/datasources":
			if r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`[{"name":"test_datasource","engine":"postgres"}]`))
			} else if r.Method == http.MethodPost {
				w.WriteHeader(http.StatusCreated)
				w.Write([]byte(`{"name":"test_datasource", "engine":"postgres"}`))
			}
		case "/api/datasources/test_datasource":
			if r.Method == http.MethodGet {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"name":"test_datasource", "engine":"postgres"}`))
			} else if r.Method == http.MethodDelete {
				w.WriteHeader(http.StatusNoContent)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	api = restapi.NewRestAPI("1234567890abc", mockServer.URL)
	ds = datasource.NewDatasource(api)
}

func teardownMockServer() {
	mockServer.Close()
}
