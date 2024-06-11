package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
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

	fmt.Printf("Connecting to %s\n", address.String())

	connection, err := net.DialTCP("tcp", nil, address)
	if err != nil {
		fmt.Printf("Failed to dial %q\n", address.String())
		return nil
	}

	return connection
}

func Start(connection net.Conn) {
	writer, reader := bufio.NewWriter(connection), bufio.NewReader(connection)

	go func(reader *bufio.Reader) {
		for {
			b, err := reader.ReadByte()
			if err != nil {
				if err.Error() == "EOF" {
					fmt.Printf("Server closed connection\n")
					connection.Close()
					os.Exit(0)
				}
				fmt.Printf("Error occurred reading from server: %s\n", err)
				connection.Close()
				return
			}
			fmt.Print(string(b))
		}
	}(reader)

	stdin := bufio.NewReader(os.Stdin)
	for {
		message, err := stdin.ReadBytes('\n')
		if err != nil {
			fmt.Printf("Error reading user input: %s\n", err)
		}

		writer.Write(message)
		err = writer.Flush()
		if err != nil {
			fmt.Printf("Failed to write to a server: %s\n", err)
			connection.Close()
			return
		}
	}
}
