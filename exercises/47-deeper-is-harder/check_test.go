package main

import (
	"reflect"
	"testing"
)

func TestSpawnTable(t *testing.T) {
	cases := []struct {
		chances map[int][]spawnChance
		floor   int
		want    map[string]int
	}{
		{enemyChances, 1, map[string]int{"Orc": 80}},
		{enemyChances, 3, map[string]int{"Orc": 80, "Troll": 15}},
		{enemyChances, 7, map[string]int{"Orc": 80, "Troll": 60}},
		{itemChances, 5, map[string]int{"Health Potion": 35, "Confusion Scroll": 10, "Lightning Scroll": 25}},
	}
	for _, c := range cases {
		if got := spawnTable(c.chances, c.floor); !reflect.DeepEqual(got, c.want) {
			t.Errorf("floor %d: want %v, got %v", c.floor, c.want, got)
		}
	}
	custom := map[int][]spawnChance{0: {{&orc, 10}}, 2: {{&orc, 1}, {&troll, 5}}}
	if got := spawnTable(custom, 2); got["Orc"] != 1 || got["Troll"] != 5 {
		t.Errorf("a deeper floor must replace the weight: %v", got)
	}
	if got := spawnTable(custom, 1); len(got) != 1 || got["Orc"] != 10 {
		t.Errorf("floor 1 must not include floor 2 entries: %v", got)
	}
}
