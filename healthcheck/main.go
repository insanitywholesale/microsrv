package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	// use URL from env var but
	// if empty default to localhost
	URL := os.Getenv("URL")
	if URL == "" {
		URL = "http://localhost"
	}

	// perform GET request
	fmt.Println("getting:", URL)
	r, err := http.Get(URL)
	if err != nil {
		os.Exit(1)
	}

	// check status code
	fmt.Println("status code:", r.StatusCode)
	if r.StatusCode != 200 {
		os.Exit(1)
	}

	// print exit message formatted
	// same as os.Exit(1) does
	fmt.Println("exit status 0")
	os.Exit(0)
}
