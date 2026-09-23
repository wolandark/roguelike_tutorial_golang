package main

import "github.com/gdamore/tcell/v2"

type Glyph struct {
	Ch     rune
	FG, BG tcell.Color
}

func (g Glyph) Style() tcell.Style {
	return tcell.StyleDefault.Foreground(g.FG).Background(g.BG)
}

type Tile struct {
	Walkable    bool
	Transparent bool
	Dark        Glyph
}

var shroud = Glyph{' ', tcell.ColorWhite, tcell.ColorBlack}

var (
	floor = Tile{
		Walkable: true, Transparent: true,
		Dark: Glyph{' ', tcell.ColorWhite, tcell.NewRGBColor(50, 50, 150)},
	}
	wall = Tile{
		Walkable: false, Transparent: false,
		Dark: Glyph{' ', tcell.ColorWhite, tcell.NewRGBColor(0, 0, 100)},
	}
)
