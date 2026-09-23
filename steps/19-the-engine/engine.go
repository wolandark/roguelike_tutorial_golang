package main

import "github.com/gdamore/tcell/v2"

type Engine struct {
	Entities []*Entity
	Player   *Entity
	GameMap  *GameMap
}

func (e *Engine) Render(screen tcell.Screen) {
	screen.Clear()
	e.GameMap.Render(screen)
	for _, ent := range e.Entities {
		screen.SetContent(ent.X, ent.Y, ent.Char, nil, tcell.StyleDefault.Foreground(ent.Color))
	}
	screen.Show()
}
