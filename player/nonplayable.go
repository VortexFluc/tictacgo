package player

import (
	"fmt"
	"math/rand"
	"tictactoe/board"
)

type NPC struct {
	mark int
}

func NewNPC() *NPC {
	return &NPC{
		mark: board.O,
	}
}

func (p *NPC) Choice(d *board.Board) {
	fmt.Println("NPC choice:")
	emptyCells := d.EmptyCells()

	selectedCell := emptyCells[rand.Intn(len(emptyCells))]

	if err := d.SetCell(selectedCell.Row, selectedCell.Col, p.mark); err != nil {
		fmt.Println("Error setting cell", selectedCell.Row, selectedCell.Col, err)
	}
}

func (p *NPC) String() string {
	return fmt.Sprint("NPC")
}

func (p *NPC) GetMark() int {
	return p.mark
}
