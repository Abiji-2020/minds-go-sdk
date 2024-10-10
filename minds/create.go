package minds

import (
	"encoding/json"
	"errors"

	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func (md *Minds) Create(name string, Datasources []datasource.DatabaseConfig) (string, error) {
	DatasourceNames := []string{}
	for _, ds := range Datasources {
		DatasourceNames = append(DatasourceNames, ds.Name)
	}
	data := CreateRequest{
		Name:            name,
		DatasourceNames: DatasourceNames,
	}

	jsonData, _ := json.Marshal(data)

	resp, err := md.config.Post("/projects/mindsdb/minds", jsonData)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("failed to create a new mind : " + name)
	}

	return name, nil

}
