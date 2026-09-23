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

type ConfusionConsumable struct {
	NumberOfTurns int
}

func (ConfusionConsumable) GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler) {
	engine.Log("Select a target location.", colorNeedsTarget)
	return nil, NewSingleRangedAttackHandler(engine, func(x, y int) Action {
		return ItemAction{Item: item, TargetX: x, TargetY: y}
	})
}

func (c ConfusionConsumable) Activate(engine *Engine, action ItemAction, consumer *Entity) error {
	if !engine.GameMap.IsVisible(action.TargetX, action.TargetY) {
		return Impossible{"You cannot target an area that you cannot see."}
	}
	target := action.TargetActor(engine)
	if target == nil {
		return Impossible{"You must select an enemy to target."}
	}
	if target == consumer {
		return Impossible{"You cannot confuse yourself!"}
	}
	engine.Log(fmt.Sprintf("The eyes of the %s look vacant, as it starts to stumble around!", target.Name), colorStatusEffectApplied)
	target.AI = &ConfusedEnemy{PreviousAI: target.AI, TurnsRemaining: c.NumberOfTurns}
	consume(consumer, action.Item)
	return nil
}

type FireballDamageConsumable struct {
	Damage       int
	Radius       int
	MaximumRange int // 0 means unlimited
}

func (f FireballDamageConsumable) GetAction(engine *Engine, consumer, item *Entity) (Action, EventHandler) {
	engine.Log("Select a target location.", colorNeedsTarget)
	return nil, NewAreaRangedAttackHandler(engine, f.Radius, func(x, y int) Action {
		return ItemAction{Item: item, TargetX: x, TargetY: y}
	})
}

func (f FireballDamageConsumable) Activate(engine *Engine, action ItemAction, consumer *Entity) error {
	if !engine.GameMap.IsVisible(action.TargetX, action.TargetY) {
		return Impossible{"You cannot target an area that you cannot see."}
	}
	if f.MaximumRange > 0 && consumer.Distance(action.TargetX, action.TargetY) > f.MaximumRange {
		return Impossible{"That is too far away."}
	}
	hit := false
	for _, actor := range engine.GameMap.Actors() {
		if actor.Distance(action.TargetX, action.TargetY) <= f.Radius {
			engine.Log(fmt.Sprintf("The %s is engulfed in a fiery explosion, taking %d damage!", actor.Name, f.Damage), colorWhite)
			actor.Fighter.SetHP(engine, actor, actor.Fighter.HP-f.Damage)
			hit = true
		}
	}
	if !hit {
		return Impossible{"There are no targets in the radius."}
	}
	consume(consumer, action.Item)
	return nil
}
