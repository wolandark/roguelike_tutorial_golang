package main

import "github.com/gdamore/tcell/v2"

func handleKey(ev *tcell.EventKey) Action {
	switch ev.Key() {
	case tcell.KeyUp:
		return BumpAction{ActionWithDirection{0, -1}}
	case tcell.KeyDown:
		return BumpAction{ActionWithDirection{0, 1}}
	case tcell.KeyLeft:
		return BumpAction{ActionWithDirection{-1, 0}}
	case tcell.KeyRight:
		return BumpAction{ActionWithDirection{1, 0}}
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return EscapeAction{}
	}
	return nil
}
