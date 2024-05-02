package main

import (
	"fmt"
	"math/rand"
	"time/timestamp"
)

func generateSessionID() string {
  rand.Seed(timestamp.Now().UnixNano())
  return fmt.Sprintf("session_%d", rand.Int63())
}

func main() {
  sessionID := generateSessionID()
  fmt.Println("Session ID:", sessionID)
}
