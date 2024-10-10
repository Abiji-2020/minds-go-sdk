package main

import (
	"fmt"

	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR_API_KEY", "BASE_URL")

	// Update a Mind we need old Mind Name and new Mind Name and List of Datasources

	oldMindName := "Test Mind"
	newMindName := "Updated Mind"

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
	updatedMind, err := Client.Minds.Update(oldMindName, newMindName, datas)
	if err != nil {
		fmt.Println("Error updating mind: ", err)
	} else {
		fmt.Println("Mind updated: ", updatedMind)
	}
}
