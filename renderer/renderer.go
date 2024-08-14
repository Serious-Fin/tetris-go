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
		currentFigure: *game.GetRandomFigure(),
	}
}

func (m gameSession) Init() tea.Cmd {
	return tickCmd()
}

func (m *gameSession) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tickMsg:
		if !m.figureActive {
			m.currentFigure = *game.GetRandomFigure()
			m.figureActive = true
		}

		if m.gameBoard.CollisionDetected(&m.currentFigure, game.Point{Row: 1, Col: 0}, m.currentFigure.GeometryIndex) {
			m.gameBoard.DrawFigureAs(&m.currentFigure, game.FilledCell)
			m.figureActive = false
		} else {
			m.TryMoveFigure(m.gameBoard.MoveDown)
		}

		return m, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "right":
			m.TryMoveFigure(m.gameBoard.MoveRight)

		case "left":
			m.TryMoveFigure(m.gameBoard.MoveLeft)

		case "up":
			m.TryMoveFigure(m.gameBoard.Rotate)
		}
	}
	return m, nil
}

func (m *gameSession) moveFigure(movementFunction func(*game.Figure)) {
	m.gameBoard.DrawFigureAs(&m.currentFigure, game.EmptyCell)
	movementFunction(&m.currentFigure)
	m.gameBoard.DrawFigureAs(&m.currentFigure, m.currentFigure.BlockType)
}

func (m *gameSession) TryMoveFigure(movementFunction func(*game.Figure)) {
	if m.figureActive {
		m.moveFigure(movementFunction)
	}
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
	return tea.Tick(660*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
