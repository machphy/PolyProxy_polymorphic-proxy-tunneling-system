package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting PolyProxy client (frpc)...")

	conn, err := net.Dial("tcp", "127.0.0.1:7000")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	message := "Hello from frpc client!"
	fmt.Println("Sending:", message)
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Fatalf("Failed to send message: %v", err)
	}

	// Read server response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading: %v", err)
	}
	fmt.Println("Server replied:", string(buffer[:n]))
}
