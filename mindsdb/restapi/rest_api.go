package restapi

import (
	"errors"
	"fmt"

	"net/http"
	"strings"
	"time"
)

// RestAPI struct to handle the HTTP Requests
type RestAPI struct {
	APIKey  string
	BaseURL string
	Client  *http.Client
}

// NewRestAPI creates a  new RestAPI instance

func NewRestAPI(apiKey, baseURL string) *RestAPI {
	if baseURL == "" {
		baseURL = "https://mdb.ai"
	}

	// Remove any trailing slashes from the baseURL to avoid issues with the double slashes
	baseURL = strings.TrimRight(baseURL, "/")

	// check if '/api' is already present at the end of the baseURL
	if !strings.HasSuffix(baseURL, "/api") {
		baseURL = baseURL + "/api"
	}

	return &RestAPI{
		APIKey:  apiKey,
		BaseURL: baseURL,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Helper function to set headers

func (api *RestAPI) headers() http.Header {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+api.APIKey)
	headers.Set("Content-Type", "application/json")
	return headers
}

// Helper function to raise error for status

func raiseForStatus(resp *http.Response) error {
	if resp.StatusCode == 404 {
		return errors.New("Object Not found")
	}

	if resp.StatusCode == 403 {
		return errors.New("Forbidden")
	}

	if resp.StatusCode == 401 {
		return errors.New("Unauthorized")

	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("Unknown error: %s", resp.Status)
	}
	return nil
}
