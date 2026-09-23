package main

import "github.com/gdamore/tcell/v2"

type GameMap struct {
	Width, Height int
	Tiles         [][]Tile
	Visible       [][]bool
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{
		Width: width, Height: height,
		Tiles:   make([][]Tile, height),
		Visible: make([][]bool, height),
	}
	for y := range m.Tiles {
		m.Tiles[y] = make([]Tile, width)
		m.Visible[y] = make([]bool, width)
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

func (m *GameMap) ComputeFOV(ox, oy, radius int) {
	for _, row := range m.Visible {
		clear(row)
	}
	for y := oy - radius; y <= oy+radius; y++ {
		for x := ox - radius; x <= ox+radius; x++ {
			if m.InBounds(x, y) {
				m.Visible[y][x] = true
			}
		}
	}
}

func (m *GameMap) Render(screen tcell.Screen) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			g := shroud
			if m.IsVisible(x, y) {
				g = m.TileAt(x, y).Light
			}
			screen.SetContent(x, y, g.Ch, nil, g.Style())
		}
	}
}
