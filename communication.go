package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler1(w http.ResponseWriter, req *http.Request) {
	// restrict to only needed commands
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
}

func handler2(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Security-Policy", `default-src 'none'; 
	front-src 'none'; img-src 'self'; object-src 'none'; script-src 'self'`)
	w.Header().Add("X-XSS-Protection", "1; mode=block")
}

func contentSecurePolicy() {
	fmt.Println("Set Content-Security-Policy if possible (learn more about it)")
}

func communication() {
	// secure communication implementation
	http.HandleFunc("/", handler2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
