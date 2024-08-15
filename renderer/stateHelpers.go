package renderer

import (
	"aug/tetris/game"

	tea "github.com/charmbracelet/bubbletea"
)

func (s *gameSession) dropNewFigure() {
	s.currentFigure = *game.GetRandomFigure()
	s.isFigureActive = true
}

func (s *gameSession) beepOrContinueGame() (tea.Model, tea.Cmd) {
	if s.aParams.timesRowFlashedBefore >= rowFlashesBeforeDisappearing {
		return s.endRowCleaning()
	} else {
		return s, deleteRowCmd()
	}
}

func (s *gameSession) endRowCleaning() (tea.Model, tea.Cmd) {
	s.removeRow(s.rowToDeleteNext)
	s.isPlayingAnimation = false
	return s, tickCmd()
}

func (s *gameSession) startCleaningFullRows() {
	for index, row := range s.board.Board {
		if isRowFull(row) {
			s.isPlayingAnimation = true
			s.rowToDeleteNext = index
			s.aParams.timesRowFlashedBefore = 0
		}
	}
}

func (s *gameSession) pickNextPixel() {
	s.aParams.lastPaintedPoint.Col -= 1
	if s.aParams.lastPaintedPoint.Col < s.halfBoardLength()-1 {
		s.aParams.lastPaintedPoint.Row -= 1
		s.aParams.lastPaintedPoint.Col = s.board.Width - 1
	}
}
