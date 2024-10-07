package datasource

import (
	"encoding/json"
	"errors"
)

func (ds *Datasource) List() ([]DatabaseConfig, error) {

	resp, err := ds.config.Get("/datasources")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var dbConfigs []DatabaseConfig

	if err := json.NewDecoder(resp.Body).Decode(&dbConfigs); err != nil {
		return nil, err
	}

	// Filter the list to exclude datasources without an engine

	filtered := []DatabaseConfig{}
	for _, item := range dbConfigs {
		if item.Engine == "" {
			// Skip items where engine is empty
			continue
		}
		filtered = append(filtered, item)
	}

	if len(filtered) == 0 {
		return nil, errors.New("no valid datasource found")
	}

	return filtered, nil
}
