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

// PoisonConsumable hurts whoever drinks it. Your turn!
type PoisonConsumable struct {
	Damage int
}

func (PoisonConsumable) GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler) {
	return ItemAction{Item: item}, nil
}

func (p PoisonConsumable) Activate(engine *Engine, action ItemAction, consumer *Entity) error {
	engine.Log(fmt.Sprintf("You drink the %s... it was poison! You lose %d HP.", action.Item.Name, p.Damage), colorEnemyAtk)
	consumer.Fighter.SetHP(engine, consumer, consumer.Fighter.HP-p.Damage)
	consume(consumer, action.Item)
	return nil
}
