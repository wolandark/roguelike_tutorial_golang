package main

import "fmt"

type Action interface {
	Perform(engine *Engine, entity *Entity)
}

type EscapeAction struct{}

func (EscapeAction) Perform(*Engine, *Entity) {}

type WaitAction struct{}

func (WaitAction) Perform(*Engine, *Entity) {}

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

func (a ActionWithDirection) TargetActor(engine *Engine, entity *Entity) *Entity {
	x, y := a.Dest(entity)
	return engine.GameMap.GetActorAt(x, y)
}

type MeleeAction struct{ ActionWithDirection }

func (a MeleeAction) Perform(engine *Engine, entity *Entity) {
	target := a.TargetActor(engine, entity)
	if target == nil {
		return
	}
	damage := entity.Fighter.Power - target.Fighter.Defense
	desc := fmt.Sprintf("%s attacks %s", entity.Name, target.Name)
	color := colorEnemyAtk
	if entity == engine.Player {
		color = colorPlayerAtk
	}
	if damage > 0 {
		engine.Log(fmt.Sprintf("%s for %d hit points.", desc, damage), color)
		target.Fighter.SetHP(engine, target, target.Fighter.HP-damage)
	} else {
		engine.Log(fmt.Sprintf("%s but does no damage.", desc), color)
	}
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
	if a.TargetActor(engine, entity) != nil {
		MeleeAction{a.ActionWithDirection}.Perform(engine, entity)
	} else {
		MovementAction{a.ActionWithDirection}.Perform(engine, entity)
	}
}
