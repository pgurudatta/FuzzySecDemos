package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/readfile", func(w http.ResponseWriter, r *http.Request) {
		// Extract the file name from the URL query parameters
		fileName := r.URL.Query().Get("file")

		// Check if the file name is empty
		if fileName == "" {
			http.Error(w, "File name not provided", http.StatusBadRequest)
			return
		}

		// Construct the absolute path to the file
		filePath := filepath.Join("uploads", fileName)

		// Attempt to read the file
		data, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		// Write the file content to the response
		fmt.Fprintf(w, "%s", data)
	})

	http.ListenAndServe(":8080", nil)
}
