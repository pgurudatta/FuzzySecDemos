package main

import (
    "net/http"
    "os/exec"
)

func handler(w http.ResponseWriter, req *http.Request) {
    // Get the value of the 'cmd' parameter from the URL query
    cmdName := req.URL.Query().Get("cmd")
    
    // Check if cmdName is empty
    if cmdName == "" {
        http.Error(w, "Command not provided", http.StatusBadRequest)
        return
    }
    
    // Execute the command
    cmd := exec.Command(cmdName)
    err := cmd.Run()
    if err != nil {
        http.Error(w, "Failed to execute command", http.StatusInternalServerError)
        return
    }
    
    // Send success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Command executed successfully"))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
