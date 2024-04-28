package display

import (
	"github.com/charmbracelet/lipgloss"
)

// Title returns title in a stable section purple
func Title(title string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("141")).
		Render(title)
}

// Secondary returns secondary in a passive cloud
func Secondary(secondary string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("244")).
		Render(secondary)
}
