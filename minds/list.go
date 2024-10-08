package minds

import (
	"encoding/json"
)

func (md *Minds) List() ([]ResponseMind, error) {
	resp, err := md.config.Get("/projects/mindsdb/minds")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var mindResponses []ResponseMind
	if err := json.NewDecoder(resp.Body).Decode(mindResponses); err != nil {
		return nil, err
	}
	return mindResponses, nil
}
