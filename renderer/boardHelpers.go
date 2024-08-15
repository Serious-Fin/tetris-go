package renderer

import "aug/tetris/game"

func (s *gameSession) removeRow(index int) {
	for row := index; row > 0; row-- {
		for col := 0; col < s.board.Width; col++ {
			s.board.Board[row][col] = s.board.Board[row-1][col]
		}
	}
}

func isRowFull(row []int) bool {
	for _, status := range row {
		if status != game.FilledCell {
			return false
		}
	}
	return true
}

func isGameOver(f *game.Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if f.MiddlePos.Row+point.Row < 0 {
			return true
		}
	}
	return false
}

func (s *gameSession) invertRowColor() {
	for index, value := range s.board.Board[s.rowToDeleteNext] {
		if value == game.FilledCell {
			s.board.Board[s.rowToDeleteNext][index] = game.EmptyCell
		} else {
			s.board.Board[s.rowToDeleteNext][index] = game.FilledCell
		}
	}
}

func (s *gameSession) halfBoardLength() int {
	return (s.board.Width / 2)
}

func (s *gameSession) paintTwoPixelsMirrored() {
	s.board.Board[s.aParams.lastPaintedPoint.Row][s.aParams.lastPaintedPoint.Col] = game.Border
	s.board.Board[s.aParams.lastPaintedPoint.Row][s.board.Width-s.aParams.lastPaintedPoint.Col-1] = game.Border
}
