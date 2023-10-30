package main

import (
	"github.com/matveynator/netchan"
	"log"
	"net"
	"time"
)

func main() {
	// Set up a TCP server to accept connections from clients on port 12345
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Println("Error setting up listener:", err)
		return
	}

	log.Println("Server is waiting for a connection...")

	// Accept a connection from a client
	connection, err := listener.Accept()
	if err != nil {
		log.Println("Error accepting connection:", err)
		return
	}

	log.Println("Connection established with client.")

	// Create a new Exporter and serve it on the connection
	exporter := netchan.NewExporter()
	go exporter.ServeConn(connection) // Start serving in a separate goroutine

	// Create channels for sending and receiving messages
	pingChannel := make(chan string)
	pongChannel := make(chan string)

	// Export channels for communication with the client
	err = exporter.Export("pingChannel", pingChannel, netchan.Send)
	if err != nil {
		log.Println("Error exporting ping channel:", err)
		return
	}
	err = exporter.Export("pongChannel", pongChannel, netchan.Recv)
	if err != nil {
		log.Println("Error exporting pong channel:", err)
		return
	}

	log.Println("Channels exported, starting ping handler...")

	// Start the ping handler in a separate goroutine
	go pingHandler(pingChannel, pongChannel)

	// Block indefinitely to keep the server running
	select {}
}

// pingHandler sends ping messages and waits for pong responses
func pingHandler(pingChannel chan string, pongChannel chan string) {
	for {
		pingChannel <- "ping"         // Send a ping message
		msg := <-pongChannel          // Wait for a pong response
		log.Println("Received:", msg) // Log the received pong message
		time.Sleep(1 * time.Second)   // Wait for a second before next ping
	}
}
