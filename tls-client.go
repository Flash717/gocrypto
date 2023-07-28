package main

import (
	"crypto/tls"
	"log"
)

func tlsClient() {
	log.SetFlags(log.Lshortfile)
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, _ := tls.Dial("tcp", "127.0.0.1:443", conf)

	defer conn.Close()

	n, _ := conn.Write([]byte("hello\n"))

	buf := make([]byte, 100)
	n, _ = conn.Read(buf)

	println(string(buf[:n]))
}
