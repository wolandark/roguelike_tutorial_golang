package main

import "github.com/gdamore/tcell/v2"

func handleKey(ev *tcell.EventKey) Action {
	switch {
	case ev.Key() == tcell.KeyUp || ev.Rune() == 'k':
		return BumpAction{ActionWithDirection{0, -1}}
	case ev.Key() == tcell.KeyDown || ev.Rune() == 'j':
		return BumpAction{ActionWithDirection{0, 1}}
	case ev.Key() == tcell.KeyLeft || ev.Rune() == 'h':
		return BumpAction{ActionWithDirection{-1, 0}}
	case ev.Key() == tcell.KeyRight || ev.Rune() == 'l':
		return BumpAction{ActionWithDirection{1, 0}}
	case ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC:
		return EscapeAction{}
	}
	return nil
}
