package main

type AI interface {
	Perform(engine *Engine, entity *Entity)
}

type HostileEnemy struct{}

func (HostileEnemy) Perform(engine *Engine, entity *Entity) {
	target := engine.Player
	dx, dy := target.X-entity.X, target.Y-entity.Y

	if engine.GameMap.IsVisible(entity.X, entity.Y) {
		if entity.Distance(target.X, target.Y) <= 1 {
			MeleeAction{ActionWithDirection{dx, dy}}.Perform(engine, entity)
			return
		}
		MovementAction{ActionWithDirection{sign(dx), sign(dy)}}.Perform(engine, entity)
		return
	}
	WaitAction{}.Perform(engine, entity)
}
