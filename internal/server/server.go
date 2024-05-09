package server

import (
	"bufio"
	"fmt"
	"net"
)

type ServerConfig struct {
	IP   string
	Port int
}

func Talk(conn net.Conn) {
	defer conn.Close()
	_, err := conn.Write([]byte("Hello, brodyaga!\n"))
	if err != nil {
		fmt.Printf("Failed to write to client: %s\n", err)
	}

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read from client: %s\n", err)
	}

	fmt.Printf("Client %q send: %s\n", conn.RemoteAddr(), message)

	conn.Write([]byte("Сам дебил!\n"))
	if err != nil {
		fmt.Printf("Failed to write to client: %s\n", err)
	}
}

func Start(conf ServerConfig) {
	address := fmt.Sprintf("%s:%d", conf.IP, conf.Port)
	listener, err := net.Listen("tcp", address)
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
		go Talk(connection)
	}
}
