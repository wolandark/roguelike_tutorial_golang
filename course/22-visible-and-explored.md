# Step 22 · Visible and explored
## Chapter: Field of view

### In this chapter

Seeing the whole dungeon spoils exploration. Roguelikes show three kinds of tile: what you can see now (lit), what you have seen before (dim, remembered), and nothing at all. The bookkeeping for that is simple and comes first; the real line-of-sight algorithm, shadowcasting, replaces a naive rule in the second step; the third step hides entities the same way.

### The problem

We need two facts per tile that change every turn, *visible* and *explored*, and three looks per tile. Where do the facts live? They could be fields on `Tile`, but tiles are copied values shared between cells, and visibility belongs to the *cell*, not the kind of tile. So the map gets two parallel slices of booleans, indexed exactly like `Tiles`. For this step the rule for "visible" is deliberately naive, everything within 8 cells, so that the drawing side can be built and tested before the algorithm arrives.

>>> 1. In `tiles.go`, add a field `Light Glyph` to `Tile`, give `floor` and `wall` a lit glyph, and declare `var shroud = Glyph{' ', tcell.ColorWhite, tcell.ColorBlack}`.
>>> 2. In `gamemap.go`, add the fields `Visible []bool` and `Explored []bool` to `GameMap` and allocate them in `NewGameMap`.
>>> 3. Add the methods `IsVisible(x, y int) bool` and `IsExplored(x, y int) bool`, each checking `InBounds` first.
>>> 4. Add a method `func (m *GameMap) setVisible(x, y int)` that marks a cell both visible and explored.
>>> 5. Add a method `func (m *GameMap) ComputeFOV(ox, oy, radius int)` that clears `Visible` and marks every cell within the square of `radius` around `ox, oy`.
>>> 6. Change `Render` to draw the `Light` glyph for visible cells, `Dark` for explored ones and `shroud` otherwise.
>>> 7. In `engine.go`, add a method `func (e *Engine) UpdateFOV()` that calls `ComputeFOV(e.Player.X, e.Player.Y, 8)`, and call it at the end of `HandleEvent`. In `main.go`, call `engine.UpdateFOV()` once before the loop.

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
