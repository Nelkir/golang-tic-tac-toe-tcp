package tictac

import "fmt"

const (
	Playing = iota
	XWins
	OWins
	Draw
	Error
)

type State int

func Start(player1 Player, player2 Player) {
	players := Players{player1, player2}

	game := NewGame(players)

	players.Message(game.PrettyField())
	for game.state == Playing {
		err := game.Move()
		if err != nil {
			fmt.Printf("Failed to perform move: %s\n", err)
		}

		switch game.state {
		case XWins:
			game.players.Message("X Wins!\n")
			return
		case OWins:
			game.players.Message("O Wins!\n")
		case Draw:
			game.players.Message("Draw!\n")
		case Error:
			fmt.Printf("Player left game!\n")
		}
		players.Message(game.PrettyField())
	}
}
