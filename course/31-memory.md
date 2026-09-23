# Step 31 · Memory

### The problem

Cells you walked past vanish into black the moment they leave your view. Roguelikes remember them: a cell once seen stays on screen, dimmed, so you can find your way back. That is a second fact per cell, **explored**, which never goes back to `false`. It is set whenever a cell becomes visible, so the natural place to set it is the same place that sets `Visible`.

>>> 1. In `gamemap.go`, add a field `Explored [][]bool` to `GameMap`, allocate its rows in `NewGameMap`, and add a method `IsExplored(x, y int) bool` like `IsVisible`.
>>> 2. Add a method `func (m *GameMap) setVisible(x, y int)` that, for an in-bounds cell, sets both `Visible` and `Explored` to `true`, and call it from `ComputeFOV` instead of setting `Visible` directly.
>>> 3. In `Render`, draw the `Light` glyph for visible cells, the `Dark` glyph for cells that are explored but not visible, and `shroud` for the rest.

!!! A lit square around you, dark blue where you have been, black elsewhere. Walls still do not block your view.

--- reveal

{{diff gamemap.go}}

- `setVisible` marks a cell visible **and** explored. It is the only place that writes `Explored`, so memory can never get out of sync with sight.
- `ComputeFOV` clears `Visible` every turn but never touches `Explored`: that is the whole difference between sight and memory.
- `Render` uses a tagless `switch`: the first `case` whose condition is true wins, so visible beats explored beats shroud.

--- end

%%% Swap the two cases in `Render` so `IsExplored` is tested first. Every visible cell is also explored, so the first case always wins and nothing is ever drawn lit. The order of cases in a tagless switch is part of its meaning.
