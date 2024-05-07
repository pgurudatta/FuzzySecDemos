/*
Sample code for vulnerable type: Open Redirect
CWE : CWE-601
Description : URL Redirection to Untrusted Site ('Open Redirect')
*/
package main

import (
	"fmt"
	"net/http"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url") //  source (User-controlled input)

	// Vulnerable: User input directly used to construct redirect URL
	http.Redirect(w, r, url, http.StatusFound) //sink
}

func main() {
	http.HandleFunc("/", handleRequest)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
