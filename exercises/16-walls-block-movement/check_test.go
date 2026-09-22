package main

import "testing"

func TestBorder(t *testing.T) {
	m := NewGameMap(12, 7)
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			border := x == 0 || y == 0 || x == m.Width-1 || y == m.Height-1
			if border && m.TileAt(x, y).Walkable {
				t.Fatalf("tile (%d,%d) is on the border but is not a wall", x, y)
			}
			if !border && !m.TileAt(x, y).Walkable {
				t.Fatalf("tile (%d,%d) is inside but is not floor", x, y)
			}
		}
	}
}
