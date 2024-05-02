package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// UserData represents the XML structure for user data
type UserData struct {
	Username string `xml:"username"`
	Password string `xml:"password"`
}

// getUsers queries the XML database for user data using XPath
func getUsers(query string) ([]UserData, error) {
	// Simulated XML database containing user data
	xmlData := `
		<users>
			<user>
				<username>john_doe</username>
				<password>secretpassword123</password>
			</user>
			<user>
				<username>alice</username>
				<password>alicepassword456</password>
			</user>
		</users>
	`

	// Parse the XML data
	var users []UserData
	err := xml.Unmarshal([]byte(xmlData), &users)
	if err != nil {
		return nil, err
	}

	// Simulate query execution using XPath (vulnerable to injection)
	var results []UserData
	for _, user := range users {
		// Vulnerable XPath query execution
		if user.Username == query || user.Password == query {
			results = append(results, user)
		}
	}

	return results, nil
}

func main() {
	// Simulated user input (XPath query)
	query := "' or '1'='1" // Simulated XPath injection attack

	// Retrieve user data based on the XPath query
	users, err := getUsers(query)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Display the retrieved user data
	if len(users) == 0 {
		fmt.Println("No matching users found")
	} else {
		fmt.Println("Matching users:")
		for _, user := range users {
			fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
		}
	}
}
