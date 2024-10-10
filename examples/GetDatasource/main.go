package main 

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {
	
	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	dataSourceName := "datasource1"

	// Get the datasource
	ds, err := Client.Datasource.Get(dataSourceName)
	if err != nil {
		fmt.Println("Error in getting datasource: ", err)
	} else {
		fmt.Println("Datasource Name: ", ds.Name)
	}
}