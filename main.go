package main

import (
	"log"
	"strconv"
	"time"

	ct "github.com/MuriloUnten/gapple/timer"
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
	timer *ct.CountdownTimer
	focusSeconds int
	chillSeconds int
	windowWidth int
	windowHeight int
}

func initialModel() model {
	timer, err := ct.NewCountdownTimer(0)
	if err != nil {
		log.Fatal(err)
	}

	return model {
		status: NONE,
		timer: timer,
		focusSeconds: 25 * 60, // TODO: review this. The initial times are hard coded
		chillSeconds: 5 * 60,
		windowWidth: -1,
		windowHeight: -1,
	}
}

type TickMsg time.Time

func tickEverySecond() tea.Cmd {
    return tea.Every(time.Second, func(t time.Time) tea.Msg {
        return TickMsg(t)
    })
}

func (m model) Init() tea.Cmd {
	return tickEverySecond()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width

	case TickMsg:
		if m.timer.Paused() {
			break
		}
		
		m.timer.Update(time.Time(msg))
		return m, tickEverySecond()

    case tea.KeyMsg:
        switch msg.String() {
		// quit
        case "ctrl+c", "q":
            return m, tea.Quit

		// pause
		case " ": 
			m.timer.TogglePause()
			if m.status == NONE {
				m.status = FOCUS
			}

		// restart
		case "r":
			m.timer.Reset()

		// clear
		case "c":
			m.status = NONE
			m.focusSeconds = 0
			m.chillSeconds = 0
			m.timer.Preset(0)


		// next
		case "n":
			if m.status == FOCUS {
				m.timer.Preset(m.chillSeconds)
				m.status = CHILL

			} else if m.status == CHILL {
				m.timer.Preset(m.focusSeconds)
				m.status = FOCUS
			}
			m.timer.Unpause()

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
		hotkeyHint("s", "set"),
		spacer,
		hotkeyHint("q", "quit"),
		spacer,
		hotkeyHint("n", "next"),
		spacer,
		hotkeyHint("?", "help"),
	)
}

func remainingTimeToString(rt time.Duration) (string, string) {
	minutes := strconv.Itoa(int(rt.Seconds() / 60))
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}

	seconds := strconv.Itoa(int(rt.Seconds()) % 60)
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}

	return minutes, seconds
}

func secondsToTimeString(s int) string {
	str := ""
	minutes := strconv.Itoa(s / 60)
	if len(minutes) == 1 {
		minutes = "0" + minutes
	}
	str += minutes
	str += ":"
	seconds := strconv.Itoa(s % 60)
	if len(seconds) == 1 {
		seconds = "0" + seconds
	}
	str += seconds

	return str
}

func numbers() map[string]string {
	m := make(map[string]string)

	m["0"] = "█████\n" +
			 "█   █\n" +
			 "█   █\n" +
			 "█   █\n" +
			 "█████"

	m["1"] = "    █\n" +
			 "    █\n" +
			 "    █\n" +
			 "    █\n" +
			 "    █"

	m["2"] = "█████\n" +
			 "    █\n" +
			 "█████ \n" +
			 "█    \n" +
			 "█████"

	m["3"] = "█████\n" +
			 "    █\n" +
			 "█████\n" +
			 "    █\n" +
			 "█████"

	m["4"] = "█   █\n" +
			 "█   █\n" +
			 "█████\n" +
			 "    █\n" +
			 "    █"

	m["5"] = "█████\n" +
			 "█    \n" +
			 "█████\n" +
			 "    █\n" +
			 "█████"

	m["6"] = "█████\n" +
			 "█    \n" +
			 "█████\n" +
			 "█   █\n" +
			 "█████"

	m["7"] = "█████\n" +
			 "    █\n" +
			 "    █\n" +
			 "    █\n" +
			 "    █"

	m["8"] = "█████\n" +
			 "█   █\n" +
			 "█████\n" +
			 "█   █\n" +
			 "█████"

	m["9"] = "█████\n" +
			 "█   █\n" +
			 "█████\n" +
			 "    █\n" +
			 "█████"

	m[":"] = "     \n" +
			 "  █  \n" +
			 "     \n" +
			 "  █  \n" +
			 "     "

	m[" "] = " \n" +
			 " \n" +
			 " \n" +
			 " \n" +
			 " "
	return m
}

func hotkeyHint(hotkey, text string) string {
	hotkeyStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#A8C4FF")).
		AlignHorizontal(lipgloss.Right)

	return hotkeyStyle.Render(hotkey) + " " + text
}

func (m model) View() string {
	n := numbers()

	mainPaneStyle := lipgloss.NewStyle().
		Height(m.windowHeight - 5).
		Width(m.windowWidth - 2).
		Align(lipgloss.Center, lipgloss.Center).
		Border(lipgloss.RoundedBorder(), true)

	hotkeysPaneStyle := lipgloss.NewStyle().
		Width(m.windowWidth - 2).
		Align(lipgloss.Center, lipgloss.Center)

	mainTimerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#A8C4FF"))
	timersInfoStyle := lipgloss.NewStyle().MarginTop(1)

	statusTextStyle := lipgloss.NewStyle().Margin(1)
	statusText := ""
	if m.status == FOCUS {
		statusText = "Deep Focus"
	} else if m.status == CHILL {
		statusText = "Chill"
	}

	minutesString, secondsString := remainingTimeToString(m.timer.RemainingTime())
	characters := make([]string, 8)
	for _, c := range minutesString {
		characters = append(characters, n[string(c)])
		characters = append(characters, n[" "])
	}
	characters = append(characters, n[":"])
	for _, c := range secondsString {
		characters = append(characters, n[" "])
		characters = append(characters, n[string(c)])
	}

	return lipgloss.JoinVertical(
		lipgloss.Center,
		mainPaneStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				statusTextStyle.Render(statusText),
				mainTimerStyle.Render(
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						characters...
					),
				),
				timersInfoStyle.Render(secondsToTimeString(m.focusSeconds)),
				timersInfoStyle.Render(secondsToTimeString(m.chillSeconds)),
			),
		),
		hotkeysPaneStyle.Border(lipgloss.RoundedBorder(), true).Render(hotkeyBar()),
	)
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		log.Fatal("Error running program: ", err)
	}
}
