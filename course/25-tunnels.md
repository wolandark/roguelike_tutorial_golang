# Step 25 · Tunnels

### The problem

A diagonal line is not walkable, but a straight horizontal or vertical one is. So a corridor between two points can be made of two straight legs that meet at a corner: across to the target's column, then up or down to it. That is the traditional roguelike **L-shaped tunnel**. There are two possible corners, `(x2, y1)` and `(x1, y2)`, and picking one at random keeps every corridor from bending the same way.

>>> 1. In `procgen.go`, add a function `func tunnelBetween(m *GameMap, x1, y1, x2, y2 int)`. Pick the corner `(x2, y1)`, or `(x1, y2)` when `rand.IntN(2)` returns 0, and set every cell of `line` from the start to the corner, and from the corner to the end, to `floor`.
>>> 2. Import `math/rand/v2` in `procgen.go`.
>>> 3. In `main.go`, replace the `line` loop with `tunnelBetween(gameMap, x1, y1, x2, y2)`.

!!! An L-shaped corridor joins the rooms, and the yellow `@` is reachable. Run it a few times: the corner flips between the two bends.

--- reveal

{{diff procgen.go}}

- `math/rand/v2` is Go's current random-number package. `rand.IntN(2)` returns 0 or 1. It seeds itself, so every run differs.
- Each leg is a `line` whose ends share a row or a column, so it is straight and walkable.

{{diff main.go}}

--- end

%%% Replace the coin flip with a fixed corner, always `(x2, y1)`. Every corridor now bends the same way, which in a full dungeon looks machine-made. Small random choices are most of what makes generated levels feel natural.
