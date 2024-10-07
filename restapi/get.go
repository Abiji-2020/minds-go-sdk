package restapi

import (
	"net/http"
)

// GET request
func (api *RestAPI) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", api.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}

	req.Header = api.headers()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := raiseForStatus(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
