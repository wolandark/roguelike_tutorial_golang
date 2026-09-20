package main

import "testing"

func TestShieldSlot(t *testing.T) {
	if Shield == Weapon || Shield == Armor {
		t.Fatal("Shield must be its own slot")
	}
	e := NewGame()
	p := e.Player
	buckler := &Entity{Name: "Buckler", Equippable: &Equippable{Type: Shield, DefenseBonus: 2}}
	tower := &Entity{Name: "Tower Shield", Equippable: &Equippable{Type: Shield, DefenseBonus: 4, PowerBonus: -1}}
	p.Inventory.Items = append(p.Inventory.Items, buckler, tower)
	p.ToggleEquip(e, buckler, false)
	if p.EquippedItem(Shield) != buckler || p.Defense() != 4 || p.Power() != 4 {
		t.Fatalf("with buckler: want defense 4 / power 4, got %d / %d", p.Defense(), p.Power())
	}
	if p.EquippedItem(Armor) == nil || p.EquippedItem(Weapon) == nil {
		t.Fatal("equipping a shield must not touch the other slots")
	}
	p.ToggleEquip(e, tower, false)
	if buckler.Equippable.Equipped || p.EquippedItem(Shield) != tower || p.Defense() != 6 || p.Power() != 3 {
		t.Fatal("equipping a second shield must unequip the first and apply its bonuses")
	}
	p.ToggleEquip(e, tower, false)
	if p.EquippedItem(Shield) != nil || p.Defense() != 2 {
		t.Fatal("unequipping the shield should drop its bonus")
	}
}
