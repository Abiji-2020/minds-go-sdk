package datasource_test

import (
	"testing"
)

func TestDAtasource_Drop(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	err := ds.Drop("test_datasource")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
