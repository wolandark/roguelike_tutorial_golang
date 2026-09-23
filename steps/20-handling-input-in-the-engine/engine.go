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
	switch action := handleKey(key).(type) {
	case EscapeAction:
		return true
	case MovementAction:
		destX, destY := e.Player.X+action.DX, e.Player.Y+action.DY
		if e.GameMap.InBounds(destX, destY) && e.GameMap.TileAt(destX, destY).Walkable {
			e.Player.Move(action.DX, action.DY)
		}
	}
	return false
}

func (e *Engine) Render(screen tcell.Screen) {
	screen.Clear()
	e.GameMap.Render(screen)
	for _, ent := range e.Entities {
		screen.SetContent(ent.X, ent.Y, ent.Char, nil, tcell.StyleDefault.Foreground(ent.Color))
	}
	screen.Show()
}
