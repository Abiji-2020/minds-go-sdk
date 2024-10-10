package client

import (

	"github.com/Abiji-2020/minds-go-sdk/datasource"
	"github.com/Abiji-2020/minds-go-sdk/restapi"
	"github.com/Abiji-2020/minds-go-sdk/minds"
)

const defaultBaseURL = "https://mdb.ai"

type Client struct {
	Minds *minds.Minds
	Datasource *datasource.Datasource
}

func NewClient(apiKey string, baseURL ...string ) *Client {
	finalBaseURL := defaultBaseURL
	if len(baseURL) > 0 && baseURL[0] != "" {
		finalBaseURL = baseURL[0]
	}

	apiConfig := restapi.NewRestAPI(apiKey, finalBaseURL)

	mindsInstance := minds.NewMinds(apiConfig)

	return &Client{
		Minds: mindsInstance,
		Datasource: datasource.NewDatasource(apiConfig),
	}
}