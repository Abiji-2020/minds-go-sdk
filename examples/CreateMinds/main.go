package main

import (
	"fmt"

	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {
	// To initialize the client, you need to provide your API key and the base URL
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	//Declare a datasource
	ds1 := datasource.DatabaseConfig{
		Name:        "datasource1",
		Engine:      "mysql",
		Description: "This is a test datasource",
		ConnectionData: map[string]string{
			"user":     "root",
			"password": "password",
			"host":     "localhost",
			"port":     "3306",
			"database": "test",
		},
		Tables: []string{"table1", "table2"},
	}

	// Create a new mind
	datas := []datasource.DatabaseConfig{ds1}
	mindName := "Test Mind"
	createdMind, err := Client.Minds.Create(mindName, datas)
	if err != nil {
		fmt.Println("Error creating mind: ", err)
	} else {
		fmt.Println("Mind created: ", createdMind)
	}
}
