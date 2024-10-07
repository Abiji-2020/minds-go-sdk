package datasource

import (
	"errors"
)

func (ds *Datasource) Drop(name string) error {
	resp, err := ds.config.Delete("/datasources/" + name)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("failed to delete datasource")
	}

	return nil
}
