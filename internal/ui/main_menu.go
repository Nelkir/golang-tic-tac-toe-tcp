package ui

import (
	"fmt"
	"net"
	"os"

	"github.com/rivo/tview"

	"tictac/internal/client"
	"tictac/internal/envs"
	"tictac/internal/server"
	"tictac/internal/tictac"
)

var Applicaton *tview.Application

func init() {
	Applicaton = tview.NewApplication()
}

func MainMenu(envs envs.Envs) {
	var connection net.Conn

	MainMenuList := tview.NewList()
	MainMenuList.AddItem("Server", "Host Tic Tac Toe game server", 's', func() {
		connection = server.Start(server.ServerConfig{
			IP:   envs.ServerIP,
			Port: envs.ServerPort,
		})
		tictac.Start(tictac.NewLocalPlayer('O'), tictac.NewRemotePlayer('X', connection))
	})
	MainMenuList.AddItem("Client", "Connect to Tic Tac Toe game server", 'c', func() {
		connection = client.Connect(client.ClientCofig{
			IP:   envs.ClientIP,
			Port: envs.ClientPort,
		})
	})
	MainMenuList.AddItem("Quit", "Quit Tic Tac Toe", 'q', func() {
		Applicaton.Stop()
	})

	Applicaton.SetRoot(MainMenuList, true).SetFocus(MainMenuList)
	if err := Applicaton.Run(); err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}
}
