package minds

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (md *Minds) Get(name string) (*ResponseMind, error) {

	resp, err := md.Config.Get("/projects/mindsdb/minds/" + name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("mind is not found")
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var mindValue ResponseMind
	if err := json.NewDecoder(resp.Body).Decode(&mindValue); err != nil {
		return nil, err
	}

	return &mindValue, nil

}
