package main

import (
	"fmt"
	"net/http"
)

// contentType will return the value of Content-Type header returned by making a
// HTTP GET requst to a url

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close() // Make sure we close the body.

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" { // Return error if Content-Type header not found.
		return "", fmt.Errorf("Can't find Content-Type header")
	}

	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
