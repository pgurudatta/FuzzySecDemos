package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := "SensitiveData123"

	// Using MD5 hash algorithm (INSECURE - DO NOT USE MD5 FOR CRYPTOGRAPHIC PURPOSES)
	hash := md5.Sum([]byte(data))
	hashString := fmt.Sprintf("%x", hash)

	fmt.Printf("MD5 Hash: %s\n", hashString)
}
