package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var userInput string

	fmt.Print("Enter a file name: ")
	fmt.Scanln(&userInput)

	// Vulnerable code: directly using user input in a system command
	cmd := exec.Command("cat", userInput)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	fmt.Printf("Command output:\n%s\n", output)
}
