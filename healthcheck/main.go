package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	URL := os.Getenv("URL")
	if URL == "" {
		URL = "http://localhost"
	}

	_, err := http.Get(URL)
	fmt.Println("getting", URL)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("exit status 0")
	os.Exit(0)
}
