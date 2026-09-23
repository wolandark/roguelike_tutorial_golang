package main

import "github.com/gdamore/tcell/v2"

type Engine struct {
	Player  *Entity
	GameMap *GameMap
}

func (e *Engine) HandleEvent(ev tcell.Event) (quit bool) {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
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
	screen.Show()
}
