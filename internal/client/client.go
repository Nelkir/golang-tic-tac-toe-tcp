package client

import (
	"fmt"
	"net"
)

type ClientCofig struct {
	IP   string
	Port int
}

func Connect(config ClientCofig) net.Conn {
	address, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.IP, config.Port))
	if err != nil {
		fmt.Printf("Failed to resolve tcp address: %s\n", err)
		return nil
	}

	fmt.Printf("Connecting to %s", address.String())

	connection, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		fmt.Printf("Failed to dial %q", address.String())
		return nil
	}

	return connection
}
