package main

import (
	"log"
	"net"
	"strings"
)

const (
	maxTcpPacketSize = 65535
)

type Server struct {
	addr   string
	nl     net.Listener
	quitch chan struct{}
	writer *RespWriter
}

func NewServer(addr string) *Server {
	return &Server{
		addr:   addr,
		quitch: make(chan struct{}),
		writer: &RespWriter{},
	}
}

func (s *Server) Serve() error {
	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	log.Printf("Serving at %s", s.addr)

	defer l.Close()
	s.nl = l

	go s.acceptCon()

	<-s.quitch

	return nil
}

func (s *Server) acceptCon() {
	for {
		conn, err := s.nl.Accept()
		if err != nil {
			log.Println("accept error", err)
			continue
		}
		log.Printf("Connected to %s ", conn.RemoteAddr().String())

		go s.readCon(conn)
	}
}

func (s *Server) readCon(con net.Conn) {
	defer con.Close()

	buf := make([]byte, maxTcpPacketSize)
	n, err := con.Read(buf)
	if err != nil {
		log.Fatal("Error while reading req: ", err.Error())
	}

	log.Printf("received %d bytes: %s ", n, buf[:n])

	s.processCommand(buf, con)

}

func (s *Server) processCommand(msg []byte, con net.Conn) {
	cmd := string(msg)
	if strings.Contains(strings.ToLower(cmd), "ping") {
		s.writer.Write(con, "PONG")
	}

}
