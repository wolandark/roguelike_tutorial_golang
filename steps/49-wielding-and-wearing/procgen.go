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

func (r RectangularRoom) Intersects(other RectangularRoom) bool {
	return r.X1 <= other.X2 && r.X2 >= other.X1 &&
		r.Y1 <= other.Y2 && r.Y2 >= other.Y1
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

type floorValue struct{ Floor, Value int }

var (
	maxItemsByFloor    = []floorValue{{1, 1}, {4, 2}}
	maxMonstersByFloor = []floorValue{{1, 2}, {4, 3}, {6, 5}}
)

type spawnChance struct {
	Template *Entity
	Weight   int
}

var (
	itemChances = map[int][]spawnChance{
		0: {{&healthPotion, 35}},
		2: {{&confusionScroll, 10}},
		4: {{&lightningScroll, 25}, {&sword, 5}},
		6: {{&fireballScroll, 25}, {&chainMail, 15}},
	}
	enemyChances = map[int][]spawnChance{
		0: {{&orc, 80}},
		3: {{&troll, 15}},
		5: {{&troll, 30}},
		7: {{&troll, 60}},
	}
)

func maxValueForFloor(values []floorValue, floor int) int {
	current := 0
	for _, v := range values {
		if v.Floor > floor {
			break
		}
		current = v.Value
	}
	return current
}

func entitiesAtRandom(chances map[int][]spawnChance, n, floor int) []*Entity {
	weights := map[*Entity]int{}
	var order []*Entity
	for f := 0; f <= floor; f++ {
		for _, c := range chances[f] {
			if _, seen := weights[c.Template]; !seen {
				order = append(order, c.Template)
			}
			weights[c.Template] = c.Weight
		}
	}
	total := 0
	for _, t := range order {
		total += weights[t]
	}
	var chosen []*Entity
	for i := 0; i < n && total > 0; i++ {
		roll := rand.IntN(total)
		for _, t := range order {
			roll -= weights[t]
			if roll < 0 {
				chosen = append(chosen, t)
				break
			}
		}
	}
	return chosen
}

func placeEntities(room RectangularRoom, dungeon *GameMap, floor int) {
	nMonsters := rand.IntN(maxValueForFloor(maxMonstersByFloor, floor) + 1)
	nItems := rand.IntN(maxValueForFloor(maxItemsByFloor, floor) + 1)

	monsters := entitiesAtRandom(enemyChances, nMonsters, floor)
	items := entitiesAtRandom(itemChances, nItems, floor)

	for _, template := range append(monsters, items...) {
		x, y := room.randomTile()
		if dungeon.EntityAt(x, y) != nil {
			continue
		}
		template.Spawn(dungeon, x, y)
	}
}

func (r RectangularRoom) randomTile() (int, int) {
	return r.X1 + 1 + rand.IntN(r.X2-r.X1-1), r.Y1 + 1 + rand.IntN(r.Y2-r.Y1-1)
}

func GenerateDungeon(dungeon *GameMap, maxRooms, roomMinSize, roomMaxSize, currentFloor int, player *Entity) {
	var rooms []RectangularRoom
	var centerOfLastRoom [2]int

	for i := 0; i < maxRooms; i++ {
		roomWidth := roomMinSize + rand.IntN(roomMaxSize-roomMinSize+1)
		roomHeight := roomMinSize + rand.IntN(roomMaxSize-roomMinSize+1)
		x := rand.IntN(dungeon.Width - roomWidth - 1)
		y := rand.IntN(dungeon.Height - roomHeight - 1)

		newRoom := NewRectangularRoom(x, y, roomWidth, roomHeight)

		overlaps := false
		for _, other := range rooms {
			if newRoom.Intersects(other) {
				overlaps = true
				break
			}
		}
		if overlaps {
			continue
		}

		newRoom.carve(dungeon)

		if len(rooms) == 0 {
			player.X, player.Y = newRoom.Center()
		} else {
			px, py := rooms[len(rooms)-1].Center()
			cx, cy := newRoom.Center()
			tunnelBetween(dungeon, px, py, cx, cy)
			centerOfLastRoom = [2]int{cx, cy}
		}
		placeEntities(newRoom, dungeon, currentFloor)
		rooms = append(rooms, newRoom)
	}

	dungeon.SetTile(centerOfLastRoom[0], centerOfLastRoom[1], downStairs)
	dungeon.DownstairsX, dungeon.DownstairsY = centerOfLastRoom[0], centerOfLastRoom[1]
}
