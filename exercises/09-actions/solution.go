package main

import "github.com/gdamore/tcell/v2"

func handleKey(ev *tcell.EventKey) Action {
	switch {
	case ev.Key() == tcell.KeyUp || ev.Rune() == 'k':
		return MovementAction{DX: 0, DY: -1}
	case ev.Key() == tcell.KeyDown || ev.Rune() == 'j':
		return MovementAction{DX: 0, DY: 1}
	case ev.Key() == tcell.KeyLeft || ev.Rune() == 'h':
		return MovementAction{DX: -1, DY: 0}
	case ev.Key() == tcell.KeyRight || ev.Rune() == 'l':
		return MovementAction{DX: 1, DY: 0}
	case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:
		return EscapeAction{}
	case ev.Rune() == 'y':
		return MovementAction{DX: -1, DY: -1}
	case ev.Rune() == 'u':
		return MovementAction{DX: 1, DY: -1}
	case ev.Rune() == 'b':
		return MovementAction{DX: -1, DY: 1}
	case ev.Rune() == 'n':
		return MovementAction{DX: 1, DY: 1}
	}
	return nil
}
