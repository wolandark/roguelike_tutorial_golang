package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Engine struct {
	Player         *Entity
	GameMap        *GameMap
	GameWorld      *GameWorld
	MessageLog     *MessageLog
	MouseX, MouseY int
}

func (e *Engine) Log(msg string, color tcell.Color) {
	e.MessageLog.AddMessage(msg, color, true)
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
	renderBar(screen, 0, 45, e.Player.Fighter.HP, e.Player.Fighter.MaxHP, 20)
	drawText(screen, 0, 47, fmt.Sprintf("Dungeon level: %d", e.GameWorld.CurrentFloor), tcell.StyleDefault)
	renderNamesAtMouse(screen, 21, 44, e)
	e.MessageLog.Render(screen, 21, 45, 40, 5)
}

func drawText(screen tcell.Screen, x, y int, text string, style tcell.Style) {
	for i, r := range []rune(text) {
		screen.SetContent(x+i, y, r, nil, style)
	}
}
