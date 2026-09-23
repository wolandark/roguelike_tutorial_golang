# Step 18 · Walls block movement

### The problem

Walk into the wall from step 17: you pass straight through. Moving never asks the map what is in the way. The simplest fix is to ask, right where the move happens: work out the destination, and only move if that cell is inside the map and walkable.

"Inside the map" needs its own check first. `TileAt` on a cell outside the map panics (see the experiment in step 17), so every coordinate we compute has to be tested before it is used.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) InBounds(x, y int) bool` that returns true when `x` is between 0 and `m.Width-1` and `y` is between 0 and `m.Height-1`, both ends included, and false otherwise.
>>> 2. In `main.go`, in the `MovementAction` case, compute `destX, destY := player.X+action.DX, player.Y+action.DY` and call `player.Move` only if `gameMap.InBounds(destX, destY) && gameMap.TileAt(destX, destY).Walkable`.

!!! Walk into the three-cell wall; the `@` stops. The map edge stops you too.

--- reveal

{{diff gamemap.go}}

- `InBounds` is four comparisons: the cell is inside when `x` is at least 0 and below `Width`, and the same for `y` and `Height`.

{{diff main.go}}

- `&&` stops at the first false operand. If the destination is out of bounds, `TileAt` is never called, which is what keeps it from panicking. The order of the two checks matters.

--- end

%%% Swap the two conditions so `Walkable` is tested first, then walk off the left edge: panic, index out of range. `&&` evaluates left to right and stops early, and the order of checks is part of correctness, not style.
