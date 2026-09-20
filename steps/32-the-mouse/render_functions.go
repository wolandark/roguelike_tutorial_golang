package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell/v2"
)

func renderBar(screen tcell.Screen, x, y, current, maximum, totalWidth int) {
	filled := 0
	if maximum > 0 {
		filled = current * totalWidth / maximum
	}
	for i := 0; i < totalWidth; i++ {
		bg := colorBarEmpty
		if i < filled {
			bg = colorBarFilled
		}
		screen.SetContent(x+i, y, ' ', nil, tcell.StyleDefault.Background(bg))
	}
	label := fmt.Sprintf("HP: %d/%d", current, maximum)
	for i, r := range label {
		bg := colorBarEmpty
		if x+i < x+filled {
			bg = colorBarFilled
		}
		screen.SetContent(x+1+i, y, r, nil, tcell.StyleDefault.Foreground(colorBarText).Background(bg))
	}
}

func getNamesAt(m *GameMap, x, y int) string {
	if !m.IsVisible(x, y) {
		return ""
	}
	var names []string
	for _, e := range m.Entities {
		if e.X == x && e.Y == y {
			names = append(names, e.Name)
		}
	}
	return strings.Join(names, ", ")
}

func renderNamesAtMouse(screen tcell.Screen, x, y int, engine *Engine) {
	drawText(screen, x, y, getNamesAt(engine.GameMap, engine.MouseX, engine.MouseY), tcell.StyleDefault)
}
