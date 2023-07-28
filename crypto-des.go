package main

import (
	"crypto/des"
	"encoding/hex"
	"fmt"
)

func cryptoDES() {
	key := "mysecretPasswordKeySiz24"
	plainText := "Secret12"

	cipherText := EncryptTripleDES([]byte(key), plainText)
	decryptedText := DecryptTripleDES([]byte(key), cipherText)

	fmt.Printf("Plain Text: %s\n", plainText)
	fmt.Printf("Ciper Text: %s \n", cipherText)
	fmt.Printf("Decrypted Text: %s\n", decryptedText)
}

func EncryptTripleDES(key []byte, plaintext string) string {
	cipher, _ := des.NewTripleDESCipher(key)

	out := make([]byte, len(plaintext))
	cipher.Encrypt(out, []byte(plaintext))
	return hex.EncodeToString(out)
}

func DecryptTripleDES(key []byte, ct string) string {
	ciphertext, _ := hex.DecodeString(ct)
	cipher, _ := des.NewTripleDESCipher([]byte(key))

	plaintext := make([]byte, len(ciphertext))
	cipher.Decrypt(plaintext, ciphertext)
	output := string(plaintext[:])
	return output
}
