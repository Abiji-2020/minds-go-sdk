package minds

import (
	"github.com/Abiji-2020/minds-go-sdk/restapi"
)

type Minds struct {
	config *restapi.RestAPI
}

type ResponseMind struct {
	CreatedAt  string            `json:"created_at"`
	dsNames    []string          `json:"datasources"`
	ModelName  string            `json:"model_name"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
	Provider   string            `json:"provider"`
	UpdatedAT  string            `json:"updated_at"`
}
type CreateRequest struct {
	Name            string   `json:"name"`
	datasourceNames []string `json:"datasources"`
}

func NewMinds(api *restapi.RestAPI) *Minds {
	return &Minds{
		config: api,
	}
}
