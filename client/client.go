package client

import (
	"bufio"
	"context"
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

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("Client is running")
	// 启动读协程
	go func() {
		readMsg(ctx, conn)
		// 退出时取消上下文
		cancel()
	}()
	// 启动写协程
	go func() {
		writeMsg(cancel, conn)
	}()

	//  等待上下文取消
	<-ctx.Done()
	fmt.Println("Client is exiting")
}

func writeMsg(cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		if message == "exit" {
			fmt.Println("Exiting...")
			cancel()
			return
		}
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			log.Printf("Error writing message: %v", err)
			cancel()
			return
		}
	}
}

func readMsg(ctx context.Context, conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		select {
		case <-ctx.Done():
			log.Printf("Connection closed")
			return
		default:
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
}
