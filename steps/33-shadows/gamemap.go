package main

import "github.com/gdamore/tcell/v2"

type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
	Visible       [][]bool
	Explored      [][]bool
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{
		Width: width, Height: height,
		Tiles:    make([][]Tile, height),
		Visible:  make([][]bool, height),
		Explored: make([][]bool, height),
	}
	for y := range m.Tiles {
		m.Tiles[y] = make([]Tile, width)
		m.Visible[y] = make([]bool, width)
		m.Explored[y] = make([]bool, width)
		for x := range m.Tiles[y] {
			m.Tiles[y][x] = wall
		}
	}
	return m
}

func (m *GameMap) InBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

func (m *GameMap) TileAt(x, y int) Tile { return m.Tiles[y][x] }

func (m *GameMap) SetTile(x, y int, t Tile) { m.Tiles[y][x] = t }

func (m *GameMap) IsVisible(x, y int) bool { return m.InBounds(x, y) && m.Visible[y][x] }

func (m *GameMap) IsExplored(x, y int) bool { return m.InBounds(x, y) && m.Explored[y][x] }

func (m *GameMap) setVisible(x, y int) {
	if m.InBounds(x, y) {
		m.Visible[y][x] = true
		m.Explored[y][x] = true
	}
}

func (m *GameMap) Render(screen tcell.Screen) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			g := shroud
			switch {
			case m.IsVisible(x, y):
				g = m.TileAt(x, y).Light
			case m.IsExplored(x, y):
				g = m.TileAt(x, y).Dark
			}
			screen.SetContent(x, y, g.Ch, nil, g.Style())
		}
	}
}
