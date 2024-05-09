package main

import (
	"tictac/internal/args"
	"tictac/internal/server"
)

func main() {
	args := args.GetArgs()
	server.Start(server.ServerConfig{
		IP:   args.ServerIP,
		Port: args.ServerPort,
	})
	return
}
