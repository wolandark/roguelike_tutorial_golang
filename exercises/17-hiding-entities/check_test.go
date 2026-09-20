package main

import (
	"testing"

	"github.com/gdamore/tcell/v2"
)

func TestMagicMapping(t *testing.T) {
	player := &Entity{Char: '@'}
	m := GenerateDungeon(30, 6, 10, 80, 45, player)
	e := &Engine{Entities: []*Entity{player}, Player: player, GameMap: m}
	e.UpdateFOV()
	before := 0
	for _, v := range m.Visible {
		if v {
			before++
		}
	}
	if quit := e.HandleEvent(tcell.NewEventKey(tcell.KeyRune, 'm', 0)); quit {
		t.Fatal("'m' must not quit")
	}
	for i, ex := range m.Explored {
		if !ex {
			t.Fatalf("tile %d is not explored after magic mapping", i)
		}
	}
	after := 0
	for _, v := range m.Visible {
		if v {
			after++
		}
	}
	if after != before {
		t.Fatalf("Visible changed: %d -> %d", before, after)
	}
	x, y := player.X, player.Y
	e.HandleEvent(tcell.NewEventKey(tcell.KeyRune, 'm', 0))
	if player.X != x || player.Y != y {
		t.Fatal("'m' moved the player")
	}
	if !e.HandleEvent(tcell.NewEventKey(tcell.KeyEscape, 0, 0)) {
		t.Fatal("escape broke")
	}
}
