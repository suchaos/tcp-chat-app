package main

import "demo/client"

func main() {
	//tcpServer := server.New(8080)
	//tcpServer.Start()

	tcpClient := client.New(8080)
	tcpClient.Start()
}
