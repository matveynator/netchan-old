package main

import (
	"github.com/matveynator/netchan-old"
	"log"
	"net"
)

func main() {
	// Connect to the server
	connection, err := net.Dial("tcp", "127.0.0.1:12345") // Replace "127.0.0.1" with the server's IP or hostname
	if err != nil {
		log.Println("Error connecting to server:", err)
		return
	}

	log.Println("Connected to server, setting up importer...")

	// Create a new Importer on the connection
	importer := netchan.NewImporter(connection)

	// Create channels for sending and receiving messages
	pingChannel := make(chan string)
	pongChannel := make(chan string)

	// Import channels for communication with the server
	err = importer.Import("pingChannel", pingChannel, netchan.Recv, 1)
	if err != nil {
		log.Println("Error importing ping channel:", err)
		return
	}
	err = importer.Import("pongChannel", pongChannel, netchan.Send, 1)
	if err != nil {
		log.Println("Error importing pong channel:", err)
		return
	}

	log.Println("Channels imported, starting pong handler...")

	// Start the pong handler
	pongHandler(pongChannel, pingChannel)
}

// pongHandler waits for ping messages and sends pong responses
func pongHandler(pongChannel chan string, pingChannel chan string) {
	for {
		msg := <-pingChannel          // Wait for a ping message
		log.Println("Received:", msg) // Log the received ping message
		pongChannel <- "pong"         // Send a pong response
	}
}
