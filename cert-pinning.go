package main

import (
	"fmt"

	"github.com/rustler47/SecureClient"
)

func certPinning() {
	proxy := "http://localhost:5555"
	url := []string{"google.com"}
	BadPinDetected := func(proxy string) {
		fmt.Println("WARNING! Failed SSL pinning - Invalid cert detected\n", "Proxy:", proxy)
	}

	pinner, err := SecureClient.New(url, true, BadPinDetected)
	if err != nil {
		return
	}

	client, err := pinner.NewClient(proxy)
	if err != nil {
		return
	}

	client.Get("https://google.com")
}
