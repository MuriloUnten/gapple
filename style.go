package main

import (
	"github.com/charmbracelet/lipgloss"
)

func mainPaneStyle(m Model) lipgloss.Style {
	return lipgloss.NewStyle().
		Height(m.windowHeight-5).
		Width(m.windowWidth-2).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)
}

func hotkeysPaneStyle(m Model) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(m.windowWidth-2).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)
}

func mainTimerStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("#A8C4FF"))
}

func timersInfoStyle() lipgloss.Style {
	return lipgloss.NewStyle().MarginTop(1)
}

func statusTextStyle() lipgloss.Style {
	return lipgloss.NewStyle().Margin(1)
}
