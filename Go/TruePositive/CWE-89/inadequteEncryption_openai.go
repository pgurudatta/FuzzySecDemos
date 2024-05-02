package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func encryptData(key []byte, plaintext []byte) ([]byte, error) {
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }

  gcm, err := cipher.NewGCM(block)
  if err != nil {
    return nil, err
  }

  nonce := make([]byte, gcm.NonceSize())
  ciphertext, err := gcm.Seal(nonce, nil, plaintext, nil)
  if err != nil {
    return nil, err
  }

  return append(nonce, ciphertext...), nil
}

func main() {
  key := []byte("secret") // This is a short and insecure key!
  data := []byte("This is some confidential data")

  ciphertext, err := encryptData(key, data)
  if err != nil {
    // Handle error
  }
  // ... use ciphertext ...
}
