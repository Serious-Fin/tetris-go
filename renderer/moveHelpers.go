package renderer

import (
	"aug/tetris/game"
)

func (s *gameSession) moveFigure(movementFunction func(*game.Figure)) {
	s.board.DrawFigureAs(&s.currentFigure, game.EmptyCell)
	movementFunction(&s.currentFigure)
	s.board.DrawFigureAs(&s.currentFigure, s.currentFigure.BlockType)
}

func (s *gameSession) tryMoveFigure(movementFunction func(*game.Figure)) {
	if s.canMoveFigure() {
		s.moveFigure(movementFunction)
	}
}

func (s *gameSession) canMoveFigure() bool {
	return s.isFigureActive && !s.isPlayingAnimation
}

func (s *gameSession) moveAndSolidifyIfCollision() {
	if !s.canMoveFigure() {
		return
	}

	if s.board.CollisionDetected(&s.currentFigure, game.Point{Row: 1, Col: 0}, s.currentFigure.GeometryIndex) {
		s.board.DrawFigureAs(&s.currentFigure, game.FilledCell)
		s.isFigureActive = false

		if isGameOver(&s.currentFigure) {
			s.isPlayingAnimation = true
		}
	} else {
		s.tryMoveFigure(s.board.MoveDown)
	}
}
