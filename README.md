# go-marvel
Golang client for [Marvel APIs](https://developer.marvel.com/docs).

### Overview

`go-marvel` is a package which helps you in writing apps that need to call REST APIs of Marvel Comics site


### Usage

```
go get github.com/amitgurav04/go-marvel
```

Some example code:

```go
package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/amitgurav04/go-marvel"
)

func main() {
	fmt.Println("This is Example of Marvel client in Go.")
	publicKey := os.Getenv("PUBLIC_KEY")
	privateKey := os.Getenv("PRIVATE_KEY")
	if publicKey == "" || privateKey == "" {
		fmt.Println("Set PUBLIC_KEY and PRIVATE_KEY for Marvel as an environment variables.")
		os.Exit(1)
	}
	httpClient := &http.Client{}
	auth := marvel.NewAuth(publicKey, privateKey)

	client := marvel.NewClient(auth, httpClient)
	fmt.Printf("Marvel Client is: %+v\n\n", client)

	fmt.Println("Fetching Details of all characters...")
	chars, err := client.CharacterDetails.ListAll()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Details of all Characters: %+v\n\n", chars)

	id := 1011334 //Id to fetch the details of the character
	fmt.Printf("Fetching Details of character whose Id is %+v...\n", id)
	char, err := client.CharacterDetails.GetById(id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Character: %+v\n", char)
}
```

How to run sample code:

```
PUBLIC_KEY=<public_key> PRIVATE_KEY=<private_key> go run main.go
```
Note: refer [Marvel Account](https://developer.marvel.com/account) page to get public_key and private_key for Marvel Comics API.