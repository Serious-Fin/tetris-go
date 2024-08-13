package renderer

import (
	"aug/tetris/game"
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type gameSession struct {
	gameBoard *game.GameBoard
}

func initialModel(board *game.GameBoard) *gameSession {
	return &gameSession{
		gameBoard: board,
	}
}

// TODO: Figure out what this thing does
func (m gameSession) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

// TODO: make it so it drives the figure while falling
func (m *gameSession) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			row := rand.Intn(m.gameBoard.Height)
			col := rand.Intn(m.gameBoard.Width)

			if m.gameBoard.Board[row][col] == game.EmptyCell {
				m.gameBoard.Board[row][col] = game.FilledCell
			} else {
				m.gameBoard.Board[row][col] = game.EmptyCell
			}
		}

	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (g *gameSession) View() string {
	return DisplayBoard(g.gameBoard)
}

func Main(board *game.GameBoard) {
	p := tea.NewProgram(initialModel(board))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
