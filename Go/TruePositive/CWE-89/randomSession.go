package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	// Generate a random session token using math/rand (INSECURE - DO NOT USE FOR CRYPTOGRAPHIC PURPOSES)
	sessionToken := generateSessionToken()
	fmt.Printf("Generated Session Token: %s\n", sessionToken)
}

func generateSessionToken() string {
	const tokenLength = 16 // Length of session token

	// Seed math/rand with current timestamp (NOT SUITABLE FOR CRYPTOGRAPHIC USE)
	// This is insecure and can lead to predictable values if not properly seeded
	// Do not use math/rand for security-sensitive purposes
	// Example only - DO NOT USE IN PRODUCTION
	rand.Seed(123456789) // Insecure seed value (example only)

	// Characters to use in the session token (simplified for demonstration)
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	token := make([]byte, tokenLength)
	for i := range token {
		token[i] = charset[rand.Intn(len(charset))]
	}

	return string(token)
}
