# Step 16 · Visible and explored
## Chapter: Field of view

### In this chapter

Seeing the whole dungeon spoils exploration. Roguelikes show three kinds of tile: what you can see now (lit), what you have seen before (dim, remembered), and nothing at all. The bookkeeping for that is simple and comes first; the real line-of-sight algorithm, shadowcasting, replaces a naive rule in the second step; the third step hides entities the same way.

### The problem

We need two facts per tile that change every turn, *visible* and *explored*, and three looks per tile. Where do the facts live? They could be fields on `Tile`, but tiles are copied values shared between cells, and visibility belongs to the *cell*, not the kind of tile. So the map gets two parallel slices of booleans, indexed exactly like `Tiles`. For this step the rule for "visible" is deliberately naive, everything within 8 cells, so that the drawing side can be built and tested before the algorithm arrives.

>>> Give `Tile` a `Light` glyph and add a black `shroud` glyph for unknown cells. Give `GameMap` `Visible` and `Explored` slices, `IsVisible`/`IsExplored` helpers that also check bounds, a `setVisible` helper that marks both, and a `ComputeFOV(ox, oy, radius)` that clears `Visible` and marks every cell within a square of `radius`. `Render` picks lit, dark or shroud. The engine recomputes after every action and once before the first frame.

!!! A lit square around you, dark blue where you have been, black elsewhere. Walls do not block your view yet.

--- reveal

{{diff tiles.go}}

{{diff gamemap.go}}

- `setVisible` marks a tile visible **and** explored. It is the only place that writes `Explored`, so memory can never get out of sync with sight.
- `for i := range m.Visible` with only an index is the idiom for "every position of the slice".
- `Render` uses a **tagless `switch`**: the first `case` whose condition is true wins, so visible beats explored beats shroud.

{{diff engine.go}}

{{diff main.go}}

- Without the call in `main.go`, the first frame would be black until the first key press.

--- end

%%% Remove the `m.Visible[i] = false` loop from `ComputeFOV`. Everything you have ever seen stays lit: visibility must be recomputed from scratch each turn, memory is what persists.
