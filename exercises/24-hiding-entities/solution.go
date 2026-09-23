package main

import "github.com/gdamore/tcell/v2"

type Engine struct {
	Entities []*Entity
	Player   *Entity
	GameMap  *GameMap
}

func (e *Engine) HandleEvent(ev tcell.Event) (quit bool) {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}
	if key.Key() == tcell.KeyRune && key.Rune() == 'm' {
		for _, row := range e.GameMap.Explored {
			for x := range row {
				row[x] = true
			}
		}
		return false
	}
	action := handleKey(key)
	if action == nil {
		return false
	}
	if _, ok := action.(EscapeAction); ok {
		return true
	}
	action.Perform(e, e.Player)
	e.UpdateFOV()
	return false
}

func (e *Engine) UpdateFOV() {
	e.GameMap.ComputeFOV(e.Player.X, e.Player.Y, 8)
}

func (e *Engine) Render(screen tcell.Screen) {
	screen.Clear()
	e.GameMap.Render(screen)
	for _, ent := range e.Entities {
		if e.GameMap.IsVisible(ent.X, ent.Y) {
			screen.SetContent(ent.X, ent.Y, ent.Char, nil, tcell.StyleDefault.Foreground(ent.Color))
		}
	}
	screen.Show()
}
