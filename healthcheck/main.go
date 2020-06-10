package main

import (
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://localhost")
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
