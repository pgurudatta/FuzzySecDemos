package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Connect to the database
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/dbname")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Vulnerable query using string concatenation
    username := "admin'; DROP TABLE users--"
    password := "password123"
    query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)

    // Execute the query
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Process the query results
    for rows.Next() {
        var userID int
        var username string
        var password string
        if err := rows.Scan(&userID, &username, &password); err != nil {
            log.Fatal(err)
        }
        fmt.Printf("User found: ID=%d, Username=%s, Password=%s\n", userID, username, password)
    }
}
