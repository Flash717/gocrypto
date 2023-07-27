package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("Secret password\n"))
	fmt.Printf("%x", h.Sum(nil))
}
