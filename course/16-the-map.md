# Step 16 · The map

### The problem

Drawing a room by hand with loops does not scale to a dungeon, and the player cannot ask a loop what is under their feet. The tiles have to be *stored*, so that drawing reads them and, next step, movement can too. The map is a grid of tiles, and the question is how to store a grid in Go. The Python tutorial uses a 2D numpy array. Go has no numpy, and a slice of slices is awkward to allocate and slow to copy. The idiomatic choice is **one flat slice** with a little arithmetic: the tile at `(x, y)` sits at index `y*width + x`. This step creates the map, fills it with floor and draws it; the arithmetic appears twice for now, and the next step puts it in one place.

>>> 1. Create `gamemap.go` and declare a struct `GameMap` with fields `Width, Height int` and `Tiles []Tile`.
>>> 2. Add a function `func NewGameMap(width, height int) *GameMap` that allocates `Tiles` with `make([]Tile, width*height)`, fills every element with `floor`, and returns the map.
>>> 3. Add a method `func (m *GameMap) Render(screen tcell.Screen)` that loops over every `y` and `x` and draws `m.Tiles[y*m.Width+x].Dark`.
>>> 4. In `main.go`, add the constants `mapWidth = 80` and `mapHeight = 45`, create `gameMap := NewGameMap(mapWidth, mapHeight)` before the loop, and replace the hand-drawn room with `gameMap.Render(screen)`.

!!! The whole screen is blue floor, except the bottom five rows, and the two `@`s stand on it.

--- reveal

{{file gamemap.go}}

- `make([]Tile, width*height)` allocates the slice; every element starts as the zero `Tile` (all `false`, blank glyph), so `NewGameMap` fills it with `floor` in a loop. `for i := range m.Tiles` with one variable gives the index.
- `m.Tiles[y*m.Width+x]` is the flat-slice arithmetic: row `y` starts at `y*Width`.

{{diff main.go}}

- The map is 80 wide but only 45 tall, leaving five rows at the bottom for a status area later.

--- end

%%% Swap the two loops in `Render` so `x` is the outer one. Nothing changes on screen, because each cell is still drawn once; but with the index formula `y*Width + x`, the inner loop of the original visits memory in order, which is why the row loop goes outside. It does not matter at this size. It is the habit that matters.
