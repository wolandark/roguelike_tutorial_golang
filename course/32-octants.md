# Step 32 · Octants

### The problem

The square of vision ignores walls, and it is a square. Real line of sight needs a scan that starts at the player and moves outwards, so that a wall can later hide what lies behind it. The standard roguelike algorithm, **shadowcasting**, splits the area around the player into eight triangles, called **octants**, and scans each one row by row, moving away from the player. This step builds only the scan, with no shadows yet; the result is a disc of vision instead of a square. The next step adds the walls.

Writing the scan eight times, once per direction, would be eight copies of the same loop. Instead the scan is written once for one octant, in "row and column" terms, and a table of four multipliers per octant turns row and column back into map `x` and `y`.

This is one of the few places in the game where you are not expected to derive the code yourself. Type it in, then poke at it with the experiment.

>>> 1. Create `fov.go` with a package variable `var octants = [8][4]int{...}` holding the eight rows of multipliers shown below.
>>> 2. In the same file, add a method `func (m *GameMap) castLight(cx, cy, radius, row int, start, end float64, xx, xy, yx, yy int)` that walks the rows of one octant from `row` out to `radius` and calls `setVisible` on every cell within the radius.
>>> 3. Add a method `func (m *GameMap) ComputeFOV(ox, oy, radius int)` in `fov.go` that clears `Visible`, marks the player's own cell, and calls `castLight` once per octant.
>>> 4. Delete the old `ComputeFOV` from `gamemap.go`; Go does not allow two methods with the same name.

!!! A roughly round patch of lit cells around you instead of a square. Walls still do not block your view.

--- reveal

{{file fov.go}}

- `octants` has type `[8][4]int`: a fixed-size **array** of eight arrays of four ints. Unlike a slice, the size is part of the type.
- `x := cx + dx*xx + dy*xy` and `y := cy + dx*yx + dy*yy` turn the position inside the octant, `dx` along the row and `dy` the row number, into map coordinates. Each octant's four multipliers flip and swap the axes so the same loop covers a different eighth of the circle.
- `lSlope` and `rSlope` are the slopes of the two edges of the current cell as seen from the player. `start` and `end` are the range of slopes the octant covers; cells outside it are skipped. For now that range is always the whole octant, 1.0 to 0.0.
- `float64(dx)` converts an `int` to a floating-point number so the division keeps its fraction.
- `dx*dx+dy*dy < radiusSq` keeps cells within a circle of `radius`, which is why the patch is round.

{{diff gamemap.go}}

--- end

%%% Delete one of the eight rows of `octants` (and change the array size to 7). One eighth of the disc stays black: each row really is one slice of the circle.
