package main

import "testing"

func TestNoMonstersInFirstRoom(t *testing.T) {
	total := 0
	for run := 0; run < 300; run++ {
		m := NewGameMap(80, 45)
		player := playerTemplate.Spawn(m, 0, 0)
		GenerateDungeon(m, 30, 6, 10, 2, player)
		for _, e := range m.Entities {
			if e == player {
				continue
			}
			total++
			if max(abs2(e.X-player.X), abs2(e.Y-player.Y)) <= 2 {
				t.Fatalf("run %d: a %s spawned at (%d,%d), next to the player at (%d,%d)", run, e.Name, e.X, e.Y, player.X, player.Y)
			}
		}
	}
	if total < 300 {
		t.Fatalf("only %d monsters in 300 dungeons: the other rooms must still get monsters", total)
	}
}

func abs2(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
