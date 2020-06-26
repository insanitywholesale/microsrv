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

	fmt.Println("getting:", URL)
	r, err := http.Get(URL)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("status code:", r.StatusCode)
	if r.StatusCode != 200 {
		os.Exit(1)
	}
	fmt.Println("exit status 0")
	os.Exit(0)
}
