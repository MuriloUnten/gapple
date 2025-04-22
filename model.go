package main

import (
	ct "github.com/MuriloUnten/gapple/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"log"
	"time"
)

type Mode int

const (
	TIMER Mode = iota
	HELP
	SET
)

var Hints = ""
var Numbers = numbers()
var Specials = specialCharacters()

type Model struct {
	status       Status
	timer        *ct.CountdownTimer
	focusSeconds int
	chillSeconds int
	windowWidth  int
	windowHeight int
	tick         int
	mode         Mode
}

func initialModel() Model {
	timer, err := ct.NewCountdownTimer(25 * 60) // TODO gonna have to change this aswell
	if err != nil {
		log.Fatal(err)
	}

	return Model{
		status:       NONE,
		timer:        timer,
		focusSeconds: 25 * 60, // TODO: review this. The initial times are hard coded
		chillSeconds: 5 * 60,
		windowWidth:  -1,
		windowHeight: -1,
		tick:         -1,
		mode:         TIMER,
	}
}

type TickMsg time.Time

func tickEveryFithSecond() tea.Cmd {
	return tea.Every(time.Millisecond*200, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return tickEveryFithSecond()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width

	case TickMsg:
		m.tick = time.Time(msg).Second()
		m.timer.Update(time.Time(msg))
		return m, tickEveryFithSecond()

	case tea.KeyMsg:
		return getBind(msg.String()).action(m)
	}

	return m, nil
}

func hotkeyBar() string {
	spacer := "        "

	keybinds := make([]string, 0)
	firstIteration := true
	for _, bind := range Hotkeys() {
		if !firstIteration {
			keybinds = append(keybinds, spacer)
		}
		hint := hotkeyHint(bind.hotkeyText, bind.shortDescription)
		keybinds = append(keybinds, hint)
		firstIteration = false
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		keybinds...,
	)
}

func hotkeyHint(hotkey, text string) string {
	hotkeyStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#A8C4FF")).
		AlignHorizontal(lipgloss.Right)

	return hotkeyStyle.Render(hotkey) + " " + text
}

func (m Model) View() string {

	return lipgloss.JoinVertical(
		lipgloss.Center,
		mainPane(m),
		hotkeyPane(m),
	)
}

func timerCharacters(minutes, seconds string) []string {
	characters := make([]string, 8)
	for _, c := range minutes {
		characters = append(characters, Numbers[string(c)])
		characters = append(characters, Specials[" "])
	}
	characters = append(characters, Specials[":"])
	for _, c := range seconds {
		characters = append(characters, Specials[" "])
		characters = append(characters, Numbers[string(c)])
	}

	return characters
}

func pauseStatusIcon(paused bool) string {
	if paused {
		return Specials["u"]
	} else {
		return Specials["p"]
	}
}

func statusText(status Status) string {
	if status == FOCUS {
		return "Deep Focus"
	} else if status == CHILL {
		return "Chill"
	}

	return ""
}

func mainPane(m Model) string {
	mainPaneStyle := mainPaneStyle(m)
	mainTimerStyle := mainTimerStyle()
	timersInfoStyle := timersInfoStyle()
	statusTextStyle := statusTextStyle()

	// Load status text
	statusText := statusText(m.status)

	// Load timer characters
	minutesString, secondsString := remainingTimeToString(m.timer.RemainingTime())
	characters := timerCharacters(minutesString, secondsString)

	// Load pause icon
	pauseIconStyle := lipgloss.NewStyle().MarginTop(2).Foreground(lipgloss.Color("#A8C4FF"))
	pauseIcon := pauseStatusIcon(m.timer.Paused())

	switch m.mode {
	case TIMER:
		return mainPaneStyle.Render(
			lipgloss.JoinVertical(
				lipgloss.Center,
				statusTextStyle.Render(statusText),
				mainTimerStyle.Render(
					lipgloss.JoinHorizontal(
						lipgloss.Center,
						characters...,
					),
				),
				pauseIconStyle.Render(pauseIcon),
				timersInfoStyle.Render(secondsToTimeString(m.focusSeconds)),
				timersInfoStyle.Render(secondsToTimeString(m.chillSeconds)),
			),
		)
	case SET:
	case HELP:
	}
	return ""
}

func hotkeyPane(m Model) string {
	hotkeysPaneStyle := hotkeysPaneStyle(m)
	return hotkeysPaneStyle.Render(Hints)
}

func RunProgram() {
	LoadHotkeys()
	Hints = hotkeyBar()
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		log.Fatal("Error running program: ", err)
	}
}
