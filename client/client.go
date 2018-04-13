package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

const addr = "127.0.0.1:8888"
const bufferSize = 256
const endLine = 10

var username string
var in *bufio.Reader

func main() {
	// By Default the first param always be the program name
	if len(os.Args) < 2 {
		fmt.Println("Use: client <username>")
		os.Exit(1)
	}

	in = bufio.NewReader(os.Stdin)

	// Get username
	username = os.Args[len(os.Args)-1]

	var conn net.Conn
	var err error

	// Trying to connect to the server
	// When it's can, it stops execution
	for {
		fmt.Printf("Trying to connect with %s...\n", addr)
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			fmt.Printf("Connected with %s!\n", addr)
			break
		}
	}

	defer conn.Close()

	// Using the go routine to receive other user's messages
	go receiveMessages(conn)

	// Handling other user's messages
	handleClientConn(conn)
}

/*
	If exists received message its printed thus:
	username -> message
*/
func handleClientConn(conn net.Conn) {
	for {
		buf, _, _ := in.ReadLine()
		if len(buf) > 0 {
			conn.Write(append([]byte(username+" -> "), append(buf, endLine)...))
		}
	}
}

/*
	Receive messages from other client's
*/
func receiveMessages(conn net.Conn) {
	var data []byte
	buffer := make([]byte, bufferSize)

	for {
		for {
			n, err := conn.Read(buffer)
			if err != nil && err == io.EOF {
				break
			}

			buffer = buffer[:n]
			data = append(data, buffer...)
			if data[len(data)-1] == endLine {
				break
			}
		}

		fmt.Printf("%s\n", data[:len(data)-1])
		data = make([]byte, 0)
	}
}
