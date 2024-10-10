package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", "BASE URL")

	// Define the database Config to create the datasource
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

	// Create the datasource with the database config and replace flag to true if you want to replace the existing datasource
	createdDs, err := Client.Datasource.Create(ds1, false)
	if err != nil {
		fmt.Println("Error creating datasource: ", err)
	} else {
		fmt.Println("Datasource created successfully: ", createdDs.Name)
	}
}
