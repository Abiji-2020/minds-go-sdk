package datasource

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// Get retrieves a datasource by name from the API.
func (ds *Datasource) Get(name string) (*DatabaseConfig, error) {

	resp, err := ds.config.Get("/datasources/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("datasource not found")
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the body of the response
	bodyBytes, err := io.ReadAll(resp.Body) // Read the body before closing
	if err != nil {
		return nil, err
	}

	// Decode the response body into DatabaseConfig
	var dbConfig DatabaseConfig
	if err := json.Unmarshal(bodyBytes, &dbConfig); err != nil {
		return nil, err
	}

	// Check if the engine field is empty
	if dbConfig.Engine == "" {
		return nil, errors.New("wrong type of datasource")
	}

	return &dbConfig, nil
}
