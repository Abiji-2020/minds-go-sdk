package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", "BASE URL")

	// Define the database Config to Where we need to update

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

	// Update the datasource with the database config
	updatedDs, err := Client.Datasource.Update(ds1)

	if err != nil {
		fmt.Println("Error updating datasource: ", err)
	} else {
		fmt.Println("Datasource updated successfully: ", updatedDs.Name)
	}
}
