package main

import (
	"fmt"
	"net"
)

func main() {
	// Écouter sur le port 8080 en UDP
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Erreur lors de l'écoute:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Serveur UDP en écoute sur le port 8080")

	buffer := make([]byte, 1024)

	for {
		// Lire les données du client
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Erreur lors de la lecture:", err)
			continue
		}

		// Afficher ce qui a été reçu
		received := string(buffer[:n])
		fmt.Printf("Reçu du client %v: %s\n", clientAddr, received)

		// Envoyer la réponse au client
		response := "HTTP/1.1 200 Bytes1"
		_, err = conn.WriteToUDP([]byte(response), clientAddr)
		if err != nil {
			fmt.Println("Erreur lors de l'envoi de la réponse:", err)
			continue
		}

		fmt.Printf("Réponse envoyée au client %v: %s\n", clientAddr, response)
	}
}
s