package main

import tea "github.com/charmbracelet/bubbletea"

// quit
func (m Model) QuitAction() (Model, tea.Cmd) {
	return m, tea.Quit
}

// pause
func (m Model) PauseAction() (Model, tea.Cmd) {
	m.timer.TogglePause()
	if m.status == NONE {
		m.status = FOCUS
	}

	return m, nil
}

// restart
func (m Model) RestartAction() (Model, tea.Cmd) {
	m.timer.Reset()

	return m, nil
}

// clear
func (m Model) ClearAction() (Model, tea.Cmd) {
	m.status = NONE
	m.focusSeconds = 0
	m.chillSeconds = 0
	m.timer.Preset(0)

	return m, nil
}

// next
func (m Model) NextAction() (Model, tea.Cmd) {
	if m.status == FOCUS {
		m.timer.Preset(m.chillSeconds)
		m.status = CHILL

	} else if m.status == CHILL {
		m.timer.Preset(m.focusSeconds)
		m.status = FOCUS
	}

	return m, nil
}

// TODO Implement
// set
func (m Model) SetAction() (Model, tea.Cmd) {

	return m, nil
}

// TODO Implement
// help
func (m Model) HelpAction() (Model, tea.Cmd) {

	return m, nil
}

func (m Model) UndefinedAction() (Model, tea.Cmd) {
	return m, nil
}
