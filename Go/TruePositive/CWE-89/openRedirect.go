package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/redirect", handleRedirect)
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":8080", nil)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {
	// Get the redirect URL from the "url" query parameter
	redirectURL := r.URL.Query().Get("url")

	if redirectURL == "" {
		http.Error(w, "Redirect URL not specified", http.StatusBadRequest)
		return
	}

	// Perform the redirect to the specified URL
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Simulated login functionality (not relevant to the vulnerability)
	fmt.Fprintf(w, "Login page")
}
