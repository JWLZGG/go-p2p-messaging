package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Start server listener
	go startServer()

	// Ask for peer address to connect to
	fmt.Print("Enter peer address (ip:port): ")
	peerAddress, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	peerAddress = peerAddress[:len(peerAddress)-1]package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// Start the server listener in a goroutine
	go startServer()

	// Allow some time for the server to start
	time.Sleep(1 * time.Second)

	// Prompt for peer address to connect to
	fmt.Print("Enter peer address (ip:port): ")
	reader := bufio.NewReader(os.Stdin)
	peerAddress, _ := reader.ReadString('\n')
	peerAddress = strings.TrimSpace(peerAddress) // Trim any newline or extra spaces

	// Debug: Print the peer address to verify formatting
	fmt.Printf("Peer address entered (trimmed): '%s'\n", peerAddress)

	// Check for valid format
	if !strings.Contains(peerAddress, ":") {
		fmt.Println("Invalid address format. Please use ip:port format, e.g., localhost:9001")
		return
	}

	// Attempt to connect to the peer
	peerXpeer(peerAddress)
}

func startServer() {
	// Listen on localhost:9000
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on localhost:9000") // Delayed message to avoid overlapping with input prompt

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("Accepted a connection from a peer!")
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed by peer.")
			break
		}
		fmt.Print("Peer: " + message)
	}
}

func peerXpeer(peerAddress string) {
	// Debug: Confirm the exact address we’re trying to connect to
	fmt.Printf("Attempting to connect to peer at: '%s'\n", peerAddress)

	// Connect to the specified peer address
	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Successfully connected to", peerAddress)
	for {
		fmt.Print("You: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(message))
	}
}


	// Connect to peer
	peerXpeer(peerAddress) // Ensure this matches the function name
}

func startServer() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:9000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Connection closed.")
			break
		}
		fmt.Print("Peer: " + message)
	}
}

func peerXpeer(peerAddress string) {
	conn, err := net.Dial("tcp", peerAddress)
	if err != nil {
		fmt.Println("Error connecting to peer:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", peerAddress)
	for {
		fmt.Print("You: ")
		message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(message))
	}
}
