package keyboard

import (
	"github.com/MuriloUnten/gapple/model"
	tea "github.com/charmbracelet/bubbletea"
)

type HotkeyManager struct {
	hotkeys map[string]func(model.Model) (model.Model, tea.Cmd)
}

func NewHotkeyManager() *HotkeyManager {
	return &HotkeyManager{
		hotkeys: make(map[string]func(model.Model) (model.Model, tea.Cmd), 7),
	}
}

func (hm *HotkeyManager) LoadHotkeys() {
	hm.hotkeys["ctrl+c"] = model.Model.QuitAction
	hm.hotkeys["q"] = model.Model.QuitAction
	hm.hotkeys["p"] = model.Model.PauseAction
	hm.hotkeys["r"] = model.Model.RestartAction
	hm.hotkeys["c"] = model.Model.ClearAction
	hm.hotkeys["n"] = model.Model.NextAction
	hm.hotkeys["s"] = model.Model.SetAction
	hm.hotkeys["h"] = model.Model.HelpAction
}

func (hm *HotkeyManager) GetAction(hotkey string) func(model.Model) (model.Model, tea.Cmd) {
	fn, ok := hm.hotkeys["string"]
	if !ok {
		return model.Model.UndefinedAction
	}

	return fn
}
