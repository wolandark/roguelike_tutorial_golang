# Step 17 · Walls on the map

### The problem

To put a wall somewhere we write `gameMap.Tiles[y][x] = wall`, and `Render` reads `m.Tiles[y][x]`. Every place that indexes the grid directly is a place where `[y][x]` can be swapped to `[x][y]`, which compiles and either panics or touches the wrong cell. So reading and writing a cell go into two small methods that take `x, y` in the order the rest of the game uses, and from here on nothing else indexes `Tiles` directly.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) TileAt(x, y int) Tile` that returns `m.Tiles[y][x]`.
>>> 2. Add a method `func (m *GameMap) SetTile(x, y int, t Tile)` that assigns `t` to the same element.
>>> 3. In `Render`, replace the index expression with `m.TileAt(x, y)`.
>>> 4. In `main.go`, after creating the map, call `gameMap.SetTile(x, 22, wall)` for `x` from 30 to 32.

!!! A darker three-cell wall in the middle of the blue floor, this time stored in the map rather than drawn by hand. You can still walk through it; that is the next step.

--- reveal

{{diff gamemap.go}}

- Methods that fit on one line are written on one line in Go; `gofmt` keeps them that way.
- `Render` no longer indexes the grid itself. The `[y][x]` order now lives in exactly two places, and everything else says `TileAt(x, y)`, with `x` first like every other coordinate in the game.

{{diff main.go}}

--- end

%%% Call `gameMap.SetTile(80, 0, wall)` once, one column past the right edge: `panic: runtime error: index out of range [80] with length 80`. Asking for a cell outside the map crashes the game, and the player walking off the edge will do exactly that. The next step adds the check for it.
