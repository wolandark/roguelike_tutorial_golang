package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Engine struct {
	Player   *Entity
	GameMap  *GameMap
	Messages []string
}

func (e *Engine) Log(msg string) {
	e.Messages = append(e.Messages, msg)
	if len(e.Messages) > 4 {
		e.Messages = e.Messages[1:]
	}
}

func (e *Engine) HandleEnemyTurns() {
	for _, ent := range append([]*Entity(nil), e.GameMap.Entities...) {
		if ent != e.Player && ent.Alive && ent.AI != nil {
			ent.AI.Perform(e, ent)
		}
	}
}

func (e *Engine) UpdateFOV() {
	e.GameMap.ComputeFOV(e.Player.X, e.Player.Y, 8)
}

func (e *Engine) Render(screen tcell.Screen) {
	e.GameMap.Render(screen)
	drawText(screen, 1, 46, fmt.Sprintf("HP: %d/%d", e.Player.Fighter.HP, e.Player.Fighter.MaxHP), tcell.StyleDefault)
	for i, msg := range e.Messages {
		drawText(screen, 21, e.GameMap.Height+i, msg, tcell.StyleDefault)
	}
}

func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range []rune(text) {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
