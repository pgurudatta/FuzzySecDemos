/*
Sample code for vulnerable type: Inadequate Encryption Strength
CWE : CWE-326
Description : Inadequate Encryption Strength
*/
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	//Generate Private Key
	pvk, err := rsa.GenerateKey(rand.Reader, 1024) //source and sink
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pvk)
}
