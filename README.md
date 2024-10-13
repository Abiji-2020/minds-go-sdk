[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Abiji-2020/minds-go-sdk/go.yml)](https://github.com/Abiji-2020/minds-go-sdk/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Abiji-2020/minds-go-sdk)](https://goreportcard.com/report/github.com/Abiji-2020/minds-go-sdk)
[![GitHub License](https://img.shields.io/github/license/Abiji-2020/minds-go-sdk)](https://github.com/Abiji-2020/minds-go-sdk/blob/main/LICENSE)


# Minds `Go` SDK

The `Minds Go SDK` is a Software Development Kit for the [Minds](https://mindsdb.com/minds) (Unofficial). Built with the REST API Endpoints given by the Minds. With the help of this the  `Go Developers` can able to easily interact with the Minds. Currently the Minds SDK allows the developer to `Create`, `Update`, `Add Datasource`, `Delete` , `List` and `Get` the Minds, and Datasource. As of now it uses the Http package from the `go` to communicate with the Endpoints of Minds. 

## Features
- **Mind Management:**  Create, Update, Delete and Fetch details of Minds from MindsDB and also has the feature to add Datasource to the existing Mind.
- **Datasource Management:** Create, Update, Delete and Fetch details of Datasources linked to MindsDB.
- **Chat Completeion:** Integreate with chat applications API for natural language processing, it can support both single chat and streaming of chats.

## Getting Started
   <h3>Prerequisites</h3>
   Before you can use the SDK, you'll need:
   - A valid Minds API key (can be accessed from [Minds](https://mdb.ai/login)
   - Go installed in your machine.
   - A project set up in Go or you can create a new one. 
 
  ### Installation
  To install the SDK, run the following command in the terminal where your project is located.
  ```bash
go get github.com/Abiji-2020/minds-go-sdk
```
Then you can use the modules in the project. 

## Usage 

Below is a simple example of how to use the `Client` package to interact with the Minds API. [More Examples on each feature](https://github.com/Abiji-2020/minds-go-sdk/tree/main/examples)

```go
package main

import (
	"fmt"

	"github.com/Abiji-2020/minds-go-sdk/client"
	"github.com/Abiji-2020/minds-go-sdk/datasource"
)

func main() {
	// To initialize the client, you need to provide your API key and the base URL
	Client := client.NewClient("YOUR API KEY", " BASE URL")

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
	mindName := "Test Mind"
	createdMind, err := Client.Minds.Create(mindName, datas)
	if err != nil {
		fmt.Println("Error creating mind: ", err)
	} else {
		fmt.Println("Mind created: ", createdMind)
	}
}
```

## Features
1. Minds Management
   - Fetch all minds
   - Get a specific mind by name
   - Add datasource to mind
   - Create, Update and delete minds
     
2. Datasource Management
   - List all datasources
   - Get a datasource by name
   - Create, update and delete a datasource from minds.

3. Completions API
   - Use natural language processing with chat completions.
  
## Additional Information

This SDK is currently under active development. Contributions are welcome.

> [!NOTE]
> For the future of the `SDK` , I have planned to add `templates` as a package of giving out all the [Datasource templates in the Minds](https://docs.mdb.ai/docs/data_sources).
>
 - If found any issues report them, or add them in `Issues` tab
 - For any inquiries and feature requests, feel free to open a discussion or contact us.

## License
This package is Licenced under the *MIT LICENSE*, see the `LICENSE` file for more details. 
