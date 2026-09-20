package main

import "github.com/gdamore/tcell/v2"

var moveKeys = map[tcell.Key][2]int{
	tcell.KeyUp:    {0, -1},
	tcell.KeyDown:  {0, 1},
	tcell.KeyLeft:  {-1, 0},
	tcell.KeyRight: {1, 0},
	tcell.KeyHome:  {-1, -1},
	tcell.KeyEnd:   {-1, 1},
	tcell.KeyPgUp:  {1, -1},
	tcell.KeyPgDn:  {1, 1},
}

var moveRunes = map[rune][2]int{
	'h': {-1, 0}, 'j': {0, 1}, 'k': {0, -1}, 'l': {1, 0},
	'y': {-1, -1}, 'u': {1, -1}, 'b': {-1, 1}, 'n': {1, 1},
	'1': {-1, 1}, '2': {0, 1}, '3': {1, 1}, '4': {-1, 0},
	'6': {1, 0}, '7': {-1, -1}, '8': {0, -1}, '9': {1, -1},
}

func handleKey(ev *tcell.EventKey) Action {
	if d, ok := moveKeys[ev.Key()]; ok {
		return BumpAction{ActionWithDirection{d[0], d[1]}}
	}
	switch ev.Key() {
	case tcell.KeyEscape, tcell.KeyCtrlC:
		return EscapeAction{}
	case tcell.KeyRune:
		if d, ok := moveRunes[ev.Rune()]; ok {
			return BumpAction{ActionWithDirection{d[0], d[1]}}
		}
		if ev.Rune() == '.' || ev.Rune() == '5' {
			return WaitAction{}
		}
	}
	return nil
}
