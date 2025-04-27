package deck

import (
	"errors"
	"fmt"
)

const (
	EMPTY = iota
	X
	O
)

type Cell struct {
	row int
	col int
	Val int
}

type Deck struct {
	Data [][]Cell
	size int
}

func NewDeck(size int) Deck {
	data := make([][]Cell, size)
	for i := range data {
		data[i] = make([]Cell, size)
	}

	return Deck{
		Data: data,
		size: size,
	}
}

func (deck *Deck) SetCell(row, col, val int) error {
	cell := deck.Data[row][col]

	if row > deck.size-1 {
		return errors.New("Row out of range")
	}

	if col > deck.size-1 {
		return errors.New("Col out of range")
	}

	if cell.Val != EMPTY {
		return errors.New("Cell not empty")
	}

	cell.Val = val
	deck.Data[row][col] = cell
	return nil
}

func (deck Deck) String() string {
	result := ""
	for _, row := range deck.Data {
		for _, cell := range row {
			switch cell.Val {
			case X:
				result += fmt.Sprint("X")
			case O:
				result += fmt.Sprint("O")
			case EMPTY:
				result += fmt.Sprint("▢")
			}
			result += fmt.Sprint(" ")
		}
		result += fmt.Sprint("\n")
	}

	return result
}
