package deck

import "fmt"

const (
	EMPTY = iota
	X
	O
)

type Cell struct {
	row int
	col int
	val int
}

type Deck struct {
	data [][]Cell
}

func NewDeck(size int) Deck {
	data := make([][]Cell, size)
	for i := range data {
		data[i] = make([]Cell, size)
	}

	return Deck{
		data: data,
	}
}

func (deck Deck) String() string {
	result := ""
	for _, row := range deck.data {
		for _, cell := range row {
			switch cell.val {
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
