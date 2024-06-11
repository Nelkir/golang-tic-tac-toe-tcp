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
			return
		case Draw:
			game.players.Message("Draw!\n")
			return
		case Error:
			game.players.Message("Player left the game!\n")
			fmt.Printf("Player left the game!\n")
			return
		}
		players.Message(game.PrettyField())
	}
}
