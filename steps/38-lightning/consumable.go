package main

import "fmt"

type Consumable interface {
	GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler)
	Activate(engine *Engine, action ItemAction, consumer *Entity) error
}

func consume(consumer, item *Entity) {
	if consumer.Inventory != nil {
		consumer.Inventory.Remove(item)
	}
}

type HealingConsumable struct {
	Amount int
}

func (HealingConsumable) GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler) {
	return ItemAction{Item: item}, nil
}

func (h HealingConsumable) Activate(engine *Engine, action ItemAction, consumer *Entity) error {
	recovered := consumer.Fighter.Heal(h.Amount)
	if recovered == 0 {
		return Impossible{"Your health is already full."}
	}
	engine.Log(fmt.Sprintf("You consume the %s, and recover %d HP!", action.Item.Name, recovered), colorHealthRecovered)
	consume(consumer, action.Item)
	return nil
}

type LightningDamageConsumable struct {
	Damage       int
	MaximumRange int
}

func (LightningDamageConsumable) GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler) {
	return ItemAction{Item: item}, nil
}

func (l LightningDamageConsumable) Activate(engine *Engine, action ItemAction, consumer *Entity) error {
	var target *Entity
	closest := l.MaximumRange + 1
	for _, actor := range engine.GameMap.Actors() {
		if actor == consumer || !engine.GameMap.IsVisible(actor.X, actor.Y) {
			continue
		}
		if d := consumer.Distance(actor.X, actor.Y); d < closest {
			target, closest = actor, d
		}
	}
	if target == nil {
		return Impossible{"No enemy is close enough to strike."}
	}
	engine.Log(fmt.Sprintf("A lightning bolt strikes the %s with a loud thunder, for %d damage!", target.Name, l.Damage), colorWhite)
	target.Fighter.SetHP(engine, target, target.Fighter.HP-l.Damage)
	consume(consumer, action.Item)
	return nil
}
