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
	if s.timesRowFlashedBefore >= rowFlashesBeforeDisappearing {
		return s.endRowCleaning()
	} else {
		return s, rowBeepCmd()
	}
}

func (s *gameSession) endRowCleaning() (tea.Model, tea.Cmd) {
	s.removeRow(s.rowToDeleteNext)
	s.isCleaningBoardInProgress = false
	return s, tickCmd()
}

func (s *gameSession) startCleaningFullRows() {
	for index, row := range s.board.Board {
		if isRowFull(row) {
			s.isCleaningBoardInProgress = true
			s.rowToDeleteNext = index
			s.timesRowFlashedBefore = 0
		}
	}
}
