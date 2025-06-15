package main

import (
	"fmt"
	"net"
)

func main() {
	// Create UDP address to listen on
	udpAddr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create UDP connection
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on port 8080")

	buffer := make([]byte, 1024)

	for {
		// Read from UDP connection
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		fmt.Printf("Received from %s: %s\n", addr.String(), string(buffer[:n]))

		// Prepare response message
		response := fmt.Sprintf("HTTP/1.1 200 Bytes1%sBytes0\r\n\r\n", string(buffer[:n]))

		// Send response back to client
		_, err = conn.WriteToUDP([]byte(response), addr)
		if err != nil {
			fmt.Println("Error writing:", err)
			continue
		}
	}
}