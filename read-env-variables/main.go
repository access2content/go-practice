package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("GOPATH")
	fmt.Println(path)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "6969"
	}

	fmt.Println(port)
}
