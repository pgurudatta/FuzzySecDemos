package main

import (
	"fmt"
	"os"
	"strings"
)

type User struct {
	ID       string
	Password string
}
var userDatabase = map[string]User{
	"john_doe": {"john_doe", "secretpassword123"},
	"jane_doe": {"jane_doe", "anothersecretpassword"},
}

func main() {
	userID := "john_doe"
  password := "secretpassword123"
	// Authenticate user
	if authenticate(userID, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed. Invalid credentials.")
	}
}

func authenticate(userID, password string) bool {
	log.Printf("Attempting login with userID: %s, password: %s", userID, password)
	user, ok := userDatabase[userID]
	if !ok {
		// User not found
		return false
	}

	// Check if the provided password matches the stored password
	return user.Password == password
}
