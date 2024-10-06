package restapi_test

import (
	"net/http"
	"testing"
)

// Test for Post method using the shared mock server
func TestRestAPI_Post(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	data := map[string]string{"key": "value"}

	resp, err := api.Post("/test-post", data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
	}
}
