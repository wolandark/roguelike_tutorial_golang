package main

import "fmt"

type Action interface {
	Perform(engine *Engine, entity *Entity)
}

type EscapeAction struct{}

func (EscapeAction) Perform(*Engine, *Entity) {}

type ActionWithDirection struct {
	DX, DY int
}

func (a ActionWithDirection) Dest(entity *Entity) (int, int) {
	return entity.X + a.DX, entity.Y + a.DY
}

func (a ActionWithDirection) BlockingEntity(engine *Engine, entity *Entity) *Entity {
	x, y := a.Dest(entity)
	return engine.GameMap.GetBlockingEntityAt(x, y)
}

type MeleeAction struct{ ActionWithDirection }

func (a MeleeAction) Perform(engine *Engine, entity *Entity) {
	target := a.BlockingEntity(engine, entity)
	if target == nil {
		return
	}
	engine.Log(fmt.Sprintf("You kick the %s, much to its annoyance!", target.Name))
}

type MovementAction struct{ ActionWithDirection }

func (a MovementAction) Perform(engine *Engine, entity *Entity) {
	destX, destY := a.Dest(entity)
	if !engine.GameMap.InBounds(destX, destY) {
		return
	}
	if !engine.GameMap.TileAt(destX, destY).Walkable {
		return
	}
	if engine.GameMap.GetBlockingEntityAt(destX, destY) != nil {
		return
	}
	entity.Move(a.DX, a.DY)
}

type BumpAction struct{ ActionWithDirection }

func (a BumpAction) Perform(engine *Engine, entity *Entity) {
	if a.BlockingEntity(engine, entity) != nil {
		MeleeAction{a.ActionWithDirection}.Perform(engine, entity)
	} else {
		MovementAction{a.ActionWithDirection}.Perform(engine, entity)
	}
}
