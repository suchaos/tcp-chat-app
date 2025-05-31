package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TCPServer struct {
	port int
}

func New(port int) *TCPServer {
	return &TCPServer{port: port}
}

func (s *TCPServer) Start() {
	addr := fmt.Sprintf(":%d", s.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()

	log.Println("Server is running on :8080")

	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go processConnect(connect)
	}
}

func processConnect(connect net.Conn) {
	log.Printf("New connection from %s", connect.RemoteAddr())
	defer connect.Close()

	reader := bufio.NewReader(connect)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		log.Printf("Received message: %s", message)
		_, err = connect.Write([]byte(message + "\n"))
		if err != nil {
			log.Printf("Error writing message: %v", err)
			return
		}
	}
}
