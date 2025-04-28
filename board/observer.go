package board

const (
	CONTINUE_GAME = iota
	WIN
	DRAW
)

type ObserverResult int

func AnalyzeBoard(d *Board, mark int) ObserverResult {

	cells := d.CellsFilledWith(mark)

	if findWinRow(cells) {
		return WIN
	}

	if findWinColumn(cells) {
		return WIN
	}

	if findWinDiagonal(d, mark) {
		return WIN
	}

	if len(d.EmptyCells()) == 0 {
		return DRAW
	}

	return CONTINUE_GAME
}

func findWinRow(c []Cell) bool {
	rowMap := make(map[int][]Cell)
	for _, sr := range c {
		if v, ok := rowMap[sr.Row]; ok {
			v = append(v, sr)
			rowMap[sr.Row] = v
		} else {
			rowMap[sr.Row] = []Cell{sr}
		}
	}

	for _, v := range rowMap {
		if len(v) == 3 {
			return true
		}
	}

	return false
}

func findWinColumn(c []Cell) bool {
	colMap := make(map[int][]Cell)
	for _, sr := range c {
		if v, ok := colMap[sr.Col]; ok {
			v = append(v, sr)
			colMap[sr.Col] = v
		} else {
			colMap[sr.Col] = []Cell{sr}
		}
	}

	for _, v := range colMap {
		if len(v) == 3 {
			return true

		}
	}

	return false
}

func findWinDiagonal(d *Board, mark int) bool {
	var foundDiagonal bool
	for _, diag := range d.Diagonals() {
		for _, cell := range diag {
			if cell.Val != mark {
				foundDiagonal = false
				break
			}
			foundDiagonal = true
		}

		if foundDiagonal {
			break
		}
	}

	return foundDiagonal
}
