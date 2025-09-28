package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/BurntSushi/toml"
)

type ProxyConfig struct {
	Name       string `toml:"name"`
	Type       string `toml:"type"`
	LocalPort  int    `toml:"localPort"`
	RemotePort int    `toml:"remotePort"`
}

type FrpcConfig struct {
	Proxies []ProxyConfig `toml:"proxies"`
}

func main() {
	fmt.Println("Starting PolyProxy client (frpc)...")

	// Load configuration
	configPath := "../../configs/example_frpc.toml"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file not found: %s", configPath)
	}

	var cfg FrpcConfig
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		log.Fatalf("Error decoding config: %v", err)
	}

	// Connect to each proxy defined in config
	for _, proxy := range cfg.Proxies {
		address := fmt.Sprintf("127.0.0.1:%d", proxy.RemotePort)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Printf("[%s] Failed to connect: %v", proxy.Name, err)
			continue
		}
		defer conn.Close()

		message := fmt.Sprintf("Hello from %s!", proxy.Name)
		fmt.Printf("[%s] Sending: %s\n", proxy.Name, message)
		conn.Write([]byte(message))

		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Printf("[%s] Error reading: %v", proxy.Name, err)
			continue
		}
		fmt.Printf("[%s] Server replied: %s\n", proxy.Name, string(buffer[:n]))
	}
}
