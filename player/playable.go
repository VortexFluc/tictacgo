package player

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"tictactoe/board"
)

type RealPlayer struct {
	mark   int
	input  *bufio.Reader
	exitCh chan bool
}

func NewPlayablePlayer(input *bufio.Reader, exitCh chan bool) *RealPlayer {
	return &RealPlayer{
		mark:   board.X,
		input:  input,
		exitCh: exitCh,
	}
}

func (p *RealPlayer) Choice(d *board.Board) {
	fmt.Println("Please enter a command. Available commands: set, quit")
	input := readInput(p.input)
	command := getCommand(input)

	switch command {
	case SET:
		if err := p.setValue(p.input, d); err != nil {
			fmt.Println("Error setting value: ", err)
		}
	case QUIT:
		p.exitCh <- true
	default:
		fmt.Println("Invalid command")
	}
}

func (p *RealPlayer) GetMark() int {
	return p.mark
}

func (p *RealPlayer) String() string {
	return fmt.Sprint("You")
}

func (p *RealPlayer) setValue(ir *bufio.Reader, d *board.Board) error {
	fmt.Println("Select a row (from 1 to 3)")
	rStr := readInput(ir)
	if row, err := strconv.Atoi(rStr); err == nil {
		fmt.Println("Select a column (from 1 to 3)")
		cStr := readInput(ir)
		if column, err := strconv.Atoi(cStr); err == nil {
			return d.SetCell(row-1, column-1, p.mark)
		} else {
			return err
		}
	} else {
		return err
	}

	return nil
}

func readInput(ir *bufio.Reader) string {
	command, _ := ir.ReadString('\n')
	command = strings.TrimSpace(command)

	return command
}

const (
	SET = iota
	QUIT
	UNKNOWN
)

func getCommand(com string) int {
	if com == "quit" {
		return QUIT
	}

	if com == "set" {
		return SET
	}

	return UNKNOWN
}
