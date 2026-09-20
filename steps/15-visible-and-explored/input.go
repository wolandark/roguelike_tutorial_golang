package main

import "github.com/gdamore/tcell/v2"

func handleKey(ev *tcell.EventKey) Action {
	switch ev.Key() {
	case tcell.KeyUp:
		return MovementAction{DX: 0, DY: -1}
	case tcell.KeyDown:
		return MovementAction{DX: 0, DY: 1}
	case tcell.KeyLeft:
		return MovementAction{DX: -1, DY: 0}
	case tcell.KeyRight:
		return MovementAction{DX: 1, DY: 0}
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return EscapeAction{}
	}
	return nil
}
