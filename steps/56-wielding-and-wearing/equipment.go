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

func (e *Entity) EquippedItem(slot EquipmentType) *Entity {
	if e.Inventory == nil {
		return nil
	}
	for _, item := range e.Inventory.Items {
		if item.Equippable != nil && item.Equippable.Equipped && item.Equippable.Type == slot {
			return item
		}
	}
	return nil
}

func (e *Entity) ItemIsEquipped(item *Entity) bool {
	return item.Equippable != nil && item.Equippable.Equipped
}

func (e *Entity) PowerBonus() int {
	bonus := 0
	for _, slot := range []EquipmentType{Weapon, Armor} {
		if item := e.EquippedItem(slot); item != nil {
			bonus += item.Equippable.PowerBonus
		}
	}
	return bonus
}

func (e *Entity) DefenseBonus() int {
	bonus := 0
	for _, slot := range []EquipmentType{Weapon, Armor} {
		if item := e.EquippedItem(slot); item != nil {
			bonus += item.Equippable.DefenseBonus
		}
	}
	return bonus
}

func (e *Entity) ToggleEquip(engine *Engine, item *Entity, addMessage bool) {
	if item.Equippable.Equipped {
		e.unequip(engine, item, addMessage)
		return
	}
	if current := e.EquippedItem(item.Equippable.Type); current != nil {
		e.unequip(engine, current, addMessage)
	}
	item.Equippable.Equipped = true
	if addMessage {
		engine.Log(fmt.Sprintf("You equip the %s.", item.Name), colorWhite)
	}
}

func (e *Entity) unequip(engine *Engine, item *Entity, addMessage bool) {
	item.Equippable.Equipped = false
	if addMessage {
		engine.Log(fmt.Sprintf("You remove the %s.", item.Name), colorWhite)
	}
}
