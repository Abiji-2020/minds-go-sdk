package datasource_test

import (
	"testing"
)

func TestDatasource_List(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	dsList, err := ds.List()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(dsList) != 1 {
		t.Errorf("Expected 1 datasource, got %d", len(dsList))
	}
}
