package main

import "github.com/gdamore/tcell/v2"

type GameMap struct {
	Width, Height int
	Tiles         []Tile
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{Width: width, Height: height, Tiles: make([]Tile, width*height)}
	for i := range m.Tiles {
		m.Tiles[i] = floor
	}
	return m
}

func (m *GameMap) Render(screen tcell.Screen) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			g := m.Tiles[y*m.Width+x].Dark
			screen.SetContent(x, y, g.Ch, nil, g.Style())
		}
	}
}
