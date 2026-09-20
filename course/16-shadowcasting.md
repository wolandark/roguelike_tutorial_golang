# Step 16 · Shadowcasting

Walls should block sight. The classic answer is **recursive shadowcasting**. Split the area around the viewer into eight triangular *octants* and scan each one row by row, moving outwards. Every tile has a left and a right edge, expressed as a *slope* from the viewer; the scan keeps `start` and `end`, the range of slopes still lit. A wall casts a shadow over everything behind it that lies between its slopes, so the scan narrows the range and recurses on the part of the octant that is still lit.

Create `fov.go`:

{{file fov.go}}

- `octants` is an array of eight arrays of four multipliers. `[8][4]int` is a fixed-size array type (as opposed to a slice): the size is part of the type. The multipliers map "row and column inside the octant" back to map coordinates, in the lines `x := cx + dx*xx + dy*xy` and `y := cy + dx*yx + dy*yy`, so `castLight` is written once and called eight times.
- `ComputeFOV` replaces the naive version: clear, mark the viewer's own tile, cast light into every octant.
- `blocksSight` treats anything outside the map as rock, and otherwise asks the tile's `Transparent` flag, which we set back in step 10.
- `castLight` is the algorithm itself. You do not need to follow every line to use it; the important properties are that it marks every tile within `radius` that has a clear line to the viewer, marks each one at most once, and recurses only when it meets a wall. `float64(dx)` converts an `int` to a floating-point number, needed for the slopes. `newStart := 0.0` declares a `float64`.

Remove the naive version from `gamemap.go` (Go refuses to compile a program that declares `ComputeFOV` twice):

{{diff gamemap.go}}

!!! Run it: rooms light up as you enter them and corridors reveal themselves cell by cell. Walls at the edge of your view are lit, so rooms have visible outlines. Stand in a doorway and look at the shadow the door frame casts.
