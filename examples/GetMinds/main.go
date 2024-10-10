package main

import (
	"fmt"

	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {
	// To initialize the client, you need to provide your API key and the base URL
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	// Get a mind by Name

	mindName := "Test Mind"
	Mind, err := Client.Minds.Get(mindName)
	if err != nil {
		fmt.Println("Error getting mind: ", err)
	} else {
		fmt.Println("Mind: ", Mind.Name)
	}
}
