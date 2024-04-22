package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "secretpassword123"

	// Generate a hash with a low cost factor (vulnerable)
	hash, err := generateHash(password, bcrypt.MinCost)
	if err != nil {
		log.Fatalf("Failed to generate hash: %v", err)
	}

	fmt.Printf("Generated Hash: %s\n", hash)

	// Validate the password against the generated hash
	err = validatePassword(password, hash)
	if err != nil {
		fmt.Printf("Password validation failed: %v\n", err)
	} else {
		fmt.Println("Password validation successful")
	}
}

// generateHash generates a bcrypt hash with the specified cost factor
func generateHash(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// validatePassword compares the provided password with the bcrypt hash
func validatePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
