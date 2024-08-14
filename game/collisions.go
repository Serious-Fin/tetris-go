package game

func (f *Figure) findLeftmostOffset() int {
	leftmostOffset := 0
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if point.Col < leftmostOffset {
			leftmostOffset = point.Col
		}
	}
	return leftmostOffset
}

func (f *Figure) findRightmostOffset() int {
	rightmostOffset := 0
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if point.Col > rightmostOffset {
			rightmostOffset = point.Col
		}
	}
	return rightmostOffset
}

func (g *GameBoard) leftWallCollisionDetected(f *Figure) bool {
	leftmostOffset := f.findLeftmostOffset()
	return f.MiddlePos.Col+leftmostOffset-1 < 0
}

func (g *GameBoard) rightWallCollisionDetected(f *Figure) bool {
	rightmostOffset := f.findRightmostOffset()
	return f.MiddlePos.Col+rightmostOffset+1 >= g.Width
}

func (g *GameBoard) CollisionDownDetected(f *Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if g.cellIsFilled(f.MiddlePos.Row+point.Row+1, f.MiddlePos.Col+point.Col) {
			return true
		}
	}
	return false
}

func (g *GameBoard) cellIsFilled(row, col int) bool {
	// if cell is above map, ignore it
	if row < 0 {
		return false
	}

	return row >= g.Height || g.Board[row][col] == FilledCell
}
