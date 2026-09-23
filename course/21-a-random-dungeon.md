# Step 21 · A random dungeon

### The problem

Rooms and tunnels by hand are fine for two rooms. A level needs dozens, in random places, without overlapping. The classic generator is almost embarrassingly simple: try a random room; if it overlaps one you already placed, throw it away; otherwise carve it and dig a tunnel to the previous room. Because every kept room connects to the one before it, the level is always fully connected, without any graph algorithm. The cost is that the number of rooms is not fixed; you choose how many *attempts* to make.

To reject overlaps you need a test for two rectangles. Think about when two rooms do *not* overlap: one is entirely to the left, right, above or below the other. Negate that and you have the test.

>>> 1. In `procgen.go`, add a method `func (r RectangularRoom) Intersects(other RectangularRoom) bool`.
>>> 2. Add a function `func GenerateDungeon(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight int, player *Entity) *GameMap` that makes a map, tries `maxRooms` random rooms, skips any that intersect an earlier one, carves the rest, puts the player at the first room's centre and tunnels each later room to the previous one.
>>> 3. In `main.go`, add the constants `roomMaxSize = 10`, `roomMinSize = 6`, `maxRooms = 30`, replace the hand-made rooms and the NPC with `gameMap := GenerateDungeon(...)`, and start the player as `&Entity{Char: '@', Color: tcell.ColorWhite}`.

!!! A different dungeon every time. You start in the middle of a room; every room is reachable.

--- reveal

{{diff procgen.go}}

- `Intersects`: "not entirely to one side", written positively as four comparisons.
- `rand.IntN(n)` is uniform in `[0, n)`, hence the `+1` for an inclusive size range. The position is chosen so the room stays inside the map.
- `continue` throws away an overlapping candidate; `rooms[len(rooms)-1]` is the last kept room.
- The function receives the player pointer and *sets* its position: a side effect through a pointer, which is fine here and is what the caller expects.

{{diff main.go}}

--- end

%%% Set `maxRooms` to 300. The level fills up with rooms until nothing fits; the generator never loops forever because every attempt is independent. Now set `roomMinSize` above `roomMaxSize` and watch `rand.IntN` panic on a negative argument; the constants are a contract.
