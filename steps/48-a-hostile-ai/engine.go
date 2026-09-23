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

func (e *Engine) HandleEvent(ev tcell.Event) (quit bool) {
	key, ok := ev.(*tcell.EventKey)
	if !ok {
		return false
	}
	action := handleKey(key)
	if action == nil {
		return false
	}
	if _, ok := action.(EscapeAction); ok {
		return true
	}
	action.Perform(e, e.Player)
	e.HandleEnemyTurns()
	e.UpdateFOV()
	return false
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
	screen.Clear()
	e.GameMap.Render(screen)
	hp := fmt.Sprintf("HP: %d/%d", e.Player.Fighter.HP, e.Player.Fighter.MaxHP)
	for j, r := range []rune(hp) {
		screen.SetContent(1+j, 46, r, nil, tcell.StyleDefault)
	}
	for i, msg := range e.Messages {
		for j, r := range []rune(msg) {
			screen.SetContent(21+j, e.GameMap.Height+i, r, nil, tcell.StyleDefault)
		}
	}
	screen.Show()
}
