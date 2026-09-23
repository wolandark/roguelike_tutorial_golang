# Step 22 · A room
## Chapter: Generating a dungeon

### In this chapter

A dungeon is carved, not built: start with solid rock, cut rooms into it, join them with corridors, put the player in the first room. This chapter does it in small steps: one room, a second room, a straight line between them, an L-shaped tunnel, then random rooms, rooms that do not overlap, and finally every room connected. All of it is arithmetic on rectangles; the new Go is multiple return values and random numbers.

### The problem

So far the map starts as floor and we place walls by hand. A dungeon works the other way round: everything starts as **rock**, and open space is cut out of it. A room is the simplest open space, a rectangle, and to carve one we need a way to describe a rectangle and a way to turn its cells into floor.

One detail decides how the dungeon will look. If two rooms are carved right next to each other they merge into one big hall, unless every room keeps a one-cell ring of wall around itself. So a room carves only its *inside*.

>>> 1. In `gamemap.go`, change `NewGameMap` to fill the map with `wall` instead of `floor`.
>>> 2. Create `procgen.go` and declare a struct `RectangularRoom` with fields `X1, Y1, X2, Y2 int`, and a function `func NewRectangularRoom(x, y, width, height int) RectangularRoom` that returns the room with its top-left corner at `x, y` and its far corner at `x+width, y+height`.
>>> 3. Add a method `func (r RectangularRoom) carve(m *GameMap)` that sets every cell from `X1+1` up to `X2-1`, and from `Y1+1` up to `Y2-1`, to `floor`.
>>> 4. In `main.go`, delete the loop that placed the three wall cells, and instead create `room := NewRectangularRoom(30, 20, 20, 10)` and call `room.carve(gameMap)`.

!!! One light-blue room in a sea of dark blue rock, with both `@`s inside it. You cannot walk out of it.

--- reveal

{{diff gamemap.go}}

{{file procgen.go}}

- The far corner, `X2, Y2`, is the first cell *outside* the rectangle, the way the end index of a Go slice is the first element outside it: `s[2:5]` holds elements 2, 3 and 4.
- `carve` starts at `X1+1` and stops before `X2`, so the outermost cells on every side stay wall. That is the ring that keeps rooms apart.
- `carve` starts with a lower-case letter, so it is *unexported*: only code in this package can call it. Everything here is one package, so it is simply a signal that this is a helper.
- The receiver is a value, `(r RectangularRoom)`: carving changes the map, not the room.

{{diff main.go}}

--- end

%%% Change the loops in `carve` to start at `X1` and `Y1` and to include `X2` and `Y2`. The room grows by one cell on every side and its border wall disappears. Now place a second room right next to it in `main.go`: the two merge into one hall.
