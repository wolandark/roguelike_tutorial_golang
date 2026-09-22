package main

import (
	"sort"

	"github.com/gdamore/tcell/v2"
)

type GameMap struct {
	Width, Height int
	Tiles         []Tile
	Visible       []bool
	Explored      []bool
	Entities      []*Entity
}

func NewGameMap(width, height int) *GameMap {
	m := &GameMap{
		Width: width, Height: height,
		Tiles:    make([]Tile, width*height),
		Visible:  make([]bool, width*height),
		Explored: make([]bool, width*height),
	}
	for i := range m.Tiles {
		m.Tiles[i] = wall
	}
	return m
}

func (m *GameMap) InBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

func (m *GameMap) TileAt(x, y int) Tile { return m.Tiles[y*m.Width+x] }

func (m *GameMap) SetTile(x, y int, t Tile) { m.Tiles[y*m.Width+x] = t }

func (m *GameMap) IsVisible(x, y int) bool { return m.InBounds(x, y) && m.Visible[y*m.Width+x] }

func (m *GameMap) IsExplored(x, y int) bool { return m.InBounds(x, y) && m.Explored[y*m.Width+x] }

func (m *GameMap) setVisible(x, y int) {
	if m.InBounds(x, y) {
		m.Visible[y*m.Width+x] = true
		m.Explored[y*m.Width+x] = true
	}
}

func (m *GameMap) GetBlockingEntityAt(x, y int) *Entity {
	for _, e := range m.Entities {
		if e.BlocksMovement && e.X == x && e.Y == y {
			return e
		}
	}
	return nil
}

func (m *GameMap) EntityAt(x, y int) *Entity {
	for _, e := range m.Entities {
		if e.X == x && e.Y == y {
			return e
		}
	}
	return nil
}

func (m *GameMap) RemoveEntity(entity *Entity) {
	for i, e := range m.Entities {
		if e == entity {
			m.Entities = append(m.Entities[:i], m.Entities[i+1:]...)
			return
		}
	}
}

func (m *GameMap) GetActorAt(x, y int) *Entity {
	for _, e := range m.Entities {
		if e.Fighter != nil && e.Alive && e.X == x && e.Y == y {
			return e
		}
	}
	return nil
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
	sorted := append([]*Entity(nil), m.Entities...)
	sort.SliceStable(sorted, func(i, j int) bool { return sorted[i].RenderOrder < sorted[j].RenderOrder })
	for _, e := range sorted {
		if m.IsVisible(e.X, e.Y) {
			bg := m.TileAt(e.X, e.Y).Light.BG
			screen.SetContent(e.X, e.Y, e.Char, nil, tcell.StyleDefault.Foreground(e.Color).Background(bg))
		}
	}
}
