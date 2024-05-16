package tictac

import (
	"fmt"
)

type Game struct {
	players    Players
	state      State
	turn       bool
	field      [9]rune
	xFieldMask int16
	oFieldMask int16
}

func NewGame(players Players) Game {
	players.Message("Hello, Players! Game is starting!")
	game := Game{
		players:    players,
		state:      Playing,
		field:      newField(),
		xFieldMask: 0,
		oFieldMask: 0,
	}
	return game
}

func (game *Game) Move() error {
	var player Player

	if game.turn {
		player = game.players[0]
	} else {
		player = game.players[1]
	}

	move, err := player.Move()
	if err != nil {
		return err
	}
	game.SetFieldCell(move, player)

	return nil
}

func (game *Game) Check() {
	var winner rune

	// Check 'O' state
	switch game.oFieldMask {
	case 7, 56, 448, 73, 146, 293, 273, 84: // Field states represented by a bit-wise field mask
		winner = 'O'
	}

	// Check 'X' state
	switch game.xFieldMask {
	case 7, 56, 448, 73, 146, 293, 273, 84: // Field states represented by a bit-wise field mask
		winner = 'X'
	}

	if winner == 'X' {
		game.state = XWins
	} else if winner == 'O' {
		game.state = OWins
	} else if game.xFieldMask^game.oFieldMask == 511 {
		game.state = Draw
	} else {
		game.state = Playing
	}
}

func (game *Game) FieldSync() {
	for i := 0; i < 9; i++ {
		if game.xFieldMask&(1<<i) > 0 {
			game.field[i] = 'X'
		} else if game.oFieldMask&(1<<i) > 0 {
			game.field[i] = 'O'
		}
	}
}

func (game *Game) SetFieldCell(move int, player Player) error {
	if move > 9 || move < 1 {
		return fmt.Errorf("Field move must be between 1 and 9")
	}

	fieldBits := game.oFieldMask ^ game.xFieldMask

	if fieldBits&(1<<(move-1)) > 0 {
		return fmt.Errorf("Field cell is busy")
	}

	switch player.Rune {
	case 'O':
		game.oFieldMask = game.oFieldMask | (1 << (move - 1))
	case 'X':
		game.xFieldMask = game.xFieldMask | (1 << (move - 1))
	}

	game.FieldSync()

	return nil
}
