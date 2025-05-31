package main

import (
	"demo/client"
	"demo/server"
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "server", "运行模式: server or client")
	port := flag.Int("port", 8080, "监听或者连接的端口号")

	flag.Parse()

	switch *mode {
	case "server":
		tcpServer := server.New(*port)
		tcpServer.Start()
	case "client":
		tcpClient := client.New(*port)
		tcpClient.Start()
	default:
		fmt.Printf("未知模式 %s, 请指定正确的运行模式", *mode)
		flag.Usage()
	}
}
