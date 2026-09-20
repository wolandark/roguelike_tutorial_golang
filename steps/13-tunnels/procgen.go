package main

import "math/rand/v2"

type RectangularRoom struct {
	X1, Y1, X2, Y2 int
}

func NewRectangularRoom(x, y, width, height int) RectangularRoom {
	return RectangularRoom{X1: x, Y1: y, X2: x + width, Y2: y + height}
}

func (r RectangularRoom) Center() (int, int) {
	return (r.X1 + r.X2) / 2, (r.Y1 + r.Y2) / 2
}

func (r RectangularRoom) carve(m *GameMap) {
	for y := r.Y1 + 1; y < r.Y2; y++ {
		for x := r.X1 + 1; x < r.X2; x++ {
			m.SetTile(x, y, floor)
		}
	}
}

func tunnelBetween(m *GameMap, x1, y1, x2, y2 int) {
	cornerX, cornerY := x2, y1
	if rand.IntN(2) == 0 {
		cornerX, cornerY = x1, y2
	}
	for _, p := range line(x1, y1, cornerX, cornerY) {
		m.SetTile(p[0], p[1], floor)
	}
	for _, p := range line(cornerX, cornerY, x2, y2) {
		m.SetTile(p[0], p[1], floor)
	}
}

func line(x1, y1, x2, y2 int) [][2]int {
	var pts [][2]int
	dx, dy := sign(x2-x1), sign(y2-y1)
	x, y := x1, y1
	for {
		pts = append(pts, [2]int{x, y})
		if x == x2 && y == y2 {
			return pts
		}
		if x != x2 {
			x += dx
		}
		if y != y2 {
			y += dy
		}
	}
}

func sign(n int) int {
	switch {
	case n < 0:
		return -1
	case n > 0:
		return 1
	}
	return 0
}
