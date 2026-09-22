package main

import "github.com/gdamore/tcell/v2"

type Glyph struct {
	Ch     rune
	FG, BG tcell.Color
}

func (g Glyph) Style() tcell.Style {
	return tcell.StyleDefault.Foreground(g.FG).Background(g.BG)
}

var floorGlyph = Glyph{' ', tcell.ColorWhite, tcell.NewRGBColor(50, 50, 150)}
