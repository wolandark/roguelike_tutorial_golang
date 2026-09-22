# Step 14 · The map

### The problem

Drawing tiles by hand does not scale to 80 by 45 of them. The map is a grid of tiles, and the question is how to store a grid in Go. The Python tutorial uses a 2D numpy array. Go has no numpy, and a slice of slices is awkward to allocate and slow to copy. The idiomatic choice is **one flat slice** with a little arithmetic: the tile at `(x, y)` sits at index `y*width + x`. This step creates the map, fills it with floor and draws it; the arithmetic appears twice for now, and the next step puts it in one place.

>>> Create `gamemap.go` with a `GameMap` holding `Width`, `Height` and a `Tiles []Tile`; `NewGameMap(width, height)` that makes the slice and fills it with `floor`; and a `Render(screen)` that draws every cell from its `Dark` glyph. In `main.go`, add `mapWidth = 80` and `mapHeight = 45`, make a map, and draw it before the entities instead of the strip.

!!! The whole screen is blue floor, except the bottom five rows, and the two `@`s stand on it.

--- reveal

{{file gamemap.go}}

- `make([]Tile, width*height)` allocates the slice; every element starts as the zero `Tile` (all `false`, blank glyph), so `NewGameMap` fills it with `floor` in a loop. `for i := range m.Tiles` with one variable gives the index.
- `m.Tiles[y*m.Width+x]` is the flat-slice arithmetic: row `y` starts at `y*Width`.

{{diff main.go}}

- The map is 80 wide but only 45 tall, leaving five rows at the bottom for a status area later.

--- end

%%% Swap the two loops in `Render` so `x` is the outer one. Nothing changes on screen, because each cell is still drawn once; but with the index formula `y*Width + x`, the inner loop of the original visits memory in order, which is why the row loop goes outside. It does not matter at this size. It is the habit that matters.
