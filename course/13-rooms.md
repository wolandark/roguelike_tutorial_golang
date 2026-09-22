# Step 13 · Rooms
## Chapter: Generating a dungeon

### In this chapter

A dungeon is carved, not built: start with solid rock, cut rooms into it, join them with tunnels, place the player in the first room. Three steps: rooms, tunnels, and the loop that places random rooms until the level is full. All of it is plain arithmetic on rectangles; the only new Go is random numbers and multiple return values.

### The problem

`NewGameMap` fills the map with floor and we place walls by hand. For carving, the default has to flip: everything is wall, and a room is a rectangle of floor. What should a room remember? Its corners are enough: from them you can compute its centre (where to put the player, where tunnels start) and, later, whether it overlaps another room. One subtlety decides how the dungeon looks: if two rooms are carved right next to each other they merge into one hall, unless each room keeps a ring of wall around itself. So carving the *inside* only is a design choice, not an accident.

>>> Flip `NewGameMap` to fill with `wall`. Create `procgen.go` with a `RectangularRoom` (two corners, the far ones exclusive like slice bounds), `NewRectangularRoom(x, y, width, height)`, `Center()` returning two ints, and a `carve` method that turns the room's inside, excluding a one-tile border, into floor. In `main.go`, carve two rooms and put the player and the NPC at their centres.

!!! Two rooms of light blue in a sea of dark blue. The `@` is in the left one and cannot leave it.

--- reveal

{{diff gamemap.go}}

{{file procgen.go}}

- `X2` and `Y2` are *exclusive*, like the end of a Go slice: the room covers `X1 <= x < X2`.
- `Center` returns **two values**; callers write `x, y := r.Center()`.
- `carve` loops from `X1+1` to `X2-1`. Its lower-case name means *unexported*: visible inside this package only, a signal that it is a helper.

{{diff main.go}}

--- end

%%% Change `carve` to loop from `X1` to `X2` inclusive and place the two rooms so they touch. They become one hall, and later the tunnel logic would dig into a room's wall from the outside. The border ring is load-bearing.
