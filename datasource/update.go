package datasource

import (
	"encoding/json"
	"errors"
)

type FilterDatabaseConfig struct {
	ConnectionData map[string]string `json:"connection_data,omitempty"`
	Tables         []string          `json:"tables,omitempty"`
	Description    string            `json:"description,omitempty"`
}

func (ds *Datasource) Update(dbConfig DatabaseConfig) (*DatabaseConfig, error) {
	filteredConfig := FilterDatabaseConfig{
		ConnectionData: dbConfig.ConnectionData,
		Tables:         dbConfig.Tables,
		Description:    dbConfig.Description,
	}

	datasourceName := dbConfig.Name
	data, _ := json.Marshal(filteredConfig)

	resp, err := ds.config.Patch("/datasources/"+datasourceName, data)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("failed on the updation of the datasource : " + datasourceName)
	}

	return &dbConfig, nil

}
