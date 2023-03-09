package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	port := "6379"
	host := "0.0.0.0"

	server := NewServer(fmt.Sprintf("%s:%s", host, port))

	err := server.Serve()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
