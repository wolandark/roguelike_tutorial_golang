# Step 23 · A second room

### The problem

The player and the bystander should start in rooms of their own, and later corridors will run from room to room. Both need the same thing: the middle cell of a room. The room knows its corners, so the middle can be computed from them, and a method is the natural home for that. The middle is a pair of numbers, and Go functions can return more than one value.

>>> 1. In `procgen.go`, add a method `func (r RectangularRoom) Center() (int, int)` that returns the cell halfway between `X1` and `X2`, and halfway between `Y1` and `Y2`.
>>> 2. In `main.go`, replace the single room with two: `room1 := NewRectangularRoom(20, 15, 10, 15)` and `room2 := NewRectangularRoom(40, 28, 12, 8)`, and carve both.
>>> 3. Create the player at `room1.Center()` and the NPC at `room2.Center()`: move the two `Entity` lines below the carving and use `px, py := room1.Center()` and `nx, ny := room2.Center()` for their positions.

!!! Two separate rooms of light blue. The white `@` is in the left one, the yellow `@` in the right one, and you cannot reach it.

--- reveal

{{diff procgen.go}}

- `(int, int)` after the parameters declares **two results**. The caller receives both at once: `px, py := room1.Center()`.
- `(r.X1 + r.X2) / 2` is integer division, so the centre is always a whole cell.

{{diff main.go}}

- The entities are now created *after* the rooms, because their positions come from the rooms.

--- end

%%% Write `px := room1.Center()`. The compiler says `assignment mismatch: 1 variable but room1.Center returns 2 values`. A function with two results must be received with two variables, or with `_` for the one you do not need: `px, _ := room1.Center()`.
