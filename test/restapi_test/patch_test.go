package restapi_test

import (
	"net/http"
	"testing"
)

// Test for Patch method using the shared mock server
func TestRestAPI_Patch(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	data := map[string]string{"key": "updatedValue"}

	resp, err := api.Patch("/test-patch", data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
}
