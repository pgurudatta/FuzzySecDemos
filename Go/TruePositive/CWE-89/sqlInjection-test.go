package main

import (
	"database/sql"
	"fmt"
)

func getUser(username string) (string, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		return "", err
	}
	defer db.Close()

	// Vulnerable: User input directly concatenated into SQL query
	query := fmt.Sprintf("SELECT email FROM users WHERE username = '%s'", username)
	row := db.QueryRow(query)

	var email string
	err = row.Scan(&email)
	if err != nil {
		return "", err
	}

	return email, nil
}

func main() {
	username := "hacker" // User-controlled input
	email, err := getUser(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Email:", email)
}
