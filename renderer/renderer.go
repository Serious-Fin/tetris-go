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

		m.MoveDownLogic()

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

		case "down":
			m.MoveDownLogic()
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

func (g *gameSession) MoveDownLogic() {
	if g.gameBoard.CollisionDetected(&g.currentFigure, game.Point{Row: 1, Col: 0}, g.currentFigure.GeometryIndex) {
		g.gameBoard.DrawFigureAs(&g.currentFigure, game.FilledCell)
		g.figureActive = false

		if g.IsGameOver(&g.currentFigure) {
			os.Exit(1)
		}
	} else {
		g.TryMoveFigure(g.gameBoard.MoveDown)
	}
}

func (g *gameSession) IsGameOver(f *game.Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if f.MiddlePos.Row+point.Row < 0 {
			return true
		}
	}
	return false
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
	return tea.Tick(500*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
