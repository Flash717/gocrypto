package main

import (
	"bufio"
	"crypto/tls"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)

	for {
		msg, _ := r.ReadString('\n')

		println(msg)

		conn.Write([]byte("world\n"))
	}
}

func tlsServer() {
	cer, _ := tls.LoadX509KeyPair("server.crt", "server.key")

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, _ := tls.Listen("tcp", ":443", config)
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
