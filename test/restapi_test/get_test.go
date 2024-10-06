package restapi_test

import (
	"net/http"
	"testing"
)

// Test for Get method using the shared mock server
func TestRestAPI_Get(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	resp, err := api.Get("/test-get")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
