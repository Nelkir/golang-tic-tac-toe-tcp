package tictac

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

type Player struct {
	Rune       rune
	readWriter *bufio.ReadWriter
}

type player interface {
	Greeting()
	Move()
	write(message string)
	read() string
}

type Players []Player

func (players Players) Message(message string) {
	for _, player := range players {
		player.write(message)
	}
}

func (player *Player) Greeting() {
	player.write(fmt.Sprintf("Hello, Player! Your character is %q\n", player.Rune))
}

func (player *Player) write(message string) (int, error) {
	return player.readWriter.WriteString(message)
}

func (player *Player) read() (string, error) {
	return player.readWriter.ReadString('\n')
}

func NewRemotePlayer(r rune, connection net.Conn) Player {
	readWriter := bufio.NewReadWriter(bufio.NewReader(connection),
		bufio.NewWriter(connection))

	return Player{
		Rune:       r,
		readWriter: readWriter,
	}
}

func NewLocalPlayer(r rune) Player {
	readWriter := bufio.NewReadWriter(bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout))

	return Player{
		Rune:       r,
		readWriter: readWriter,
	}
}

func (player *Player) Move() (int, error) {
	_, err := player.write(fmt.Sprintf("Your move %q: ", player.Rune))
	if err != nil {
		fmt.Printf("Failed to write to a Player %q: %s\n", player.Rune, err)
		return -1, err
	}

	message, err := player.read()
	if err != nil {
		fmt.Printf("Failed to read from a Player %q: %s\n", player.Rune, err)
	}

	messageInt, err := strconv.Atoi(message)
	if err != nil {
		fmt.Printf("Failed to parse Player %q move: %s\n", player.Rune, err)
	}

	return messageInt, nil
}
