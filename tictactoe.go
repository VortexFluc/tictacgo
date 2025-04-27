package main

import (
	"fmt"
	"os"
	"tictactoe/deck"
	"tictactoe/player"
)

func main() {
	fmt.Println("Hello! Welcome to test TicTacToe game!")
	fmt.Println("Preparing deck for you...")
	d := deck.NewDeck(3)

	fmt.Println("Deck initialized. Here's the deck:")
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

		result := deck.AnalyzeDeck(&d, curPlayer.GetMark())

		if result == deck.WIN {
			fmt.Println(curPlayer, "won!")
			break
		}

		if result == deck.DRAW {
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
