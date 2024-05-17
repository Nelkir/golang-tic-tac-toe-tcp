package server

import (
	"fmt"
	"net"
)

type ServerConfig struct {
	IP   string
	Port int
}

func Start(conf ServerConfig) net.Conn {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", conf.IP, conf.Port))
	if err != nil {
		fmt.Printf("Failed to resolve tcp address: %s\n", err)
		return nil
	}

	fmt.Printf("Starting on %q\n", address.String())

	listener, err := net.ListenTCP("tcp", address)
	if err != nil {
		fmt.Printf("Server failed to listen on %q: %s\n", address, err)
		return nil
	}

	fmt.Printf("Server listening on %q\n", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Printf("Client failed to connect: %s\n", err)
			continue
		}
		fmt.Printf("Client connected %q\n", connection.RemoteAddr())

		return connection
	}
}
