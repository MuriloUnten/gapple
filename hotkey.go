package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Bind struct {
	action func(Model) (Model, tea.Cmd)
	hotkeyText string
	shortDescription string
	longDescription string
}

var hotkeys map[string]Bind

func Hotkeys() map[string]Bind {
	return hotkeys
}

func LoadHotkeys() {
	hotkeys = make(map[string]Bind, 7) 

	// TODO try reading from config file first, if it doesnt exist OR contains error, use default hotkeys

	hotkeys = loadDefaults()
}

// TODO Action && descriptions should be a predefined set packed together, only the hotkeyText should be changeable
func loadDefaults() map[string]Bind {

	hotkeys["ctrl+c"] = Bind{
		action: Model.QuitAction,
		hotkeyText: "ctrl+c",
		shortDescription: "quit",
		longDescription: "Close program",
	}

	hotkeys["q"] = Bind{
		action: Model.QuitAction,
		hotkeyText: "q",
		shortDescription: "quit",
		longDescription: "Close program",
	}

	hotkeys[" "] = Bind{
		action: Model.PauseAction,
		hotkeyText: "spc",
		shortDescription: "pause",
		longDescription: "Pause timer",
	}

	hotkeys["r"] = Bind{
		action: Model.RestartAction,
		hotkeyText: "r",
		shortDescription: "restart",
		longDescription: "Restart timer",
	}

	hotkeys["c"] = Bind{
		action: Model.ClearAction,
		hotkeyText: "c",
		shortDescription: "clear",
		longDescription: "Clear timers",
	}

	hotkeys["n"] = Bind{
		action: Model.NextAction,
		hotkeyText: "n",
		shortDescription: "next",
		longDescription: "Advance to next timer",
	}

	hotkeys["s"] = Bind{
		action: Model.SetAction,
		hotkeyText: "s",
		shortDescription: "set",
		longDescription: "Set timers",
	}

	hotkeys["?"] = Bind{
		action: Model.HelpAction,
		hotkeyText: "?",
		shortDescription: "help",
		longDescription: "Toggle help menu",
	}

	return hotkeys
}

func getBind(hotkey string) Bind {
	bind, ok := hotkeys[hotkey]

	// hotkey not registered
	if !ok {
		return Bind{
			action: Model.UndefinedAction,
		}
	}

	return bind
}
