package main

import (
        "fmt"
        "net/http"
        "path/filepath"
        "io/ioutil"
)

func main() {
        http.HandleFunc("/readfile", func(w http.ResponseWriter, r *http.Request) {
                fileName := getSource(r)
                if fileName == "" {
                        http.Error(w, "File name not provided", http.StatusBadRequest)
                        return
                }
                filePath := filepath.Join("uploads", fileName)
                data, err := readFile(filePath)
                if err != nil {
                        http.Error(w, "Failed to read file", http.StatusInternalServerError)
                        return
                }
                fmt.Fprintf(w, "%s", data)
        })

        http.ListenAndServe(":8080", nil)
}

func getSource(r *http.Request) string {
        return r.URL.Query().Get("file")
}

func readFile(filePath string) ([]byte, error) {
        data, err := ioutil.ReadFile(filePath)
        if err != nil {
                return nil, err
        }
        return data, nil
}
