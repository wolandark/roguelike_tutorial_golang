# Step 26 · Random rooms

### The problem

Rooms placed by hand are fine for two. A level needs dozens in different places every time. So the rooms get a random size and a random position, and the player starts in the first one. The size is a number between a minimum and a maximum; the position must keep the whole room inside the map.

This step only scatters rooms. It does not check whether they overlap and does not connect them yet; both come next, and seeing the result without them shows why they are needed.

>>> 1. In `procgen.go`, add a function `func GenerateDungeon(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight int, player *Entity) *GameMap` that creates a map, then `maxRooms` times: picks a width and a height from `roomMinSize` to `roomMaxSize`, picks `x` and `y` so the room fits inside the map, and carves the room. Put the player at the centre of the first room. Return the map.
>>> 2. In `main.go`, add the constants `roomMaxSize = 10`, `roomMinSize = 6` and `maxRooms = 30`.
>>> 3. Replace the rooms, the tunnel and the NPC with `player := &Entity{Char: '@', Color: tcell.ColorWhite}` followed by `gameMap := GenerateDungeon(maxRooms, roomMinSize, roomMaxSize, mapWidth, mapHeight, player)`, and give the engine only the player as its entity.

!!! A different scattering of rooms on every run, some overlapping into odd shapes. You start in one of them and can only walk where rooms happen to touch.

--- reveal

{{diff procgen.go}}

- `rand.IntN(n)` returns a number from 0 up to, but not including, `n`. So `roomMinSize + rand.IntN(roomMaxSize-roomMinSize+1)` is a size from `roomMinSize` up to and including `roomMaxSize`.
- `x := rand.IntN(dungeon.Width - roomWidth - 1)` keeps the room's far edge inside the map.
- `len(rooms) == 0` is true only for the first room carved, which is where the player goes. The function sets the player's position through the pointer it was given.

{{diff main.go}}

--- end

%%% Set `roomMinSize` to 11, above `roomMaxSize`. `rand.IntN` receives a negative number and panics: `invalid argument to IntN`. The two constants have to agree; nothing else checks that for you.
