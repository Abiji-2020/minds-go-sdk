package minds

import (
	"encoding/json"
	"errors"

	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func (md *Minds) AddDatasource(name string, Datasource datasource.DatabaseConfig, checkConnection bool) error {

	if name == "" {
		return errors.New("Name cannot be empty")
	}

	req := AddDatasourceRequest{
		Name:            Datasource.Name,
		CheckConnection: checkConnection,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	_, err = md.config.Post("/projects/mindsdb/minds/"+name+"/datasources", body)
	if err != nil {
		return err
	}

	return nil

}
