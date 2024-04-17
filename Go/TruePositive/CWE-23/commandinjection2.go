package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	var userInput string

	fmt.Print("Enter a directory path: ")
	fmt.Scanln(&userInput)

	// Vulnerable code: using template literals with user input
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("ls -l %s", userInput))
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	fmt.Printf("Command output:\n%s\n", output)
}
