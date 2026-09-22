# Step 15 · Walls on the map

### The problem

To put a wall somewhere we have to write `gameMap.Tiles[y*gameMap.Width+x] = wall`, and the same arithmetic already sits in `Render`. Two copies of a formula is one too many; the third copy is the one that gets a typo. Two small methods hide it, and a third guards against asking for a cell outside the map, which would panic on an index out of range, or worse, silently return a cell from the previous row for a negative `x`.

>>> Add `InBounds(x, y) bool`, `TileAt(x, y) Tile` and `SetTile(x, y, t Tile)` to `GameMap`, use `TileAt` in `Render`, and in `main.go` turn three cells of row 22 into `wall`.

!!! A darker three-cell wall in the middle of the blue floor. You can still walk through it; that is the next step.

--- reveal

{{diff gamemap.go}}

- `InBounds` must be checked before `TileAt` or `SetTile`; nothing else in the game will index `Tiles` directly.
- Methods that fit on one line are written on one line in Go; `gofmt` keeps them that way.

{{diff main.go}}

--- end

%%% Ask for `gameMap.TileAt(-1, 1)` somewhere and print the result to the log later; there is no panic, you get the last tile of row 0, index 79. Then `TileAt(-1, 0)` does panic. Both are wrong, one loudly, one quietly. That quiet one is why `InBounds` exists.
