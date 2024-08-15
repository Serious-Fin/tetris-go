package renderer

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func tickCmd() tea.Cmd {
	return tea.Tick(450*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func rowBeepCmd() tea.Cmd {
	return tea.Tick(85*time.Millisecond, func(t time.Time) tea.Msg {
		return rowBeepMsg(1)
	})
}
