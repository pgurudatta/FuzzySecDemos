package main

import (
    "database/sql"
    "fmt"
    "net/http"
)

func handler(db *sql.DB, w http.ResponseWriter, req *http.Request) {
    category := req.URL.Query().Get("category")
    if category == "" {
        http.Error(w, "Category not provided", http.StatusBadRequest)
        return
    }

    q := fmt.Sprintf("SELECT ITEM, PRICE FROM PRODUCT WHERE ITEM_CATEGORY='%s' ORDER BY PRICE", category)
    rows, err := db.Query(q)
    if err != nil {
        http.Error(w, "Failed to execute query", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    // Process the rows (not shown in this example)

    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Query executed successfully"))
}

func main() {
    // Create a new HTTP server
    http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        // Pass the database connection to the handler
        handler(nil, w, req)
    })

    // Start the HTTP server and listen for incoming requests on port 8080
    http.ListenAndServe(":8080", nil)
}
