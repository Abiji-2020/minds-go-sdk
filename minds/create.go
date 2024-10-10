package minds

import (
	"encoding/json"
	"errors"

	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func (md *Minds) Create(name string, Datasources []datasource.DatabaseConfig) (string, error) {
	var NotAvailable = errors.New("datasource not found")
	DatasourceNames := []string{}
	datasource := datasource.NewDatasource(md.Config)
	for _, ds := range Datasources {

		_, err := datasource.Get(ds.Name)
		if err == NotAvailable {
			datasource.Create(ds, false)
		}
		DatasourceNames = append(DatasourceNames, ds.Name)
	}
	data := CreateRequest{
		Name:            name,
		DatasourceNames: DatasourceNames,
	}

	jsonData, _ := json.Marshal(data)

	resp, err := md.Config.Post("/projects/mindsdb/minds", jsonData)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("failed to create a new mind : " + name)
	}

	return name, nil

}
