package main

import (
	"net/http"
	"os/exec"
	"strings"
)

func main() {
	http.HandleFunc("/runcmd", func(w http.ResponseWriter, r *http.Request) {
		cmdName := getSource(r)
		if cmdName == "" {
			http.Error(w, "Command name not provided", http.StatusBadRequest)
			return
		}
		err := runCommand(cmdName)
		if err != nil {
			http.Error(w, "Failed to execute command", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("Command executed successfully"))
	})

	http.ListenAndServe(":8080", nil)
}

func getSource(r *http.Request) string {
	return r.URL.Query().Get("cmd")
}

func runCommand(cmdName string) error {
	// Split the command name into command and arguments
	parts := strings.Fields(cmdName)
	cmd := exec.Command(parts[0], parts[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
