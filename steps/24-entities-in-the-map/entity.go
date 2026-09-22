package main

import "github.com/gdamore/tcell/v2"

type Entity struct {
	X, Y           int
	Char           rune
	Color          tcell.Color
	Name           string
	BlocksMovement bool
}

func (e Entity) Spawn(m *GameMap, x, y int) *Entity {
	clone := e
	clone.X, clone.Y = x, y
	m.Entities = append(m.Entities, &clone)
	return &clone
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}
