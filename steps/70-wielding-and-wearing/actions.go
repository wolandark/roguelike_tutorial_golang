package main

import "fmt"

type Action interface {
	Perform(engine *Engine, entity *Entity) error
}

type EscapeAction struct{}

func (EscapeAction) Perform(*Engine, *Entity) error { return nil }

type WaitAction struct{}

func (WaitAction) Perform(*Engine, *Entity) error { return nil }

type TakeStairsAction struct{}

func (TakeStairsAction) Perform(engine *Engine, entity *Entity) error {
	if entity.X != engine.GameMap.DownstairsX || entity.Y != engine.GameMap.DownstairsY {
		return Impossible{"There are no stairs here."}
	}
	engine.GameWorld.GenerateFloor(engine)
	engine.Log("You descend the staircase.", colorDescend)
	return nil
}

type PickupAction struct{}

func (PickupAction) Perform(engine *Engine, entity *Entity) error {
	inventory := entity.Inventory
	for _, item := range engine.GameMap.Entities {
		if !item.IsItem() || item.X != entity.X || item.Y != entity.Y {
			continue
		}
		if len(inventory.Items) >= inventory.Capacity {
			return Impossible{"Your inventory is full."}
		}
		engine.GameMap.RemoveEntity(item)
		inventory.Items = append(inventory.Items, item)
		engine.Log(fmt.Sprintf("You picked up the %s!", item.Name), colorWhite)
		return nil
	}
	return Impossible{"There is nothing here to pick up."}
}

type EquipAction struct {
	Item *Entity
}

func (a EquipAction) Perform(engine *Engine, entity *Entity) error {
	entity.ToggleEquip(engine, a.Item, true)
	return nil
}

type DropItem struct {
	Item *Entity
}

func (a DropItem) Perform(engine *Engine, entity *Entity) error {
	entity.Inventory.Drop(engine, entity, a.Item)
	return nil
}

type ItemAction struct {
	Item             *Entity
	TargetX, TargetY int
}

func (a ItemAction) TargetActor(engine *Engine) *Entity {
	return engine.GameMap.GetActorAt(a.TargetX, a.TargetY)
}

func (a ItemAction) Perform(engine *Engine, entity *Entity) error {
	return a.Item.Consumable.Activate(engine, a, entity)
}

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

func (a MeleeAction) Perform(engine *Engine, entity *Entity) error {
	target := a.TargetActor(engine, entity)
	if target == nil {
		return Impossible{"Nothing to attack."}
	}
	damage := entity.Power() - target.Defense()
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
	return nil
}

type MovementAction struct{ ActionWithDirection }

func (a MovementAction) Perform(engine *Engine, entity *Entity) error {
	destX, destY := a.Dest(entity)
	if !engine.GameMap.InBounds(destX, destY) || !engine.GameMap.TileAt(destX, destY).Walkable {
		return Impossible{"That way is blocked."}
	}
	if engine.GameMap.GetBlockingEntityAt(destX, destY) != nil {
		return Impossible{"That way is blocked."}
	}
	entity.Move(a.DX, a.DY)
	return nil
}

type BumpAction struct{ ActionWithDirection }

func (a BumpAction) Perform(engine *Engine, entity *Entity) error {
	if a.TargetActor(engine, entity) != nil {
		return MeleeAction{a.ActionWithDirection}.Perform(engine, entity)
	}
	return MovementAction{a.ActionWithDirection}.Perform(engine, entity)
}
