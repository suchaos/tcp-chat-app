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

	go readMsg(conn)

	writeMsg(err, conn)
}

func writeMsg(err error, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "exit" {
			break
		}
		_, err = conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Fatalf("Error writing message: %v", err)
			return
		}
	}
}

func readMsg(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading message: %v", err)
			return
		}
		if response == "\n" {
			continue
		}
		log.Printf("Received message: %s", response)
	}
}
