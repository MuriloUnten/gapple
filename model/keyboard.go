package model

import (
	tea "github.com/charmbracelet/bubbletea"
)

type HotkeyManager struct {
	hotkeys map[string]func(Model) (Model, tea.Cmd)
}

func NewHotkeyManager() *HotkeyManager {
	return &HotkeyManager{
		hotkeys: make(map[string]func(Model) (Model, tea.Cmd), 7),
	}
}

func (hm *HotkeyManager) LoadHotkeys() {
	hm.hotkeys["ctrl+c"] = Model.QuitAction
	hm.hotkeys["q"] = Model.QuitAction
	hm.hotkeys["p"] = Model.PauseAction
	hm.hotkeys["r"] = Model.RestartAction
	hm.hotkeys["c"] = Model.ClearAction
	hm.hotkeys["n"] = Model.NextAction
	hm.hotkeys["s"] = Model.SetAction
	hm.hotkeys["h"] = Model.HelpAction
}

func (hm *HotkeyManager) GetAction(hotkey string) func(Model) (Model, tea.Cmd) {
	fn, ok := hm.hotkeys["string"]
	if !ok {
		return Model.UndefinedAction
	}

	return fn
}
