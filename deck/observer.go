package deck

const (
	CONTINUE_GAME = iota
	X_WIN
	O_WIN
	DRAW
)

type ObserverResult int

func AnalyzeDeck(d *Deck) ObserverResult {
	return CONTINUE_GAME
}
