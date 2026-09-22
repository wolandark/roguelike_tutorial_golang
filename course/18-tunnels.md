# Step 18 · Tunnels

### The problem

Two rooms, no way between them. A corridor could be a straight diagonal line, but roguelike corridors are traditionally L-shaped: go horizontally to a corner, then vertically, or vertically first. Each leg is a straight line along one axis, so "every cell on the line" is a loop that steps one cell at a time. The only decision left is which way the L bends, and making that random is what stops every corridor from looking the same.

>>> In `procgen.go`, write `line(x1, y1, x2, y2)` returning every cell from one point to the other (both inclusive), `sign(n)` returning -1, 0 or 1, and `tunnelBetween(m, x1, y1, x2, y2)` that picks a corner at random (`math/rand/v2`, `rand.IntN(2)`) and carves both legs. Connect the two rooms' centres in `main.go`.

!!! An L-shaped corridor joins the rooms; the yellow `@` is reachable. Run it a few times: the corner flips between the two bends.

--- reveal

{{diff procgen.go}}

- `math/rand/v2` is Go's current random-number package. `rand.IntN(2)` is 0 or 1. It seeds itself, so every run differs.
- `line` returns `[][2]int`: a slice of two-element arrays, one `[x, y]` per cell. `var pts [][2]int` declares an empty slice; `append` grows it. The loop steps both coordinates towards the target and stops when both have arrived; for our axis-aligned legs only one ever moves.
- `sign` uses a `switch` with no value, which is an `if`/`else if` chain.

{{diff main.go}}

--- end

%%% Replace the coin flip with a constant so the corner is always `(x2, y1)`. Every corridor now bends the same way and the dungeon looks machine-made. Randomness in small choices is most of what makes generated levels feel natural.
