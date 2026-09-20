package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Fighter struct {
	HP, MaxHP int
	Defense   int
	Power     int
}

func (f *Fighter) SetHP(engine *Engine, entity *Entity, hp int) {
	f.HP = max(0, min(hp, f.MaxHP))
	if f.HP == 0 && entity.Alive {
		entity.Die(engine)
	}
}

// Heal restores up to amount HP and returns how much was actually recovered. Your turn!
func (f *Fighter) Heal(amount int) int {
	return 0
}

func (e *Entity) Die(engine *Engine) {
	if e == engine.Player {
		engine.Log("You died!")
	} else {
		engine.Log(fmt.Sprintf("%s is dead!", e.Name))
	}
	e.Char = '%'
	e.Color = tcell.NewRGBColor(191, 0, 0)
	e.BlocksMovement = false
	e.Alive = false
	e.Name = "remains of " + e.Name
	e.RenderOrder = RenderCorpse
}
