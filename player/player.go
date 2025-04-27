package player

import (
	"bufio"
	"errors"
	"os"
	"tictactoe/deck"
)

type Player interface {
	Choice(d *deck.Deck)
	GetMark() int
}

func PreparePlayers() []Player {
	players := make([]Player, 0)

	inputReader := bufio.NewReader(os.Stdin)
	players = append(players, NewPlayablePlayer(inputReader))

	players = append(players, NewNPC())

	return players
}

func ChangePlayer(curPlayer Player, players []Player) (Player, error) {
	for _, player := range players {
		if curPlayer != player {
			return player, nil
		}
	}

	return nil, errors.New("Another player not found")
}
