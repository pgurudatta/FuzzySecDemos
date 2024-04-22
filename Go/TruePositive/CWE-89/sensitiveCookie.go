package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/setCookie", setCookieHandler)
	http.HandleFunc("/readCookie", readCookieHandler)

	// Start HTTPS server on port 8443
	err := http.ListenAndServeTLS(":8443", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatalf("Failed to start HTTPS server: %v", err)
	}
}

func setCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Set a sensitive cookie without the 'Secure' attribute
	cookie := &http.Cookie{
		Name:    "sessionID",
		Value:   "exampleSessionToken",
		Expires: time.Now().Add(24 * time.Hour), // Expires in 24 hours
		HttpOnly: true, // Protect cookie from JavaScript access
		// Missing 'Secure' attribute
	}
	http.SetCookie(w, cookie)

	fmt.Fprintf(w, "Cookie set successfully")
}

func readCookieHandler(w http.ResponseWriter, r *http.Request) {
	// Read the cookie value from the request
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		http.Error(w, "Failed to read cookie", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Session ID: %s", cookie.Value)
}
