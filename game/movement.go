package game

func (g *GameBoard) MoveDown(f *Figure) {
	f.MiddlePos.Row += 1
}

func (g *GameBoard) MoveLeft(f *Figure) {
	if g.leftWallCollisionDetected(f) {
		return
	}

	// TODO: Check collision with any blocks on the left

	f.MiddlePos.Col -= 1
}

func (g *GameBoard) MoveRight(f *Figure) {
	if g.rightWallCollisionDetected(f) {
		return
	}

	// TODO: Check collision with any blocks on the right

	f.MiddlePos.Col += 1
}

func (g *GameBoard) Rotate(f *Figure) {
	// TODO: rotation collision check
	f.GeometryIndex = (f.GeometryIndex + 1) % len(f.Geometries)
}
