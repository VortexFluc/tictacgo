package main

import (
	"fmt"
	"os"
	"tictactoe/board"
	"tictactoe/player"
)

func main() {
	fmt.Println("Hello! Welcome to test TicTacToe game!")
	fmt.Println("Preparing board for you...")
	d := board.NewBoard(3)

	fmt.Println("Board initialized. Here's the board:")
	fmt.Println(d)

	fmt.Println("Initializing players.")
	players, exitCh := player.PreparePlayers()

	curPlayer := players[0]

	go func() {
		exit := <-exitCh

		if exit {
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		}
	}()

	for {
		curPlayer.Choice(&d)
		fmt.Println(d)

		result := board.AnalyzeBoard(&d, curPlayer.GetMark())

		if result == board.WIN {
			fmt.Println(curPlayer, "won!")
			break
		}

		if result == board.DRAW {
			fmt.Println("Draw!")
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
