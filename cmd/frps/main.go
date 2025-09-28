package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := "7000"
	fmt.Println("Starting PolyProxy server (frps)...")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer ln.Close()
	fmt.Printf("Server is listening on port %s...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("New client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from client: %v", err)
		return
	}

	message := string(buffer[:n])
	fmt.Println("Received from client:", message)

	// Send response back
	conn.Write([]byte("Server received: " + message))
}
