package main

import (
	"fmt"
	"net/http"
	"xml/xpath"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")

	// Vulnerable: User input directly concatenated into XPath expression
	expr := fmt.Sprintf("user/name[text() = '%s']", username)

	// ... (code to process XML document using the vulnerable expression)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
