package renderer

import (
	"aug/tetris/game"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type gameSession struct {
	board                     *game.GameBoard
	currentFigure             game.Figure
	isFigureActive            bool
	isCleaningBoardInProgress bool
	rowToDeleteNext           int
	timesRowFlashedBefore     int
}

type tickMsg time.Time
type rowBeepMsg int

const rowFlashesBeforeDisappearing = 4

func initialModel(board *game.GameBoard) *gameSession {
	return &gameSession{
		board:                     board,
		isFigureActive:            true,
		currentFigure:             *game.GetRandomFigure(),
		isCleaningBoardInProgress: false,
		rowToDeleteNext:           -1,
		timesRowFlashedBefore:     0,
	}
}

func (s gameSession) Init() tea.Cmd {
	return tickCmd()
}

func (s *gameSession) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case rowBeepMsg:
		s.invertRowColor()
		s.timesRowFlashedBefore += 1
		return s.beepOrContinueGame()

	case tickMsg:
		if !s.isFigureActive {
			s.startCleaningFullRows()
			if s.isCleaningBoardInProgress {
				return s, rowBeepCmd()
			}
			s.dropNewFigure()
		}
		s.moveAndSolidifyIfCollision()
		return s, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit

		case "right":
			s.tryMoveFigure(s.board.MoveRight)

		case "left":
			s.tryMoveFigure(s.board.MoveLeft)

		case "up":
			s.tryMoveFigure(s.board.Rotate)

		case "down":
			s.moveAndSolidifyIfCollision()
		}
	}
	return s, nil
}

func (s *gameSession) View() string {
	return s.board.StringifyBoard()
}

func StartGame(board *game.GameBoard) {
	p := tea.NewProgram(initialModel(board))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
