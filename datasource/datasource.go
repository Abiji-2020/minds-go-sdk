package datasource

import (
	"github.com/Abiji-2020/minds-go-sdk/restapi"
)

type DatabaseConfig struct {
	Name           string            `json:"name"`
	Engine         string            `json:"engine"`
	Description    string            `json:"description"`
	ConnectionData map[string]string `json:"connection_data,omitempty"`
	Tables         []string          `json:"tables,omitempty"`
}

type Datasource struct {
	config *restapi.RestAPI
}

func NewDatasource(api *restapi.RestAPI) *Datasource {
	return &Datasource{
		config: api,
	}
}
