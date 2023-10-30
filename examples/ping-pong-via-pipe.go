package main

import (
	"fmt"
	"github.com/matveynator/netchan" // Import the netchan package
	"net"
	"time"
)

func main() {
	// Create a pair of connected network endpoints using net.Pipe()
	conn1, conn2 := net.Pipe()

	// Create a new Exporter and serve it on one of the connections
	exporter := netchan.NewExporter()
	go exporter.ServeConn(conn1) // Start serving in a separate goroutine

	// Create a new Importer on the other connection
	importer := netchan.NewImporter(conn2)

	// Create channels for sending and receiving messages
	pingChan := make(chan string)
	pongChan := make(chan string)

	// Export channels: this makes them available for other processes to connect to
	// Exporting the ping channel for sending messages
	err := exporter.Export("pingChan", pingChan, netchan.Send)
	if err != nil {
		fmt.Println("Export error:", err)
		return
	}
	// Exporting the pong channel for receiving messages
	err = exporter.Export("pongChan", pongChan, netchan.Recv)
	if err != nil {
		fmt.Println("Export error:", err)
		return
	}

	// Import channels: this connects to exported channels from other processes
	// Importing the pong channel for sending messages
	err = importer.Import("pongChan", pongChan, netchan.Send, 1)
	if err != nil {
		fmt.Println("Import error:", err)
		return
	}
	// Importing the ping channel for receiving messages
	err = importer.Import("pingChan", pingChan, netchan.Recv, 1)
	if err != nil {
		fmt.Println("Import error:", err)
		return
	}

	// Start the ping and pong handlers in separate goroutines
	go pingHandler(pingChan, pongChan)
	go pongHandler(pongChan, pingChan)

	// Send the initial ping message to kickstart the ping-pong exchange
	go func() {
		pingChan <- "ping"
	}()

	// Allow the ping-pong exchange to run for a while
	time.Sleep(10 * time.Second)
}

// pingHandler sends ping messages and waits for pong responses
func pingHandler(pingChan chan string, pongChan chan string) {
	for {
		msg := <-pongChan           // Wait for a pong response
		fmt.Println(msg)            // Print the received pong message
		time.Sleep(1 * time.Second) // Wait for a second
		pingChan <- "ping"          // Send a new ping message
	}
}

// pongHandler waits for ping messages and sends pong responses
func pongHandler(pongChan chan string, pingChan chan string) {
	for {
		msg := <-pingChan  // Wait for a ping message
		fmt.Println(msg)   // Print the received ping message
		pongChan <- "pong" // Send a pong response
	}
}
