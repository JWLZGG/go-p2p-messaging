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
	peerAddress = peerAddress[:len(peerAddress)-1]

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
