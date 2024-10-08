package datasource_test

import (
	"testing"

	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func TestDatasource_Update(t *testing.T) {
	setupMockServer()
	defer teardownMockServer()

	dsConfig := datasource.DatabaseConfig{
		Name:        "test_datasource",
		Engine:      "postgres",
		Description: "Test description",
	}
	updatedDs, err := ds.Update(dsConfig)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if updatedDs.Name != "test_datasource" {
		t.Errorf("Expected datasource name 'test_datasource', got %s", updatedDs.Name)
	}

}
