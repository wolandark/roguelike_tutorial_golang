package main

import "github.com/gdamore/tcell/v2"

type Entity struct {
	X, Y  int
	Char  rune
	Color tcell.Color
}

func (e *Entity) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}
