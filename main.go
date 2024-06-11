package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"tictac/internal/client"
	"tictac/internal/envs"
	"tictac/internal/server"
	"tictac/internal/tictac"
)

func main() {
	envs := envs.GetEnvs()

	var connection net.Conn
	fmt.Printf("Starting as %s\n", envs.Mode)
	switch strings.ToLower(envs.Mode) {
	case "client":
		connection = client.Connect(client.ClientCofig{
			IP:   envs.ClientIP,
			Port: envs.ClientPort,
		})
		if connection == nil {
			fmt.Printf("Failed to create connection\n")
			os.Exit(-1)
		}
		client.Start(connection)
	default:
		connection = server.Start(server.ServerConfig{
			IP:   envs.ServerIP,
			Port: envs.ServerPort,
		})
		if connection == nil {
			fmt.Printf("Failed to create connection\n")
			os.Exit(-1)
		}
		tictac.Start(tictac.NewLocalPlayer('O'), tictac.NewRemotePlayer('X', connection))
	}

	return
}
