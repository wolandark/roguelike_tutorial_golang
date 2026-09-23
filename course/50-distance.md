# Step 50 · Distance

### The problem

A monster has to know whether it is next to the player, so it can attack instead of walking. "Next to" in a grid game with diagonal moves means one king's move away: the larger of the horizontal and the vertical difference is 1. That measure is useful for many things later (the range of a spell, the size of a blast), so it becomes a method on `Entity`.

Go's standard library has `math.Abs` for floating-point numbers but no absolute value for integers, so that needs a three-line helper.

>>> 1. In `entity.go`, add a function `func abs(n int) int` that returns `-n` for a negative `n` and `n` otherwise.
>>> 2. Add a method `func (e *Entity) Distance(x, y int) int` that returns the larger of `abs(x-e.X)` and `abs(y-e.Y)`.

!!! Nothing changes on screen; the next step uses it.

--- reveal

{{diff entity.go}}

- `max(a, b)` is a built-in function (Go 1.21 or newer), like `min`.
- This is called **Chebyshev distance**: the number of king's moves between two cells. Diagonal neighbours are at distance 1, like straight ones.

--- end

%%% Compute `abs(x-e.X) + abs(y-e.Y)` instead. That is Manhattan distance, where a diagonal neighbour is 2 away. In the next step, monsters standing diagonally next to you would stop attacking.
