package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Status int

const (
	FOCUS Status = iota
	CHILL
	NONE
)

type model struct {
	status Status
	remainingSeconds int
	focusSeconds int
	chillSeconds int
	paused bool
	windowWidth int
	windowHeight int
}

func initialModel() model {
	return model {
		status: NONE,
		remainingSeconds: -1,
		focusSeconds: 0,
		chillSeconds: 0,
		paused: true,
		windowWidth: -1,
		windowHeight: -1,
	}
}

func (m model) Init() tea.Cmd {
    // Just return `nil`, which means "no I/O right now, please."
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width

    case tea.KeyMsg:
        switch msg.String() {
		// quit
        case "ctrl+c", "q":
            return m, tea.Quit

		// pause
		case " ": 
			m.paused = !m.paused
			if m.status == NONE {
				m.status = FOCUS
			}

		// restart
		case "r":
			m.paused = true
			if m.status == FOCUS {
				m.remainingSeconds = m.focusSeconds

			}
			if m.status == CHILL {
				m.remainingSeconds = m.chillSeconds
			}

		// clear
		case "c":
			m.status = NONE
			m.remainingSeconds = -1
			m.focusSeconds = 0
			m.chillSeconds = 0
			m.paused = true


		// next
		case "n":
			if m.status == FOCUS {
				m.remainingSeconds = m.chillSeconds
				m.status = CHILL

			}
			if m.status == CHILL {
				m.remainingSeconds = m.focusSeconds
				m.status = FOCUS
			}
			m.paused = false

		// set
		case "s":

		// help
		case "?":

        }
    }

    return m, nil
}

func hotkeyBar() string {
	spacer := "        "
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		hotkeyHint("r", "restart"),
		spacer,
		hotkeyHint("spc", "toggle pause"),
		spacer,
		hotkeyHint("c", "clear"),
		spacer,
		hotkeyHint("s", "skip"),
		spacer,
		hotkeyHint("q", "quit"),
		spacer,
		hotkeyHint("?", "help"),
	)
}

func hotkeyHint(hotkey, text string) string {
	hotkeyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#A8C4FF")).AlignHorizontal(lipgloss.Right)

	return hotkeyStyle.Render(hotkey) + " " + text
}

func (m model) View() string {
	mainPaneStyle := lipgloss.NewStyle().Height(m.windowHeight - 5).Width(m.windowWidth - 2)
	hotkeysPaneStyle := lipgloss.NewStyle().Width(m.windowWidth - 2)

	return lipgloss.JoinVertical(
		lipgloss.Center,
		mainPaneStyle.Border(lipgloss.RoundedBorder(), true).Render(mainPaneStyle.Render("13:37")),
		hotkeysPaneStyle.Border(lipgloss.RoundedBorder(), true).Render(hotkeysPaneStyle.Render(hotkeyBar())),
	)
}

func main() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		log.Fatal("Error running program: ", err)
	}
}
