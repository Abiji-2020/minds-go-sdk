package main

import (
	"fmt"
	"github.com/Abiji-2020/minds-go-sdk/client"
)

func main() {

	// Initalize the Client
	Client := client.NewClient("YOUR API KEY", " BASE URL")

	// Define the model and message for chat completion

	modelName := "gpt-3.5-turbo"
	userMessage := "Hello, How can you assist me ?"

	// 1. For Non-Streamed Completion
	response, _, err := Client.Minds.Completion(modelName, userMessage, false)
	if err != nil {
		fmt.Println("Error in non-streaming completion: ", err)
	} else {
		fmt.Println("Non-streaming  Response: ", response)
	}

	// 2. For Streamed Completion
	_, stream, err := Client.Minds.Completion(modelName, userMessage, true)
	if err != nil {
		fmt.Println("Error in streaming completion: ", err)
	} else {
		fmt.Println("Streaming Response: ")
		for message := range stream {
			fmt.Println(message)
		}
		fmt.Println("Streaming Completed")
	}
}
