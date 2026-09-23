package main

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
