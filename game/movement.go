package game

func (g *GameBoard) MoveDown(f *Figure) {
	f.MiddlePos.Row += 1
}

func (g *GameBoard) MoveLeft(f *Figure) {
	if g.CollisionDetected(f, Point{Row: 0, Col: -1}, f.GeometryIndex) {
		return
	}
	f.MiddlePos.Col -= 1
}

func (g *GameBoard) MoveRight(f *Figure) {
	if g.CollisionDetected(f, Point{Row: 0, Col: 1}, f.GeometryIndex) {
		return
	}
	f.MiddlePos.Col += 1
}

func (g *GameBoard) Rotate(f *Figure) {
	newRotationIndex := (f.GeometryIndex + 1) % len(f.Geometries)

	if g.CollisionDetected(f, Point{Row: 0, Col: 0}, newRotationIndex) {
		return
	}

	f.GeometryIndex = newRotationIndex
}
