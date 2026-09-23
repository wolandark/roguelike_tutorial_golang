package main

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
