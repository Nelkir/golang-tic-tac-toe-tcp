package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"tictac/internal/client"
	"tictac/internal/envs"
	"tictac/internal/server"
	"tictac/internal/tictac"
)

func main() {
	envs := envs.GetEnvs()
	connection_chan := make(chan net.Conn, 2)

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
		go server.HandleConnections(server.ServerConfig{
			IP:   envs.ServerIP,
			Port: envs.ServerPort,
		}, connection_chan)
		for {
			switch len(connection_chan) {
			case 2:
				go HandlePlayers(<-connection_chan, <-connection_chan)
			case 1:
				fmt.Printf("Waiting for second player\n")
				time.Sleep(time.Second)
			default:
				time.Sleep(time.Second)
			}
		}
	}

	return
}

func HandlePlayers(player1 net.Conn, player2 net.Conn) {
	defer player1.Close()
	defer player2.Close()
	tictac.Start(
		tictac.NewRemotePlayer('O', player1),
		tictac.NewRemotePlayer('X', player2),
	)
}
