package main

import (
	"fmt"
	"net"
	"os"
	"web-server/lib"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	listener, error := net.Listen(CONN_TYPE, CONN_HOST + ":" + CONN_PORT)
	if error != nil {
		fmt.Println("Unable to open listener: ", error.Error())
		os.Exit(1)
	}

	defer listener.Close()
	fmt.Println("Listening on: " + CONN_HOST + ":" + CONN_PORT)

	for {
		incomingConnection, error := listener.Accept()
		if error != nil {
			fmt.Println("Error accepting request: " , error.Error())
			os.Exit(1)
		}

		go handleRequest(incomingConnection)
	}
}

func handleRequest(connection net.Conn) {
	buffer := make([]byte, 1024)
	_, error := connection.Read(buffer)

	if error != nil {
		panic("Error reading from incoming connection: " + error.Error())
	}

	httpRequest := tokeniser.TokeniseRequest(string(buffer))

	writeToConsoles(httpRequest)
	writeAndClose(connection)
}

func writeToConsoles(message tokeniser.HttpRequest) {
	fmt.Printf("[req] %s %s | Headers: %d \r\n", message.Method, message.Path, len(message.Headers))
}

func writeAndClose(conn net.Conn) {
	conn.Write([]byte("Hello world!"))
	conn.Close()
}