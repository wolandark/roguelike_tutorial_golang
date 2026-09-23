# Step 33 · Shadows

### The problem

The octant scan sees through walls. A wall should cast a **shadow**: everything behind it, as seen from the player, is hidden. In the scan, each cell covers a range of slopes; when the scan meets a wall, the slopes behind that wall are in shadow for all the rows further out. So the scan narrows its range of lit slopes as it passes walls, and for the part of the octant that is still lit on the other side of a wall it starts a new scan, which is the "recursive" part of recursive shadowcasting.

>>> 1. In `fov.go`, add a method `func (m *GameMap) blocksSight(x, y int) bool` that returns `true` for a cell outside the map or a tile that is not `Transparent`.
>>> 2. In `castLight`, after lighting a cell, handle walls: when a cell blocks sight and the scan was not already along a wall, recurse into the next row with the range narrowed to this cell's `lSlope`, and remember `rSlope` as `newStart`. While the scan continues along a wall, keep moving `newStart`; when it leaves the wall, set `start` to `newStart`.
>>> 3. After each row, stop scanning if the row ended on a wall.

!!! Rooms light up as you enter them and corridors reveal themselves cell by cell. Walls at the edge of your view are lit, so rooms have outlines. Stand in a doorway and look at the shadow the frame casts.

--- reveal

{{diff fov.go}}

- `blocked` records whether the previous cell in this row was a wall. `newStart := 0.0` declares a `float64`, because a slope has a fraction.
- When a wall starts, `castLight` calls **itself** for the next row with the range `start` to `lSlope`: the part of the octant on the near side of the wall. That call scans outwards on its own, and this call carries on along the current row.
- When the wall ends, `start = newStart` narrows the current scan to the slopes beyond the wall. A wall never lets light through; only the cells beside it continue.
- Walls themselves are lit, because the cell is lit before the wall check runs. That is what gives rooms visible outlines.

--- end

%%% Set the radius in `UpdateFOV` to 3, then to 40. At 40 you see a whole level as you enter it; the algorithm is the same, only `radiusSq` changes.

%%% Make `wall.Transparent` true in `tiles.go`. Walls still block movement but no longer block sight: the two properties are independent on purpose (think of a window or a pit).
