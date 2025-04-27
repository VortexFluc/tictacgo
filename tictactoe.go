package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const (
	EMPTY = iota
	X
	O
)

const (
	SET = iota
	QUIT
	UNKNOWN
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Hello! Welcome to test tictactoe game!")
	fmt.Println("Preparing deck for you...")

	deck := NewDeck()

	fmt.Println("Deck initialized. Here's the deck:")
	PrintDeck(deck)

	for {
		fmt.Println("Please enter a command. Available commands: set, quit")
		input := readInput(inputReader)
		command := GetCommand(input)

		switch command {
		case SET:
			SetValue(inputReader, &deck)
		case QUIT:
			break
		default:
			fmt.Println("Invalid command")
			continue
		}

		PrintDeck(deck)
		isXWin := processStage(&deck, X)
		if isXWin {
			fmt.Println("You won!")
			break
		}

		fmt.Println("Player 2 is making a choice...")
		performAIStep(&deck)
		PrintDeck(deck)
		isOWin := processStage(&deck, O)
		if isOWin {
			fmt.Println("Player 2 won!")
			break
		}

	}

	fmt.Print("Bye!")
}

func performAIStep(deck *[3][3]int) {
	availableCells := make([]DeckCell, 0)

	for rowIdx, row := range deck {
		for colIdx, cell := range row {
			if cell == EMPTY {
				availableCells = append(availableCells, DeckCell{
					row: rowIdx,
					col: colIdx,
				})
			}
		}
	}

	aiChoice := availableCells[rand.Intn(len(availableCells))]

	deck[aiChoice.row][aiChoice.col] = O
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

func GetCommand(com string) int {
	if com == "quit" {
		return QUIT
	}

	if com == "set" {
		return SET
	}

	return UNKNOWN
}

func SetValue(ir *bufio.Reader, deck *[3][3]int) {
	fmt.Println("Select a row (from 1 to 3)")
	rStr := readInput(ir)
	if row, err := strconv.Atoi(rStr); err == nil {
		fmt.Println("Select a column (from 1 to 3)")
		cStr := readInput(ir)
		if column, err := strconv.Atoi(cStr); err == nil {
			deck[row-1][column-1] = X
		}
	}
}

func NewDeck() [3][3]int {
	newDeck := [3][3]int{}
	return newDeck
}

func readInput(ir *bufio.Reader) string {
	command, _ := ir.ReadString('\n')
	command = strings.TrimSpace(command)

	return command
}

func PrintDeck(deck [3][3]int) {
	for _, row := range deck {
		for _, cell := range row {
			switch cell {
			case X:
				fmt.Print("X")
			case O:
				fmt.Print("O")
			case EMPTY:
				fmt.Print("▢")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
