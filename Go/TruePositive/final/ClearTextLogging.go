/*
Sample code for vulnerable type: Clear Text Logging
CWE : CWE-200, CWE-312
Description : Exposure of Sensitive Information to an Unauthorized Actor, Cleartext Storage of Sensitive Information
*/
package main

import (
	"fmt"
	"log"
)

func main() {
	username := "john_doe"
	password := "secretpassword123" //source

	// Log sensitive information in clear text
	log.Printf("Attempting login with username: %s, password: %s", username, password) //sink

	// Simulate login logic (for demonstration purposes)
	if authenticate(username, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed. Invalid credentials.")
	}
}

func authenticate(username, password string) bool {
	// Perform authentication logic (e.g., check username and password against database)
	// This is a placeholder function for demonstration
	return username == "john_doe" && password == "secretpassword123"
}
