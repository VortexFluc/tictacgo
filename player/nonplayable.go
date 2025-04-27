package player

import (
	"fmt"
	"math/rand"
	"tictactoe/deck"
)

type NPC struct {
	mark int
}

func NewNPC() *NPC {
	return &NPC{
		mark: deck.O,
	}
}

func (p *NPC) Choice(d *deck.Deck) {
	emptyCells := d.EmptyCells()

	selectedCell := emptyCells[rand.Intn(len(emptyCells))]

	if err := d.SetCell(selectedCell.Row, selectedCell.Col, p.mark); err != nil {
		fmt.Println("Error setting cell", selectedCell.Row, selectedCell.Col, err)
	}
}
