package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "exit" {
			break
		}
		_, err = conn.Write([]byte(message + "\n"))
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
}
