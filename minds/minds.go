package minds

import (
	"github.com/Abiji-2020/minds-go-sdk/restapi"
)

type Minds struct {
	Config *restapi.RestAPI
}

type ResponseMind struct {
	CreatedAt  string            `json:"created_at"`
	DsNames    []string          `json:"datasources"`
	ModelName  string            `json:"model_name"`
	Name       string            `json:"name"`
	Parameters map[string]string `json:"parameters"`
	Provider   string            `json:"provider"`
	UpdatedAt  string            `json:"updated_at"`
}
type CreateRequest struct {
	Name            string   `json:"name"`
	DatasourceNames []string `json:"datasources"`
}

type AddDatasourceRequest struct {
	Name            string `json:"name"`
	CheckConnection bool   `json:"check_connection"`
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int64          `json:"created"`
	Model   string         `json:"model"`
	Choices []StreamChoice `json:"choices"`
}

type StreamChoice struct {
	Delta MessageDelta `json:"delta"`
	Index int          `json:"index"`
}

type MessageDelta struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

func NewMinds(api *restapi.RestAPI) *Minds {
	return &Minds{
		Config: api,
	}
}
