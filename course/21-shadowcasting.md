# Step 21 · Shadowcasting

### The problem

The naive rule lets you see through walls. Real line of sight: a cell is visible if the straight line from you to it crosses no wall. Checking that line for every cell in range works but is slow and produces ugly artefacts. The standard roguelike answer is **recursive shadowcasting**: divide the area around the viewer into eight triangular *octants*, scan each one row by row outwards, and describe each cell by the *slopes* of its two edges as seen from the viewer. A wall casts a shadow, a range of slopes, over everything behind it, so the scan narrows the lit range and recurses only for the part that is still lit. Every cell is visited once, and walls are lit but nothing behind them is.

This is the one piece of the game where you are not expected to derive the code yourself. Read it, then poke at it with the experiments.

>>> Create `fov.go` with a `ComputeFOV` that clears `Visible`, marks the viewer's cell and casts light into each of eight octants using a multiplier table, and a `blocksSight(x, y)` that treats out-of-map cells as rock and otherwise asks the tile's `Transparent` flag. Remove the naive `ComputeFOV` from `gamemap.go`, or the compiler will complain about a duplicate.

!!! Rooms light up as you enter them and corridors reveal themselves cell by cell. Walls at the edge of your view are lit, so rooms have outlines. Stand in a doorway and look at the shadow the frame casts.

--- reveal

{{file fov.go}}

- `octants` is `[8][4]int`, a fixed-size array type: the size is part of the type. The four multipliers map "row and column inside the octant" back to map coordinates in `x := cx + dx*xx + dy*xy` and `y := cy + dx*yx + dy*yy`, so `castLight` is written once and called eight times.
- `castLight` keeps `start` and `end`, the range of slopes still lit. `lSlope`/`rSlope` are the edges of the current cell. When it meets a wall it recurses for the rows behind, with the range narrowed to `lSlope`, and continues its own row with `newStart` after the wall ends. `float64(dx)` converts an `int` for the division; `newStart := 0.0` declares a `float64`.

{{diff gamemap.go}}

--- end

%%% Set the radius in `UpdateFOV` to 3, then to 40. At 40 you see a whole level as you enter it; the algorithm is the same, only `radiusSq` changes.

%%% Make `wall.Transparent` true in `tiles.go`. Walls still block movement but no longer block sight: the two properties are independent on purpose (think of a window or a pit).
