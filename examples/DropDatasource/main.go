package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	dataSourceName := "datasource1"

	// Drop the datasource
	err := Client.Datasource.Drop(dataSourceName)
	if err != nil {
		fmt.Println("Error in dropping datasource: ", err)
	} else {
		fmt.Println("Datasource dropped successfully: ", dataSourceName)
	}
}
