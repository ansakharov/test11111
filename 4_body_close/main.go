package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com/")
	if err != nil {
		fmt.Println(err)

		return
	}

	// response body must be closed (bodyclose)
	defer resp.Body.Close() // OK

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body), err)
}
