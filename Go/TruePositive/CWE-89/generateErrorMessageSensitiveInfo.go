package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	password = "secretpassword123" // Simulated sensitive information
)

func generateErrorMessage(username string) string {
	// Simulated error message generation
	return fmt.Sprintf("Error: authentication failed for user %s with password %s", username, password)
}

func main() {
	// Simulated multiple requests
	usernames := []string{"john_doe", "alice", "bob"}

	var wg sync.WaitGroup
	wg.Add(len(usernames))

	for _, username := range usernames {
		go func(u string) {
			defer wg.Done()

			// Simulated authentication check
			if u != "john_doe" {
				// Generate error message containing sensitive information
				errMsg := generateErrorMessage(u)

				// Print error message
				fmt.Println(errMsg)

				// Log error message to file (simulated)
				f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					fmt.Println("Failed to log error:", err)
					return
				}
				defer f.Close()
				if _, err := f.WriteString(errMsg + "\n"); err != nil {
					fmt.Println("Failed to log error:", err)
				}
			}
		}(username)
	}

	wg.Wait()
}
