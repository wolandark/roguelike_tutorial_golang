# Step 17 · Walls on the map

### The problem

To put a wall somewhere we have to write `gameMap.Tiles[y*gameMap.Width+x] = wall`, and the same formula already sits in `Render`. Every place that repeats it is a place where `y*Width+x` can be mistyped as `x*Width+y`, which compiles and puts the wall somewhere else. So the formula goes into two small methods, one to read a cell and one to write it, and from here on nothing else in the game indexes `Tiles` directly.

>>> Add `TileAt(x, y) Tile` and `SetTile(x, y, t Tile)` to `GameMap`, use `TileAt` in `Render`, and in `main.go` turn three cells of row 22 into `wall`.

!!! A darker three-cell wall in the middle of the blue floor, this time stored in the map rather than drawn by hand. You can still walk through it; that is the next step.

--- reveal

{{diff gamemap.go}}

- Methods that fit on one line are written on one line in Go; `gofmt` keeps them that way.
- `Render` no longer knows the index formula. If the storage ever changes, `TileAt` and `SetTile` are the only two places to touch.

{{diff main.go}}

--- end

%%% Call `gameMap.SetTile(80, 0, wall)` once, one column past the right edge. No panic: index `0*80+80` is 80, the first cell of row 1, so the wall appears at the left of the *next* row. Now try `SetTile(-1, 0, wall)`: that one panics. Asking for a cell outside the map is wrong either way, quietly or loudly, and the next step adds the check for it.
