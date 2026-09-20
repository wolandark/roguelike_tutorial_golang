# Step 13 · Tunnels

Rooms need corridors. The simplest tunnel is an L: go horizontally to a corner, then vertically to the destination, or the other way round. Each leg is a straight line, so "every cell on the line" is a loop that steps one cell at a time.

Add to `procgen.go`:

{{diff procgen.go}}

- `import "math/rand/v2"` is Go's current random-number package. `rand.IntN(2)` is 0 or 1: a coin flip that decides which way the L bends, so tunnels do not all look the same. It seeds itself; every run is different.
- `line` returns `[][2]int`: a slice of two-element arrays, one `[x, y]` per cell. `var pts [][2]int` declares an empty slice; `append` grows it. The loop steps `x` and `y` towards the target and stops when both have arrived; for our horizontal and vertical legs only one of them ever moves.
- `sign` squeezes any difference into -1, 0 or 1: the direction to step. A `switch` with no value is an `if`/`else if` chain.
- `tunnelBetween` digs both legs. Two `for ... range` loops over the cells `line` returns; `p[0]` and `p[1]` are the x and y of each.

Connect the two rooms in `main.go`:

{{diff main.go}}

!!! Run it: an L-shaped corridor joins the rooms; the yellow `@` is reachable now. Run it a few times: the corner flips between the two possible bends.
