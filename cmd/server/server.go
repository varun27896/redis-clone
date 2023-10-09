package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	portNumber := ":6379"

	listener, err := net.Listen("tcp", portNumber)
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Enabling connection on port:", portNumber)

	connection, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Listening on port:", portNumber)

	defer connection.Close() // close connection once finished

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err = connection.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("error reading from client: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("buf after", string(buf))

		// ignore request and send back a PONG
		connection.Write([]byte("-OK\r\n"))
	}
}
