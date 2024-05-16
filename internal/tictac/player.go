package tictac

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Rune       rune
	readWriter *bufio.ReadWriter
	Connected  bool
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
	player.write(fmt.Sprintf("Your character is %q\n", player.Rune))
}

func (player *Player) write(message string) (int, error) {
	size, err := player.readWriter.WriteString(message)
	if err != nil {
		player.Connected = false
		return 0, err
	}
	return size, player.readWriter.Flush()
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
		Connected:  true,
	}
}

func NewLocalPlayer(r rune) Player {
	readWriter := bufio.NewReadWriter(bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout))

	return Player{
		Rune:       r,
		readWriter: readWriter,
		Connected:  true,
	}
}

func (player *Player) Move() (int, error) {
	_, err := player.write(fmt.Sprintf("Your move %q: ", player.Rune))
	if err != nil {
		fmt.Printf("Failed to write to a Player %q: %s\n", player.Rune, err)
		return 0, nil
	}

	message, err := player.read()
	if err != nil {
		fmt.Printf("Failed to read from a Player %q: %s\n", player.Rune, err)
		return 0, nil
	}

	messageInt, err := strconv.Atoi(strings.Trim(message, "\n"))
	if err != nil {
		fmt.Printf("Failed to parse Player %q move: %s\n", player.Rune, err)
		return 0, nil
	}

	return messageInt, nil
}
