package renderer

import (
	"aug/tetris/game"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type gameSession struct {
	board              *game.GameBoard
	currentFigure      game.Figure
	isFigureActive     bool
	isPlayingAnimation bool
	rowToDeleteNext    int
	aParams            animationParams
}

type animationParams struct {
	timesRowFlashedBefore int
	lastPaintedPoint      game.Point
}

type tickMsg time.Time
type playAnimationMsg int

const gameOverAnimation = 1
const deleteRowAnimation = 2
const rowFlashesBeforeDisappearing = 4

func initialModel(board *game.GameBoard) *gameSession {
	return &gameSession{
		board:              board,
		isFigureActive:     true,
		currentFigure:      *game.GetRandomFigure(),
		isPlayingAnimation: false,
		rowToDeleteNext:    -1,
		aParams: animationParams{
			timesRowFlashedBefore: -1,
			lastPaintedPoint: game.Point{
				Row: board.Height - 1,
				Col: board.Width - 1,
			},
		},
	}
}

func (s gameSession) Init() tea.Cmd {
	return tickCmd()
}

func (s *gameSession) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case playAnimationMsg:
		if msg == deleteRowAnimation {
			s.invertRowColor()
			s.aParams.timesRowFlashedBefore += 1
			return s.beepOrContinueGame()
		}

		if msg == gameOverAnimation {
			s.paintTwoPixelsMirrored()
			if s.aParams.lastPaintedPoint.Row == 0 && s.aParams.lastPaintedPoint.Col < s.halfBoardLength() {
				os.Exit(1)
			}
			s.pickNextPixel()

			return s, gameOverCmd()
		}

	case tickMsg:
		if !s.isFigureActive {
			s.startCleaningFullRows()
			if s.isPlayingAnimation {
				return s, deleteRowCmd()
			}
			s.dropNewFigure()
		}
		s.moveAndSolidifyIfCollision()

		if s.isPlayingAnimation {
			return s, gameOverCmd()
		}

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
