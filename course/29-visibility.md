# Step 29 · Visibility
## Chapter: Field of view

### In this chapter

Seeing the whole dungeon spoils exploration. Roguelikes show three kinds of cell: what you can see now, what you have seen before, and nothing at all. This chapter builds that in small steps: first a record of which cells are visible, then a brighter look for them, then a memory of cells seen before, and then the real line-of-sight algorithm in two parts. The last step hides monsters the same way.

### The problem

We need to know, for every cell, whether the player can see it this turn. That fact changes every turn and belongs to the *cell*, not to the kind of tile: two floor cells can be one visible and one not. So it cannot live in `Tile`. The map gets a second grid shaped like `Tiles`, of booleans: `Visible[y][x]`.

For now the rule for "visible" is deliberately simple: every cell within 8 cells of the player, walls or not. That lets us build the drawing side and test it before the real algorithm arrives. Cells that are not visible are drawn black.

>>> 1. In `gamemap.go`, add a field `Visible [][]bool` to `GameMap`, allocate its rows in `NewGameMap` next to the rows of `Tiles`, and add a method `func (m *GameMap) IsVisible(x, y int) bool` that returns `false` outside the map and the stored value otherwise.
>>> 2. Add a method `func (m *GameMap) ComputeFOV(ox, oy, radius int)` that first sets every cell of `Visible` to `false`, then sets to `true` every in-bounds cell whose `x` is within `radius` of `ox` and whose `y` is within `radius` of `oy`.
>>> 3. In `tiles.go`, add a package variable `var shroud = Glyph{' ', tcell.ColorWhite, tcell.ColorBlack}`, and in `Render` draw `shroud` for every cell that is not visible.
>>> 4. In `engine.go`, add a method `func (e *Engine) UpdateFOV()` that calls `e.GameMap.ComputeFOV(e.Player.X, e.Player.Y, 8)`, call it at the end of `HandleEvent`, and in `main.go` call `engine.UpdateFOV()` once before the loop.

!!! A square of map around you, 17 by 17 cells, and black everywhere else. The square follows you.

--- reveal

{{diff gamemap.go}}

- The struct literal is now written one field per line, because it no longer fits on one.
- `for _, row := range m.Visible { clear(row) }` resets every row to `false`. `clear` is a built-in (Go 1.21 or newer) that sets every element of a slice to its zero value. `row` is a copy of the row's slice header, but it points at the same cells, so clearing it clears the grid.
- `IsVisible` checks `InBounds` first for the same reason movement does: the square around a player near the edge reaches outside the map.

{{diff tiles.go}}

{{diff engine.go}}

{{diff main.go}}

- Without the call in `main.go`, the first frame would be all black until the first key press.

--- end

%%% Remove the `clear(row)` loop from `ComputeFOV`. Everything you have ever been near stays visible: visibility has to be computed from scratch every turn.
