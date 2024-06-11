package server

import (
	"fmt"
	"net"
)

type ServerConfig struct {
	IP   string
	Port int
}

func HandleConnections(conf ServerConfig, connections_chan chan net.Conn) {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", conf.IP, conf.Port))
	if err != nil {
		fmt.Printf("Failed to resolve tcp address: %s\n", err)
		return
	}

	fmt.Printf("Starting on %q\n", address.String())

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		fmt.Printf("Server failed to listen on %q: %s\n", address, err)
		return
	}

	fmt.Printf("Server listening on %q\n", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Client failed to connect: %s\n", err)
			continue
		}
		fmt.Printf("Client connected %q\n", connection.RemoteAddr())

		connections_chan <- connection
	}
}
