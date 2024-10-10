package minds

import (
	"encoding/json"
	"errors"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func (md *Minds) Update(Oldname, newName string, Datasources []datasource.DatabaseConfig) (string, error) {
	datasourceName := []string{}
	for _, ds := range Datasources {
		datasourceName = append(datasourceName, ds.Name)
	}
	data := CreateRequest{
		Name:            newName,
		DatasourceNames: datasourceName,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := md.Config.Patch("/projects/mindsdb/minds/"+Oldname, jsonData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("cannot update the minds" + Oldname)
	}
	return newName, nil

}
