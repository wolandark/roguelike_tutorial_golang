# Step 14 · A random dungeon

Rooms plus tunnels plus a loop that tries random rooms and keeps the ones that fit: that is the whole generator. Each new room is connected to the previous one, so the dungeon is always fully connected, and the player starts in the first room.

Add to `procgen.go`:

{{diff procgen.go}}

- `Intersects` is true unless one room is entirely left of, right of, above or below the other. Written positively, that is the four comparisons.
- `GenerateDungeon` *tries* `maxRooms` times. Each try picks a random size within the limits (`rand.IntN(n)` is uniform in `[0, n)`, hence the `+1`) and a random position that keeps the room inside the map. If the candidate overlaps a room we already have, `continue` throws it away and tries again, so the real number of rooms is usually smaller than `maxRooms`.
- `len(rooms) == 0` identifies the first room: the player goes to its centre. Every later room gets a tunnel from the previous room's centre (`rooms[len(rooms)-1]` is the last element).
- The function receives the player pointer and *sets* its position, then returns the finished map.

`main.go` gets three knobs and loses the hand-made rooms and the NPC:

{{diff main.go}}

- Note `player := &Entity{Char: '@', Color: tcell.ColorWhite}`: the position is left at zero and filled in by the generator.

!!! Run it: a different dungeon every time. You start in the middle of a room; every room is reachable. The whole map is visible, which spoils exploration; that is chapter 4.
