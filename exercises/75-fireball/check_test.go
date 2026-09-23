package main

import "testing"

func TestFireballRange(t *testing.T) {
	m := NewGameMap(30, 12)
	for y := range m.Tiles {
		for x := range m.Tiles[y] {
			m.Tiles[y][x] = floor
		}
	}
	e := &Engine{GameMap: m, MessageLog: &MessageLog{}}
	e.Player = playerTemplate.Spawn(m, 2, 5)
	e.UpdateFOV()
	near := orc.Spawn(m, 5, 5)
	far := orc.Spawn(m, 8, 5)
	dark := orc.Spawn(m, 25, 5)
	scroll := &Entity{Name: "Short Fireball", Consumable: FireballDamageConsumable{Damage: 12, Radius: 0, MaximumRange: 4}}
	e.Player.Inventory.Items = append(e.Player.Inventory.Items, scroll)
	err := (ItemAction{Item: scroll, TargetX: dark.X, TargetY: dark.Y}).Perform(e, e.Player)
	if imp, ok := err.(Impossible); !ok || imp.Msg != "You cannot target an area that you cannot see." {
		t.Fatalf("invisible target: want the visibility message first, got %v", err)
	}
	err = (ItemAction{Item: scroll, TargetX: far.X, TargetY: far.Y}).Perform(e, e.Player)
	if imp, ok := err.(Impossible); !ok || imp.Msg != "That is too far away." {
		t.Fatalf("distance 6 with range 4: want 'That is too far away.', got %v", err)
	}
	if far.Fighter.HP != 10 || len(e.Player.Inventory.Items) != 1 {
		t.Fatal("an impossible throw must not hurt anyone or consume the scroll")
	}
	if err := (ItemAction{Item: scroll, TargetX: near.X, TargetY: near.Y}).Perform(e, e.Player); err != nil {
		t.Fatalf("distance 3 with range 4 should work: %v", err)
	}
	if near.Alive || len(e.Player.Inventory.Items) != 0 {
		t.Fatal("the near orc should be dead and the scroll used up")
	}
	unlimited := &Entity{Name: "Fireball Scroll", Consumable: FireballDamageConsumable{Damage: 12, Radius: 3}}
	e.Player.Inventory.Items = append(e.Player.Inventory.Items, unlimited)
	if err := (ItemAction{Item: unlimited, TargetX: far.X, TargetY: far.Y}).Perform(e, e.Player); err != nil {
		t.Fatalf("MaximumRange 0 must mean unlimited: %v", err)
	}
}
