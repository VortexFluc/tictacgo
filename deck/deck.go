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
	Row int
	Col int
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

	for rowIdx, row := range data {
		for colIdx, _ := range row {
			data[rowIdx][colIdx] = Cell{Row: rowIdx, Col: colIdx, Val: EMPTY}
		}
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

func (deck *Deck) CellsFilledWith(mark int) []Cell {
	result := make([]Cell, 0)

	for _, row := range deck.Data {
		for _, cell := range row {
			if cell.Val == mark {
				result = append(result, cell)
			}
		}
	}

	return result
}

func (deck *Deck) EmptyCells() []Cell {
	result := make([]Cell, 0)

	for _, row := range deck.Data {
		for _, cell := range row {
			if cell.Val == EMPTY {
				result = append(result, cell)
			}
		}
	}

	return result
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
