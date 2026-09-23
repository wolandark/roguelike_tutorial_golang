# Step 16 · The map

### The problem

Drawing a room by hand with loops does not scale to a dungeon, and the player cannot ask a loop what is under their feet. The tiles have to be *stored*, so that drawing reads them and, next step, movement can too. The map is a grid of tiles, and Go stores a grid as a **slice of slices**: `Tiles[y]` is one row, a `[]Tile`, and `Tiles[y][x]` is one cell. Go has no built-in 2D type, so the grid is built in two stages: first the outer slice with one entry per row, then each row with one entry per column. This step creates the map, fills it with floor and draws it.

>>> 1. Create `gamemap.go` and declare a struct `GameMap` with fields `Width, Height int` and `Tiles [][]Tile`.
>>> 2. Add a function `func NewGameMap(width, height int) *GameMap` that allocates the rows with `make([][]Tile, height)`, then each row with `make([]Tile, width)`, sets every cell to `floor`, and returns the map.
>>> 3. Add a method `func (m *GameMap) Render(screen tcell.Screen)` that loops over every `y` and `x` and draws `m.Tiles[y][x].Dark`.
>>> 4. In `main.go`, add the constants `mapWidth = 80` and `mapHeight = 45`, create `gameMap := NewGameMap(mapWidth, mapHeight)` before the loop, and replace the hand-drawn room with `gameMap.Render(screen)`.

!!! The whole screen is blue floor, except the bottom five rows, and the two `@`s stand on it.

--- reveal

{{file gamemap.go}}

- `make([][]Tile, height)` makes the outer slice: `height` rows, each still `nil`. The loop gives every row its own `make([]Tile, width)`. Forgetting that second `make` is the classic mistake: the rows stay `nil` and the first `m.Tiles[y][x] = ...` panics with `index out of range [0] with length 0`.
- Every new `Tile` is the zero value (all `false`, blank glyph), so the inner loop fills the row with `floor`. `for y := range m.Tiles` with one variable gives the index.
- `m.Tiles[y][x]`: row first, then column. The outer index is `y` because a row is a horizontal line of cells.

{{diff main.go}}

- The map is 80 wide but only 45 tall, leaving five rows at the bottom for a status area later.

--- end

%%% Delete the line `m.Tiles[y] = make([]Tile, width)` and run: `panic: runtime error: index out of range [0] with length 0`. The outer `make` only creates the rows; each row is a slice of its own and needs its own allocation.

%%% Write `m.Tiles[x][y]` in `Render` instead of `m.Tiles[y][x]`. It panics as soon as `x` passes 44, because there are only 45 rows. Row first, then column, every time.
