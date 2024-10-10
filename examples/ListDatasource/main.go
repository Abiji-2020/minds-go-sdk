package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	ds, err := Client.Datasource.List()
	if err != nil {
		fmt.Println("Error in listing datasources: ", err)
	} else {
		for _, datasource := range ds {
			fmt.Println("Datasource Name: ", datasource.Name)
		}
	}
}
