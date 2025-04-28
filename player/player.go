package player

import (
	"bufio"
	"errors"
	"os"
	"tictactoe/board"
)

type Player interface {
	Choice(d *board.Board)
	GetMark() int
}

func PreparePlayers() ([]Player, chan bool) {
	players := make([]Player, 0)

	inputReader := bufio.NewReader(os.Stdin)
	exitCh := make(chan bool)
	players = append(players, NewPlayablePlayer(inputReader, exitCh))

	players = append(players, NewNPC())

	return players, exitCh
}

func ChangePlayer(curPlayer Player, players []Player) (Player, error) {
	for _, player := range players {
		if curPlayer != player {
			return player, nil
		}
	}

	return nil, errors.New("Another player not found")
}
