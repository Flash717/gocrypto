package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

//Code: Key Derivation

func main() {
	salt := make([]byte, 8)
	rand.Read(salt)
	dk := pbkdf2.Key([]byte("SecretPassword"), salt, 1000, 32, sha256.New)
	fmt.Print(dk)
}
