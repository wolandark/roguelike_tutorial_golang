# Step 24 · Lines

### The problem

The rooms need a way between them. The first idea is the direct one: carve every cell on the straight line from one centre to the other. To do that we need the list of cells on a line. Stepping from the start towards the end one cell at a time works: each step moves `x` one closer to the target, and `y` one closer to its target, until both have arrived.

"One closer" means +1, -1 or 0 depending on which side the target is. That small function is called `sign`.

>>> 1. In `procgen.go`, add a function `func sign(n int) int` that returns -1 for a negative `n`, 1 for a positive `n`, and 0 for zero.
>>> 2. Add a function `func line(x1, y1, x2, y2 int) [][2]int` that returns every cell from `x1, y1` to `x2, y2`, both ends included: start at the first point, and repeatedly add `sign(x2-x)` to `x` and `sign(y2-y)` to `y` until both equal the target.
>>> 3. In `main.go`, after carving the rooms, get both centres and set every cell of `line(x1, y1, x2, y2)` to `floor`.

!!! A diagonal staircase of floor cells between the two rooms. Try to walk along it: you can't. The cells only touch at their corners, and the `@` only moves up, down, left and right.

--- reveal

{{diff procgen.go}}

- `sign` is a `switch` with no value: each `case` is a condition, like the key switch in step 8.
- `[][2]int` is a slice of two-element arrays: one `[x, y]` per cell. `var pts [][2]int` starts empty and `append` grows it.
- Both coordinates move on the same iteration, which is why a line whose ends differ in both `x` and `y` comes out diagonal.
- `for { ... }` loops until the `return` inside it runs, once both coordinates have reached the target.

{{diff main.go}}

- `p[0]` and `p[1]` are the `x` and `y` of each cell.

--- end

%%% Change `room2` to `NewRectangularRoom(40, 15, 12, 15)`, so its centre is on the same row as `room1`'s. The line is now horizontal and you can walk along it. Only lines that are straight along one axis are walkable, which is exactly what the next step builds on.
