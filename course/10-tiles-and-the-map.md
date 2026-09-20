# Step 10 · Tiles and the map

### The problem

The world is empty. We want floor to walk on and walls to bump into, drawn as coloured cells. Two design choices to make. First, what a **tile** knows: whether you can walk on it, whether you can see through it (chapter 4 will need that), and how it looks. Second, how to store an 80 by 45 grid of them. The Python tutorial uses a 2D numpy array; Go has no numpy, and a slice of slices is clumsy to allocate and slow to copy. One flat slice with a little index arithmetic is the idiomatic answer, as long as the arithmetic lives in exactly one place.

>>> Create `tiles.go` with a `Glyph` (rune, foreground, background) that can turn itself into a `tcell.Style`, a `Tile` (walkable, transparent, glyph) and two package-level tiles, `floor` and `wall`, drawn as spaces with different blue backgrounds. Create `gamemap.go` with a `GameMap` holding the tiles in one slice, `NewGameMap` filling it with floor, `InBounds`, `TileAt`, `SetTile` and a `Render` that draws every cell. In `main.go`, make a map with a few wall tiles and draw it before the entities.

!!! The whole screen is blue floor with a darker three-cell wall in the middle. You can still walk through it.

--- reveal

{{file tiles.go}}

- `Style()` has a *value* receiver: it only reads `g`, so no pointer is needed. It turns the two colours into the `tcell.Style` that `SetContent` wants.
- `tcell.NewRGBColor(50, 50, 150)` is a true-colour value; tcell approximates it on terminals with fewer colours.
- `var ( ... )` declares package-level variables. `floor` and `wall` are *values*; the map copies them into its cells, so changing a cell never changes these.

{{file gamemap.go}}

- The tile at `(x, y)` is `Tiles[y*Width + x]`. `TileAt` and `SetTile` hide that; nothing else ever touches the slice directly.
- `make([]Tile, width*height)` allocates the slice; every element starts as the zero `Tile`, so `NewGameMap` fills it with `floor`.
- `InBounds` must be checked before `TileAt`: an index outside the slice panics, and a negative `x` would silently land on the previous row.

{{diff main.go}}

- The map is 80 wide but only 45 tall, leaving five rows at the bottom for a status area later.

--- end

%%% Ask for `gameMap.TileAt(-1, 0)` somewhere and run: no panic, you silently get the last tile of the previous row, because `-1 + 0*80` is `-1`... which *does* panic. Now try `TileAt(-1, 1)`: index 79, the end of row 0, no panic, wrong tile. That is why `InBounds` exists and why step 11 checks it first.
