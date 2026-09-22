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

func (m *GameMap) TileAt(x, y int) Tile { return m.Tiles[y*m.Width+x] }

func (m *GameMap) SetTile(x, y int, t Tile) { m.Tiles[y*m.Width+x] = t }

func (m *GameMap) Render(screen tcell.Screen) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			g := m.TileAt(x, y).Dark
			screen.SetContent(x, y, g.Ch, nil, g.Style())
		}
	}
}
