package main

import (
	"io"
	"log"
	"net"
	"os"
)

const addr = "127.0.0.1:8888"
const bufferSize = 256
const endLine = 10

var clients []net.Conn

/*
	The server waiting for some connection
*/
func main() {
	clients = make([]net.Conn, 0)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Ocurred a problem on listen " + addr)
		os.Exit(1)
	}

	for {
		conn, _ := listener.Accept()
		clients = append(clients, conn)

		// Go routine
		go handleConnection(conn)
	}
}

/*
	Buffer control to send messages to others clients.
*/
func handleConnection(conn net.Conn) {
	defer conn.Close()

	var data []byte
	buffer := make([]byte, bufferSize)

	for {
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				if err == io.EOF {
					break
				}
			}

			buffer = buffer[:n]
			data = append(data, buffer...)
			if data[len(data)-1] == endLine {
				break
			}
		}

		sendMessageToOtherClients(conn, data)

		data = make([]byte, 0)
	}
}

/*
	Sending messages to other clients
*/
func sendMessageToOtherClients(sender net.Conn, data []byte) {
	for i := 0; i < len(clients); i++ {
		if clients[i] != sender {
			clients[i].Write(data)
		}
	}
}
