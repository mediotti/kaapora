package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <server-ip> <server-port>\n", os.Args[0])
		os.Exit(1)
	}

	serverIP := os.Args[1]
	serverPort := os.Args[2]

	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", serverIP, serverPort))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	message := "Hello, rendezvous!"
	fmt.Printf("Sending message to %s:%s: %s\n", serverIP, serverPort, message)

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving:", err.Error())
		return
	}

	fmt.Printf("Received response: %s\n", string(buffer[:n]))
}

