package main

import (
	"net"
	"fmt"
)

const (
	MAXBUFFER = 5000
)

func startListening(address string, port int) {
	bind := fmt.Sprintf("%s:%d", address, port)
	listener, err := net.Listen("tcp", bind)

	if err != nil {
		fatal("Unable to start server", err)
	}

	defer listener.Close()

	for {
		info("Waiting for connections on " + bind)
		connection, err := listener.Accept()

		if err != nil {
			warn("Failed to accept a connection")
			warn(err.Error())
			continue
		}

		go handleConnection(connection)

	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	buffer := make([]byte, MAXBUFFER)

	for {
		mlen, err := c.Read(buffer)

		if err != nil {
			return
		}

		if mlen < 3 {
			continue
		}

		if mlen > MAXBUFFER {
			msg := fmt.Sprintf("Message exceeded %d bytes", MAXBUFFER)
			warn(msg)
			c.Write([]byte(msg))
			return
		}

		resp := handleCommand(string(buffer[:mlen]))

		c.Write([]byte(resp))
	}

}