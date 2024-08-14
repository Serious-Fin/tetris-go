package game

func (g *GameBoard) CollisionDownDetected(f *Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if g.cellIsFilled(f.MiddlePos.Row+point.Row+1, f.MiddlePos.Col+point.Col) {
			return true
		}
	}
	return false
}

func (g *GameBoard) CollisionRightDetected(f *Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if g.cellIsFilled(f.MiddlePos.Row+point.Row, f.MiddlePos.Col+point.Col+1) {
			return true
		}
	}
	return false
}

func (g *GameBoard) CollisionLeftDetected(f *Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if g.cellIsFilled(f.MiddlePos.Row+point.Row, f.MiddlePos.Col+point.Col-1) {
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
