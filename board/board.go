package board

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

type Board struct {
	Data [][]Cell
	size int
}

func NewBoard(size int) Board {
	data := make([][]Cell, size)
	for i := range data {
		data[i] = make([]Cell, size)
	}

	for rowIdx, row := range data {
		for colIdx, _ := range row {
			data[rowIdx][colIdx] = Cell{Row: rowIdx, Col: colIdx, Val: EMPTY}
		}
	}

	return Board{
		Data: data,
		size: size,
	}
}

func (d *Board) SetCell(row, col, val int) error {
	cell := d.Data[row][col]

	if row > d.size-1 {
		return errors.New("Row out of range")
	}

	if col > d.size-1 {
		return errors.New("Col out of range")
	}

	if cell.Val != EMPTY {
		return errors.New("Cell not empty")
	}

	cell.Val = val
	d.Data[row][col] = cell
	return nil
}

func (d *Board) CellsFilledWith(mark int) []Cell {
	result := make([]Cell, 0)

	for _, row := range d.Data {
		for _, cell := range row {
			if cell.Val == mark {
				result = append(result, cell)
			}
		}
	}

	return result
}

func (d *Board) EmptyCells() []Cell {
	result := make([]Cell, 0)

	for _, row := range d.Data {
		for _, cell := range row {
			if cell.Val == EMPTY {
				result = append(result, cell)
			}
		}
	}

	return result
}

func (d *Board) Diagonals() [][]Cell {
	result := make([][]Cell, 0)
	firstDiagonal := make([]Cell, 0)
	for i := 0; i < d.size; i++ {
		firstDiagonal = append(firstDiagonal, d.Data[i][i])
	}

	result = append(result, firstDiagonal)

	secondDiagonal := make([]Cell, 0)
	for i := 0; i < d.size; i++ {
		secondDiagonal = append(secondDiagonal, d.Data[d.size-i-1][i])
	}
	result = append(result, secondDiagonal)

	return result
}

func (d Board) String() string {
	result := ""
	for _, row := range d.Data {
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
