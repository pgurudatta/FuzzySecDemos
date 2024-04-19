package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	username := "john_doe"
	password := "secretpassword123"

	err := authenticate(username, password)
	if err != nil {
		log.Printf("Error occurred: %s", err.Error()) // Logging error message containing sensitive information
		fmt.Println("Login failed. Please try again.")
		return
	}

	fmt.Println("Login successful!")
}

func authenticate(username, password string) error {
	// Simulate authentication logic (e.g., check username and password against database)
	// This is a placeholder function for demonstration
	expectedUsername := "john_doe"
	expectedPassword := "secretpassword123"

	if username != expectedUsername || password != expectedPassword {
		return errors.New("invalid credentials: username or password is incorrect")
	}

	return nil
}
