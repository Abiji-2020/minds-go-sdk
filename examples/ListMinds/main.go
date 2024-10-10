package main

import (
	"fmt"

	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {
	// To initialize the client, you need to provide your API key and the base URL
	Client := client.NewClient("YOUR_API_KEY", "BASE_URL")

	// List all minds
	Minds, err := Client.Minds.List()
	if err != nil {
		fmt.Println(err)
	}

	// Print the name of all the Minds

	for _, mind := range Minds {
		fmt.Println(mind.Name)
	}
}
