package game

func (g *GameBoard) CollisionDetected(f *Figure, offset Point, figureIndex int) bool {
	for _, point := range f.Geometries[figureIndex].Points {
		if g.cellIsFilled(f.MiddlePos.Row+point.Row+offset.Row, f.MiddlePos.Col+point.Col+offset.Col) {
			return true
		}
	}
	return false
}

func (g *GameBoard) cellIsFilled(row, col int) bool {
	// Side walls has highest priority
	if col < 0 || col >= g.Width {
		return true
	}

	// if cell is above map, ignore it
	if row < 0 {
		return false
	}

	// if we hit bottom of screen or another cell
	return row >= g.Height || g.Board[row][col] == FilledCell
}
