# Step 10 · Tiles and the map

The world needs floor and walls. A **tile** is what one cell of the map is made of; the **map** is a grid of them. Two new files, no behaviour change yet: the map is drawn, but nothing stops you walking into a wall.

Create `tiles.go`:

{{file tiles.go}}

- A `Glyph` is how something looks: a character plus foreground and background colours. `Style()` is a method with a *value* receiver: it only reads `g`, so no pointer is needed. It turns the two colours into the `tcell.Style` that `SetContent` wants.
- `Tile` knows whether you can stand on it, whether you can see through it (used in chapter 4) and how it looks. Both tiles draw a **space with a coloured background**; floor is a lighter blue than wall.
- `tcell.NewRGBColor(50, 50, 150)` is a true-colour value. tcell approximates it on terminals with only 256 or 16 colours.
- `var ( ... )` declares package-level variables. `floor` and `wall` are *values*; the map copies them into its cells, so changing a cell never changes these.

Create `gamemap.go`:

{{file gamemap.go}}

- The tiles live in **one flat slice**, row after row. The tile at `(x, y)` is `Tiles[y*Width + x]`. `TileAt` and `SetTile` hide that arithmetic; nothing else touches the slice directly. This is the Go equivalent of the 2D numpy array in the Python tutorial.
- `make([]Tile, width*height)` allocates the slice; every element starts as the zero `Tile` (all `false`, blank glyph), so `NewGameMap` fills it with `floor`.
- `InBounds` must be checked before `TileAt`: indexing outside the slice panics, and a negative `x` would silently land on the previous row.
- `Render` is a double loop, one `SetContent` per cell, using the tile's `Dark` glyph (the only look we have so far).

And draw it in `main.go`:

{{diff main.go}}

- Two more constants: the map is 80 wide but only 45 tall, leaving five rows at the bottom for a status area later.
- Three tiles at row 22 become wall, something to bump into next step.

!!! Run it: the whole screen is blue floor with a darker three-cell wall in the middle. You can still walk through it.
