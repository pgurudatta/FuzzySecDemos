package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url") // User-controlled input

	// Vulnerable: User input directly used to construct redirect URL
	http.Redirect(w, r, url, http.StatusFound)
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
