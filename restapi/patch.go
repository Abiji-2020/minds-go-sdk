package restapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (api *RestAPI) Patch(url string, data interface{}) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PATCH", api.BaseURL+url, bytes.NewBuffer(jsonData))
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
