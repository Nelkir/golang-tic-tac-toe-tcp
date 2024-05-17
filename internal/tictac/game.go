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
	players.Message("Hello, Players! Game is starting!\n")
	players[0].Greeting()
	players[1].Greeting()
	game := Game{
		players:    players,
		state:      Playing,
		field:      NewField(),
		xFieldMask: 0,
		oFieldMask: 0,
	}
	return game
}

func NewField() [9]rune {
	var field [9]rune
	for i := range field {
		field[i] = '_'
	}

	return field
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

	if !player.Connected {
		game.state = Error
	}

	err = game.SetFieldCell(move, player)
	if err != nil {
		return fmt.Errorf("Failed to set field cell: %s\n", err)
	}

	game.Check()

	game.turn = !game.turn
	game.players.Message("Still Playing!\n")
	return nil
}

func (game *Game) Check() {
	var winner rune

	// Check 'O' state
	switch game.oFieldMask {
	case int16(0b111000000),
		int16(0b000111000),
		int16(0b000000111),
		int16(0b100100100),
		int16(0b010010010),
		int16(0b001001001),
		int16(0b100010001),
		int16(0b001010100): // Field states represented by a bit-wise field mask
		winner = 'O'
	}

	// Check 'X' state
	switch game.xFieldMask {
	case int16(0b111000000),
		int16(0b000111000),
		int16(0b000000111),
		int16(0b100100100),
		int16(0b010010010),
		int16(0b001001001),
		int16(0b100010001),
		int16(0b001010100): // Field states represented by a bit-wise field mask
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
		} else {
			game.field[i] = '_'
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

func (game *Game) PrettyField() string {
	screen := ""

	for i, cell := range game.field {
		switch i {
		case 2, 5, 8:
			screen += fmt.Sprintf(" %c \t-\t%d | %d | %d \n", cell, i-1, i, i+1)
		default:
			screen += fmt.Sprintf(" %c |", cell)
		}
	}

	return screen
}
