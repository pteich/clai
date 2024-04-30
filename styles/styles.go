package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Description = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{
		Light: "#B40BB4",
		Dark:  "#E7B4E7",
	}).MarginTop(1).MarginBottom(1)

	Prompt = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.AdaptiveColor{
		Light: "#128321",
		Dark:  "#54A05A",
	})

	Command = lipgloss.NewStyle().Bold(false)
)
