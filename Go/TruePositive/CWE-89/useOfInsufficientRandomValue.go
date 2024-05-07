/*
Sample code for vulnerable type: Use of Insufficiently Random Values
CWE : CWE-330
Description : Use of Insufficiently Random Values
*/
package main

import (
	"math/rand"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func generatePassword() string {
	s := make([]rune, 20)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]  //source and sink
	}
	return string(s)
}
func main() {
        password := generatePassword()
        println(password)
}
