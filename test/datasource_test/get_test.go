package datasource_test

import (
	"testing"
)

func TestDatasource_Get(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	ds, err := ds.Get("test_datasource")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if ds.Name != "test_datasource" {
		t.Errorf("Expected datasource name to be 'test_datasource', got %s", ds.Name)
	}
}
