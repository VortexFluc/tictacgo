package main

import (
	"fmt"
	"tictactoe/deck"
	"tictactoe/player"
)

const (
	SET = iota
	QUIT
	UNKNOWN
)

func main() {
	fmt.Println("Hello! Welcome to test TicTacToe game!")
	fmt.Println("Preparing deck for you...")
	d := deck.NewDeck(3)

	fmt.Println("Deck initialized. Here's the deck:")
	fmt.Println(d)

	fmt.Println("Initializing players.")
	players := player.PreparePlayers()

	curPlayer := players[0]

	for {
		curPlayer.Choice(&d)
		fmt.Println(d)

		adaptData := adapt(d.Data)
		isXWin := processStage(&adaptData, deck.X)
		if isXWin {
			fmt.Println("You won!")
			break
		}

		if newPlayer, err := player.ChangePlayer(curPlayer, players); err != nil {
			fmt.Println(err)
			break
		} else {
			curPlayer = newPlayer
		}
	}
}

type DeckCell struct {
	row int
	col int
}

func processStage(deck *[3][3]int, mark int) bool {
	result := scanDeck(deck, mark)
	rowMap := make(map[int][]DeckCell)
	for _, sr := range result {
		if v, ok := rowMap[sr.row]; ok {
			v = append(v, sr)
			rowMap[sr.row] = v
		} else {
			rowMap[sr.row] = []DeckCell{sr}
		}
	}

	colMap := make(map[int][]DeckCell)
	for _, sr := range result {
		if v, ok := colMap[sr.col]; ok {
			v = append(v, sr)
			colMap[sr.col] = v
		} else {
			colMap[sr.col] = []DeckCell{sr}
		}
	}

	var win bool
	for _, v := range rowMap {
		if len(v) == 3 {
			win = true
			break
		}
	}

	for _, v := range colMap {
		if len(v) == 3 {
			win = true
			break
		}
	}

	return win
}

func adapt(d [][]deck.Cell) [3][3]int {
	data := [3][3]int{}
	for i, row := range d {
		for j, cell := range row {
			data[i][j] = cell.Val
		}
	}

	return data
}

func scanDeck(deck *[3][3]int, mark int) []DeckCell {
	result := make([]DeckCell, 0)
	for rowIdx, row := range deck {
		for colIdx, cell := range row {
			if cell == mark {
				result = append(result, DeckCell{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}

	return result
}
