package renderer

import (
	"aug/tetris/game"
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type gameSession struct {
	gameBoard     *game.GameBoard
	currentFigure game.Figure
	figureActive  bool
}

type tickMsg time.Time

func initialModel(board *game.GameBoard) *gameSession {
	return &gameSession{
		gameBoard:     board,
		figureActive:  true,
		currentFigure: game.FigureT,
	}
}

func (m gameSession) Init() tea.Cmd {
	return tickCmd()
}

func (m *gameSession) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		m.MoveFigure(m.gameBoard.MoveDown)
		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "right":
			m.MoveFigure(m.gameBoard.MoveRight)

		case "left":
			m.MoveFigure(m.gameBoard.MoveLeft)

		case "up":
			m.MoveFigure(m.gameBoard.Rotate)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *gameSession) MoveFigure(movementFunction func(*game.Figure)) {
	m.gameBoard.CleanPreviousFigurePosition(&m.currentFigure)
	movementFunction(&m.currentFigure)
	m.gameBoard.DrawFigureOnBoard(&m.currentFigure)
}

func (g *gameSession) View() string {
	return g.gameBoard.StringifyBoard()
}

func StartGame(board *game.GameBoard) {
	p := tea.NewProgram(initialModel(board))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
