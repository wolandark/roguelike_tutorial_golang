package main

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
