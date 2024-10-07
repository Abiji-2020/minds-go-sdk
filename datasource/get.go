package datasource

import (
	"encoding/json"
	"errors"
)

func (ds *Datasource) Get(name string) (*DatabaseConfig, error) {

	resp, err := ds.config.Get("/datasources/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, errors.New("datasource not found")
	}

	var dbConfig DatabaseConfig
	if err := json.NewDecoder(resp.Body).Decode(&dbConfig); err != nil {
		return nil, err
	}

	if dbConfig.Engine == "" {
		return nil, errors.New("wrong type of datsource")
	}

	return &dbConfig, nil

}
