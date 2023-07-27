package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	key := []byte("YourPasswordMustBe32BitforAES256")

	var plainText = "Very sensitive secret data"

	encrypted := encrypt(key, plainText)

	//Print key and cypher
	fmt.Printf("CIPHER KEY: %s \n", string(key))
	fmt.Printf("ENCRYPTED: %s \n", encrypted)

	//Decrypt text
	decrypted := decrypt(key, encrypted)

	//Print re-decrypted text
	fmt.Printf("DECRYPTED: %s \n", decrypted)
	fmt.Printf("ORIGINAL: %s \n", plainText)
}

func encrypt(key []byte, message string) (encoded string) {
	plainText := []byte(message)
	block, _ := aes.NewCipher(key)
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv stands for initialization value or variable, it is the cyphertext up to the blocksize
	iv := cipherText[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)

	//Encrypt data
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return base64.RawStdEncoding.EncodeToString(cipherText)
}

func decrypt(key []byte, message string) (encoded string) {
	cipherText, _ := base64.RawStdEncoding.DecodeString(message)

	block, _ := aes.NewCipher(key)

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText)
}
