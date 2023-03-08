package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	maxTcpPacketSize = 65535
)

func main() {

	port := "6379"
	host := "0.0.0.0"

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	log.Printf("Server listing on %s:%s", host, port)
	for {
		conn, err := l.Accept()
		log.Printf("Connected to %s ", conn.RemoteAddr().String())
		if err != nil {
			log.Fatal("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		buffer := make([]byte, maxTcpPacketSize)
		_, err = conn.Read(buffer)
		if err != nil {
			log.Fatal("Error while reading req: ", err.Error())
			os.Exit(1)
		}
	}
}
