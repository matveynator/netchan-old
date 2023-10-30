# Welcome to the netchan Wiki!

<p align="right">
    <img align="right" property="og:image" src="https://repository-images.githubusercontent.com/710838463/86ad7361-2608-4a70-9197-e66883eb9914" width="30%">
</p>


## Overview
`netchan` stands as a robust library for the Go programming language, offering convenient and secure abstractions for network channel interactions. Inspired by Rob Pike’s initial concept, it aims to deliver an interface that resonates with the simplicity and familiarity of Go’s native channels.

## Key Features
- **Ease of Use**: The library’s interface is intuitively designed, mirroring the standard channel operations in Go, ensuring a seamless integration for Go developers.
- **Secure by default**: Emphasizing security, `netchan` employs cutting-edge encryption techniques, alongside reliable authentication and authorization practices, ensuring the integrity and "security by default" of your network communications.
- **Scalability**: Crafted with distributed systems in mind, `netchan` excels in environments demanding high throughput and scalability.
- **High Performance**: Performance is paramount; hence, `netchan` is meticulously optimized to ensure low overhead and swift data transmissions.
- **Network Adherence to CSP Principles**: Complete alignment with the Communicating Sequential Processes (CSP) model as articulated by Tony Hoare, extended to function at the network layer.
- **Principles of Pure Go Programming**: A seamless adherence to the [Go proverbs](https://go-proverbs.github.io), harmonizing the coding practices with the most esteemed and recognized conventions of Go programming principles.


## Getting Started
To embark on your journey with `netchan`, install the library using `go get`:
```
go get -u github.com/matveynator/netchan
```
## Usage Example:

### cat examples/client-server/server.go  
```
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
```

### cat examples/client-server/client.go 
```
package main

import (
	"github.com/matveynator/netchan"
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
```

## Documentation
- [v1.0 Plan](wiki/v1-plan.md)
- Usage Examples
- API References
- Secure by default

## Community and Support
Should you have inquiries or suggestions, feel free to open an [issue](https://github.com/matveynator/netchan/issues) in our GitHub repository.

## License
`netchan` is distributed under the BSD-style License. For detailed information, please refer to the [LICENSE](https://github.com/matveynator/netchan/blob/master/LICENSE) file.

## Recovering original netchan 2013 code by Rob Pike:
```
git clone https://go.googlesource.com/exp
cd exp
git reset --hard a05e19747a0323e4140d9d054102e7d7f9c0812a
cd old/netchan
```



