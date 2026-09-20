package main

var octants = [8][4]int{
	{1, 0, 0, 1}, {0, 1, 1, 0}, {0, -1, 1, 0}, {-1, 0, 0, 1},
	{-1, 0, 0, -1}, {0, -1, -1, 0}, {0, 1, -1, 0}, {1, 0, 0, -1},
}

func (m *GameMap) ComputeFOV(ox, oy, radius int) {
	for i := range m.Visible {
		m.Visible[i] = false
	}
	m.setVisible(ox, oy)
	for _, o := range octants {
		m.castLight(ox, oy, radius, 1, 1.0, 0.0, o[0], o[1], o[2], o[3])
	}
}

func (m *GameMap) blocksSight(x, y int) bool {
	return !m.InBounds(x, y) || !m.TileAt(x, y).Transparent
}

func (m *GameMap) castLight(cx, cy, radius, row int, start, end float64, xx, xy, yx, yy int) {
	if start < end {
		return
	}
	radiusSq := radius * radius
	for j := row; j <= radius; j++ {
		dx, dy := -j-1, -j
		blocked := false
		newStart := 0.0
		for dx <= 0 {
			dx++
			x := cx + dx*xx + dy*xy
			y := cy + dx*yx + dy*yy
			lSlope := (float64(dx) - 0.5) / (float64(dy) + 0.5)
			rSlope := (float64(dx) + 0.5) / (float64(dy) - 0.5)
			if start < rSlope {
				continue
			} else if end > lSlope {
				break
			}
			if dx*dx+dy*dy < radiusSq {
				m.setVisible(x, y)
			}
			if blocked {
				if m.blocksSight(x, y) {
					newStart = rSlope
					continue
				}
				blocked = false
				start = newStart
			} else if m.blocksSight(x, y) && j < radius {
				blocked = true
				m.castLight(cx, cy, radius, j+1, start, lSlope, xx, xy, yx, yy)
				newStart = rSlope
			}
		}
		if blocked {
			break
		}
	}
}
