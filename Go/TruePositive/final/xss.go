package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "path/filepath"
)

func main() {
        http.HandleFunc("/readfile", func(w http.ResponseWriter, r *http.Request) {
                fileName := getSource(r)
                if fileName == "" {
                        http.Error(w, "File name not provided", http.StatusBadRequest)
                        return
                }
                filePath := filepath.Join("uploads", fileName)
                data, err := ioutil.ReadFile(filePath)
                if err != nil {
                        http.Error(w, "Failed to read file", http.StatusInternalServerError)
                        return
                }
                writeResponse(w, data) // Sink
        })

        http.ListenAndServe(":8080", nil)
}

func getSource(r *http.Request) string {
        return r.URL.Query().Get("file") // Source
}

func writeResponse(w http.ResponseWriter, data []byte) {
        fmt.Fprintf(w, "%s", data) // Vulnerable to XSS
}
