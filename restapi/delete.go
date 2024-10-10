package restapi

import (
	"net/http"
)

func (api *RestAPI) Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", api.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header = api.Headers()

	resp, err := api.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if err := raiseForStatus(resp); err != nil {
		return nil, err
	}
	return resp, nil
}
