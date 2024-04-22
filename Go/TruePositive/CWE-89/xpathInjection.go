package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Example XML data structure
type User struct {
	XMLName  xml.Name `xml:"user"`
	Username string   `xml:"username"`
	Password string   `xml:"password"`
}

func main() {
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Get username and password from request parameters
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Construct XPath query using user-supplied data
	xpathQuery := fmt.Sprintf("/users/user[username/text()='%s' and password/text()='%s']", username, password)

	// Simulate querying XML data based on the constructed XPath query
	xmlData := `<users>
					<user>
						<username>admin</username>
						<password>strongPassword123</password>
					</user>
					<user>
						<username>john</username>
						<password>password123</password>
					</user>
			   </users>`

	// Unmarshal XML data into User struct
	var users []User
	err := xml.Unmarshal([]byte(xmlData), &users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Find user based on XPath query
	var foundUser User
	for _, user := range users {
		if matchesXPath(user, xpathQuery) {
			foundUser = user
			break
		}
	}

	// Validate if user was found
	if foundUser.Username != "" {
		fmt.Fprintf(w, "Login successful for user: %s\n", foundUser.Username)
	} else {
		fmt.Fprintf(w, "Login failed: User not found\n")
	}
}

// Function to check if the given user matches the provided XPath query
func matchesXPath(user User, xpathQuery string) bool {
	// Use strings.Contains for a simple check (vulnerable to injection)
	return strings.Contains(xpathQuery, user.Username) && strings.Contains(xpathQuery, user.Password)
}
