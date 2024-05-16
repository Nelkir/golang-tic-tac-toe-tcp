package main

import (
	"net"

	"tictac/internal/client"
	"tictac/internal/envs"
	"tictac/internal/server"
)

func main() {
	envs := envs.GetEnvs()

	var connection net.Conn
	switch envs.Mode {
	case "client":
		connection = client.Connect(client.ClientCofig{
			IP:   envs.ClientIP,
			Port: envs.ClientPort,
		})
	default:
		connection = server.Start(server.ServerConfig{
			IP:   envs.ServerIP,
			Port: envs.ServerPort,
		})
	}

	_ = connection

	return
}
