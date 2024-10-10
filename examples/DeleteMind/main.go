package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	mindName := "test_mind"

	// Delete the mind
	err := Client.Minds.Delete(mindName)
	if err != nil {
		fmt.Println("Failed to delete mind: ", mindName)
		return
	}
	fmt.Println("Mind deleted successfully: ", mindName)
}
