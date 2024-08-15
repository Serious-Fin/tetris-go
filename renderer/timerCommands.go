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

func deleteRowCmd() tea.Cmd {
	return tea.Tick(85*time.Millisecond, func(t time.Time) tea.Msg {
		return playAnimationMsg(deleteRowAnimation)
	})
}

func gameOverCmd() tea.Cmd {
	return tea.Tick(20*time.Millisecond, func(t time.Time) tea.Msg {
		return playAnimationMsg(gameOverAnimation)
	})
}
