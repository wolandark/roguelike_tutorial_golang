package main

import "fmt"

type Inventory struct {
	Capacity int
	Items    []*Entity
}

func (inv *Inventory) Remove(item *Entity) {
	for i, it := range inv.Items {
		if it == item {
			inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			return
		}
	}
}

func (inv *Inventory) Drop(engine *Engine, owner, item *Entity) {
	if owner.ItemIsEquipped(item) {
		owner.ToggleEquip(engine, item, true)
	}
	inv.Remove(item)
	item.X, item.Y = owner.X, owner.Y
	engine.GameMap.Entities = append(engine.GameMap.Entities, item)
	engine.Log(fmt.Sprintf("You dropped the %s.", item.Name), colorWhite)
}
