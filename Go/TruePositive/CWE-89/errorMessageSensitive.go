package main

import (
	"database/sql"
	"fmt"
)

func getUser(userID int) (string, error) {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
	if err != nil {
		return "", fmt.Errorf("Error connecting to database: %v", err) // Leaks database credentials in error message
	}
	defer db.Close()

	// ... (code to query the database using userID)

	return username, nil
}

func main() {
	username, err := getUser(123)
	if err != nil {
		fmt.Println(err) // Error message might be displayed to user
	} else {
		fmt.Println("Username:", username)
	}
}
