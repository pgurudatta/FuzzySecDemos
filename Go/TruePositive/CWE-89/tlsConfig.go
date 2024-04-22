package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://example.com"

	// Disable TLS certificate verification (INSECURE - DO NOT USE IN PRODUCTION)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// Send HTTP GET request to the URL
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response Status: %s\n", resp.Status)
}
