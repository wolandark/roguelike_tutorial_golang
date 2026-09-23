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
	before := countTrue(m.Visible)
	if quit := e.HandleEvent(tcell.NewEventKey(tcell.KeyRune, 'm', 0)); quit {
		t.Fatal("'m' must not quit")
	}
	if n := countTrue(m.Explored); n != m.Width*m.Height {
		t.Fatalf("only %d of %d tiles are explored after magic mapping", n, m.Width*m.Height)
	}
	after := countTrue(m.Visible)
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

func countTrue(grid [][]bool) int {
	n := 0
	for _, row := range grid {
		for _, v := range row {
			if v {
				n++
			}
		}
	}
	return n
}
