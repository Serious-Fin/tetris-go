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
		if s.timesRowFlashedBefore >= 4 {
			return s.endRowCleaning()
		} else {
			return s, rowBeepCmd()
		}

	case tickMsg:
		if !s.isFigureActive {
			s.cleanFullRows()
			if s.isCleaningBoardInProgress {
				return s, rowBeepCmd()
			}
			s.currentFigure = *game.GetRandomFigure()
			s.isFigureActive = true
		}

		s.MoveDownLogic()

		return s, tickCmd()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return s, tea.Quit

		case "right":
			s.TryMoveFigure(s.board.MoveRight)

		case "left":
			s.TryMoveFigure(s.board.MoveLeft)

		case "up":
			s.TryMoveFigure(s.board.Rotate)

		case "down":
			s.MoveDownLogic()
		}
	}
	return s, nil
}

func (m *gameSession) moveFigure(movementFunction func(*game.Figure)) {
	m.board.DrawFigureAs(&m.currentFigure, game.EmptyCell)
	movementFunction(&m.currentFigure)
	m.board.DrawFigureAs(&m.currentFigure, m.currentFigure.BlockType)
}

func (m *gameSession) TryMoveFigure(movementFunction func(*game.Figure)) {
	if m.isFigureActive {
		m.moveFigure(movementFunction)
	}
}

func (g *gameSession) MoveDownLogic() {
	if !g.isFigureActive || g.isCleaningBoardInProgress {
		return
	}

	if g.board.CollisionDetected(&g.currentFigure, game.Point{Row: 1, Col: 0}, g.currentFigure.GeometryIndex) {
		g.board.DrawFigureAs(&g.currentFigure, game.FilledCell)
		g.isFigureActive = false

		if g.IsGameOver(&g.currentFigure) {
			os.Exit(1)
		}
	} else {
		g.TryMoveFigure(g.board.MoveDown)
	}
}

func (s *gameSession) endRowCleaning() (tea.Model, tea.Cmd) {
	s.removeRow(s.rowToDeleteNext)
	s.isCleaningBoardInProgress = false
	return s, tickCmd()
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

func (s *gameSession) cleanFullRows() {
	for index, row := range s.board.Board {
		if isRowFull(row) {
			s.isCleaningBoardInProgress = true
			s.rowToDeleteNext = index
			s.timesRowFlashedBefore = 0
		}
	}
}

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

func (g *gameSession) IsGameOver(f *game.Figure) bool {
	for _, point := range f.Geometries[f.GeometryIndex].Points {
		if f.MiddlePos.Row+point.Row < 0 {
			return true
		}
	}
	return false
}

func (g *gameSession) View() string {
	return g.board.StringifyBoard()
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

func rowBeepCmd() tea.Cmd {
	return tea.Tick(85*time.Millisecond, func(t time.Time) tea.Msg {
		return rowBeepMsg(1)
	})
}
