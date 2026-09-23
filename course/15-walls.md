# Step 15 · Walls

### The problem

A floor with no walls is a field. Walls need a second look, and they will need two more facts the floor also has: can you stand on it, can you see through it (chapter 5). A look plus those two facts is a **tile**, and floor and wall are the first two kinds. This step only draws them; the map that stores them is next.

>>> 1. In `tiles.go`, declare a struct `Tile` with fields `Walkable bool`, `Transparent bool` and `Dark Glyph`.
>>> 2. Replace `floorGlyph` with two package variables in a `var ( )` block: `floor` (walkable, transparent, a space on RGB 50, 50, 150) and `wall` (neither, a space on RGB 0, 0, 100).
>>> 3. In `main.go`, inside the floor loops, pick `g := floor.Dark`, and switch to `wall.Dark` when the cell is on the edge of the rectangle (`x == 30 || x == 49 || y == 20 || y == 29`). Draw `g`.

!!! A blue room with a darker border. You can still walk through the border; nothing checks it yet.

--- reveal

{{diff tiles.go}}

- `var ( ... )` declares several package-level variables in one block. `floor` and `wall` are *values*; anything that copies them gets its own copy, so a map cell can never change these.

{{diff main.go}}

- The `if` picks the wall glyph on the four edges of the rectangle; everything inside stays floor.

--- end

%%% Give the wall a glyph of `'#'` with a white foreground instead of a space. Walls now read as walls without any colour; most terminal roguelikes draw them that way.
