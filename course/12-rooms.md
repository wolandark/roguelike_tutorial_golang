# Step 12 · Rooms
## Chapter: Generating a dungeon

A dungeon is carved out of rock. From now on a new map is solid wall, and **rooms** are rectangles we carve floor into.

First flip the default in `gamemap.go`:

{{diff gamemap.go}}

Create `procgen.go` (procedural generation):

{{file procgen.go}}

- A room is its two corners. `X2` and `Y2` are *exclusive*, like the end of a Go slice: the room covers `X1 <= x < X2`. `NewRectangularRoom` builds one from a corner and a size.
- `Center` returns **two values**. Go functions can; callers write `x, y := r.Center()`.
- `carve` loops from `X1+1` to `X2-1`: only the *inside* becomes floor. The one-tile ring that stays wall is what keeps two rooms placed side by side from merging into one hall. Its name starts with a lower-case letter, which in Go means it is *unexported*: visible inside this package only. Everything in the game is one package, so this is just a signal that it is a helper.

Then `main.go` carves two rooms by hand and puts the player and the NPC in their centres:

{{diff main.go}}

!!! Run it: two rooms of light blue in a sea of dark blue. The `@` is in the left one and cannot leave it.
