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

	id := 1011334
	fmt.Printf("Fetching Details of character whose Id is %+v...\n", id)
	char, err := client.CharacterDetails.GetById(id)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("Character: %+v\n", char)
}
