package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/login", handleLogin)
	http.ListenAndServe(":8080", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Simulated login functionality (authenticate user)
	authenticated := true

	if authenticated {
		// Set a sensitive cookie without the "HttpOnly" flag
		cookie := http.Cookie{
			Name:     "sessionID",
			Value:    "exampleSessionToken",
			Expires:  time.Now().Add(24 * time.Hour), // Expires in 24 hours
			Path:     "/",
			HttpOnly: false, // Vulnerable: Cookie accessible to client-side JavaScript
		}
		http.SetCookie(w, &cookie)

		fmt.Fprintf(w, "Login successful. Session cookie set.")
	} else {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
	}
}
