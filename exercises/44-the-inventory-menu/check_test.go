package main

import (
	"strings"
	"testing"
)

func TestPoison(t *testing.T) {
	m := NewGameMap(10, 10)
	for y := range m.Tiles {
		for x := range m.Tiles[y] {
			m.Tiles[y][x] = floor
		}
	}
	e := &Engine{GameMap: m, MessageLog: &MessageLog{}}
	e.Player = playerTemplate.Spawn(m, 2, 2)
	var c Consumable = PoisonConsumable{Damage: 3}
	bottle := &Entity{Char: '!', Name: "Murky Potion", Consumable: c}
	e.Player.Inventory.Items = append(e.Player.Inventory.Items, bottle)
	action, handler := c.GetAction(e, e.Player, bottle)
	if handler != nil {
		t.Fatal("poison should not need a target handler")
	}
	ia, ok := action.(ItemAction)
	if !ok || ia.Item != bottle {
		t.Fatalf("GetAction should return an ItemAction for the bottle, got %#v", action)
	}
	if err := ia.Perform(e, e.Player); err != nil {
		t.Fatalf("poison must never be Impossible: %v", err)
	}
	if e.Player.Fighter.HP != 27 || len(e.Player.Inventory.Items) != 0 {
		t.Fatalf("want HP 27 and an empty inventory, got %d and %d items", e.Player.Fighter.HP, len(e.Player.Inventory.Items))
	}
	last := e.MessageLog.Messages[len(e.MessageLog.Messages)-1]
	if !strings.Contains(last.Text, "poison") || !strings.Contains(last.Text, "Murky Potion") || last.Color != colorEnemyAtk {
		t.Fatalf("unexpected message: %q", last.Text)
	}
	e.Player.Fighter.HP = 2
	deadly := &Entity{Name: "Deadly Potion", Consumable: PoisonConsumable{Damage: 10}}
	e.Player.Inventory.Items = append(e.Player.Inventory.Items, deadly)
	if err := (ItemAction{Item: deadly}).Perform(e, e.Player); err != nil {
		t.Fatal(err)
	}
	if e.Player.Alive {
		t.Fatal("a 10 damage poison at 2 HP should kill")
	}
}
