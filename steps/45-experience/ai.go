package main

import (
	"fmt"
	"math/rand/v2"
)

type AI interface {
	Perform(engine *Engine, entity *Entity) error
}

type HostileEnemy struct{}

func (HostileEnemy) Perform(engine *Engine, entity *Entity) error {
	target := engine.Player
	dx, dy := target.X-entity.X, target.Y-entity.Y

	if engine.GameMap.IsVisible(entity.X, entity.Y) {
		if entity.Distance(target.X, target.Y) <= 1 {
			return MeleeAction{ActionWithDirection{dx, dy}}.Perform(engine, entity)
		}
		path := FindPath(engine.GameMap, entity.X, entity.Y, target.X, target.Y)
		if len(path) > 0 {
			next := path[0]
			return MovementAction{ActionWithDirection{next[0] - entity.X, next[1] - entity.Y}}.Perform(engine, entity)
		}
	}
	return WaitAction{}.Perform(engine, entity)
}

type ConfusedEnemy struct {
	PreviousAI     AI
	TurnsRemaining int
}

var directions = [][2]int{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (c *ConfusedEnemy) Perform(engine *Engine, entity *Entity) error {
	if c.TurnsRemaining <= 0 {
		engine.Log(fmt.Sprintf("The %s is no longer confused.", entity.Name), colorWhite)
		entity.AI = c.PreviousAI
		return nil
	}
	c.TurnsRemaining--
	d := directions[rand.IntN(len(directions))]
	return BumpAction{ActionWithDirection{d[0], d[1]}}.Perform(engine, entity)
}
