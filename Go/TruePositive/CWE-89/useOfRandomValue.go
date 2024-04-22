package main

import (
    "fmt"
    "math/rand"
)

func main() {
    // Vulnerable: No seeding
    sessionID := rand.Intn(100000) // Predictable sequence
    fmt.Println("Session ID:", sessionID)
}
