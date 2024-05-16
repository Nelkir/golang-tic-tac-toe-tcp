package tictac

const (
	Playing = iota
	XWins
	OWins
	Draw
)

type State int

func Start(player1 Player, player2 Player) {
	players := Players{player1, player2}

	game := NewGame(players)

	for game.state == Playing {
		if game.turn {
			game.players[0].Move()
		} else {
		}
	}
}

func newField() [9]rune {
	var field [9]rune
	for _, cell := range field {
		cell = ' '
		_ = cell
	}

	return field
}
