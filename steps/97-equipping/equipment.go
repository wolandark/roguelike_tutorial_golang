package main

import "fmt"

type EquipmentType int

const (
	Weapon EquipmentType = iota
	Armor
)

type Equippable struct {
	Type         EquipmentType
	PowerBonus   int
	DefenseBonus int
	Equipped     bool
}

func (e *Entity) ItemIsEquipped(item *Entity) bool {
	return item.Equippable != nil && item.Equippable.Equipped
}

func (e *Entity) ToggleEquip(engine *Engine, item *Entity) {
	if item.Equippable.Equipped {
		e.unequip(engine, item)
		return
	}
	item.Equippable.Equipped = true
	engine.Log(fmt.Sprintf("You equip the %s.", item.Name), colorWhite)
}

func (e *Entity) unequip(engine *Engine, item *Entity) {
	item.Equippable.Equipped = false
	engine.Log(fmt.Sprintf("You remove the %s.", item.Name), colorWhite)
}
