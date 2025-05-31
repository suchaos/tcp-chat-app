package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type TCPClient struct {
	port int
}

func New(port int) *TCPClient {
	return &TCPClient{
		port: port,
	}
}

func (c *TCPClient) Start() {
	addr := fmt.Sprintf(":%d", c.port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer conn.Close()

	fmt.Println("Client is running")

	message := "Hello, Server!\n"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading message: %v", err)
	}
	fmt.Println("Received message:", response)
}
