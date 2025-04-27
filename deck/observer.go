package deck

const (
	CONTINUE_GAME = iota
	WIN
	DRAW
)

type ObserverResult int

func AnalyzeDeck(d *Deck, mark int) ObserverResult {

	cells := d.CellsFilledWith(mark)

	rowMap := make(map[int][]Cell)
	for _, sr := range cells {
		if v, ok := rowMap[sr.Row]; ok {
			v = append(v, sr)
			rowMap[sr.Row] = v
		} else {
			rowMap[sr.Row] = []Cell{sr}
		}
	}

	colMap := make(map[int][]Cell)
	for _, sr := range cells {
		if v, ok := colMap[sr.Col]; ok {
			v = append(v, sr)
			colMap[sr.Col] = v
		} else {
			colMap[sr.Col] = []Cell{sr}
		}
	}

	for _, v := range rowMap {
		if len(v) == 3 {
			return WIN
		}
	}

	for _, v := range colMap {
		if len(v) == 3 {
			return WIN

		}
	}

	if len(d.EmptyCells()) == 0 {
		return DRAW
	}

	return CONTINUE_GAME
}
