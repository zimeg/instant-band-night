package display

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Title returns title in a stable section purple
func Title(title string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("141")).
		Render(title)
}

// Secondary returns secondary in a passive cloud
func Secondary(secondary string, a ...any) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("244")).
		Render(fmt.Sprintf(secondary, a...))
}

// Faint returns faint in dimmed markings
func Faint(faint string) string {
	return lipgloss.NewStyle().
		Faint(true).
		Render(faint)
}
