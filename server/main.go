package main

import (
	"fmt"
	"net"
)

func main() {
	port := 1304 // Change this to the port you want to listen on
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4zero, Port: port})
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer conn.Close()
	fmt.Printf("Server is listening on port %d\n", port)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Received %d bytes from %s: %s\n", n, addr.String(), string(buffer[:n]))

		// You can add your processing logic here

		// Echo back the received data
		message := []byte("Hello, client!")
		_, err = conn.WriteToUDP(message, addr)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
