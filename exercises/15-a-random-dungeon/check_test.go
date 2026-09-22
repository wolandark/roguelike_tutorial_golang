package main

import "testing"

func TestContains(t *testing.T) {
	r := NewRectangularRoom(10, 5, 6, 4)
	for _, p := range [][2]int{{11, 6}, {15, 8}, {13, 7}} {
		if !r.Contains(p[0], p[1]) {
			t.Errorf("(%d,%d) should be inside %+v", p[0], p[1], r)
		}
	}
	for _, p := range [][2]int{{10, 6}, {16, 6}, {11, 5}, {11, 9}, {0, 0}} {
		if r.Contains(p[0], p[1]) {
			t.Errorf("(%d,%d) should be outside %+v", p[0], p[1], r)
		}
	}
	m := NewGameMap(30, 20)
	r.carve(m)
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.TileAt(x, y).Walkable != r.Contains(x, y) {
				t.Fatalf("carve and Contains disagree at (%d,%d)", x, y)
			}
		}
	}
}
