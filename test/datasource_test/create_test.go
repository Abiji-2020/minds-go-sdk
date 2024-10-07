package datasource_test

import (
	"testing"

	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func TestDatasource_create(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	dsConfig := datasource.DatabaseConfig{
		Name:        "test_datasource",
		Engine:      "postgres",
		Description: "Test description",
	}

	createdDs, err := ds.Create(dsConfig, false)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdDs.Name != "test_datasource" {
		t.Errorf("Expected datasource name to be 'test_datasource', got %s", createdDs.Name)
	}
}
