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
		path := FindPath(engine.GameMap, entity.X, entity.Y, target.X, target.Y)
		if len(path) > 0 {
			next := path[0]
			MovementAction{ActionWithDirection{next[0] - entity.X, next[1] - entity.Y}}.Perform(engine, entity)
			return
		}
	}
	WaitAction{}.Perform(engine, entity)
}
