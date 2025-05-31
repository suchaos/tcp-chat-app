# TCP Chat Application

This is a simple TCP-based chat application implemented in Go. It consists of a server and a client component.

## Project Structure
- `server/` - Server implementation
- `client/` - Client implementation
- `main.go` - Entry point for both server and client

## Current Implementation

### Server
Currently, the server is a simple echo server that:
- Listens for incoming TCP connections
- Receives messages from clients
- Sends back an exact copy of received messages (with "echo: " prefix)

### Client
The client:
- Connects to the server
- Sends messages entered by the user
- Displays received echo messages

## Usage

### Start the Server
```bash
go run main.go -mode=server -port=8080
```

### Start a Client
```bash
go run main.go -mode=client -port=8080