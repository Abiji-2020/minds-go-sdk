package datasource

import (
	"encoding/json"
	"errors"
)

func (ds *Datasource) Create(dbConfig DatabaseConfig, replace bool) (*DatabaseConfig, error) {
	if replace {
		if _, err := ds.Get(dbConfig.Name); err == nil {
			ds.Drop(dbConfig.Name)
		}
	}

	data, _ := json.Marshal(dbConfig)
	resp, err := ds.config.Post("/datasources", data)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to create datasource")
	}
	return &dbConfig, nil

}
