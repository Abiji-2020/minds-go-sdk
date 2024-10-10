package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR_API_KEY", "BASE_URL")

	mindName := "Test Mind"

	// Need to provide the Database Config to update the Mind
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

	// Add the datasource to the mind with the name mindName set the flag to check the connection to database
	err := Client.Minds.AddDatasource(mindName, ds1, true)
	if err != nil {
		fmt.Println("Error adding datasource: ", err)
	} else {
		fmt.Println("Datasource added successfully")
	}
}
