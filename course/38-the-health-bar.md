# Step 38 · The health bar

### The problem

`HP: 24/30` works, but a bar you can read at a glance is a roguelike tradition. There is no widget library; a bar is a row of cells with two background colours, filled up to the current fraction, with the numbers written on top so they stay readable across the boundary.

>>> 1. Create `render_functions.go` with a function `func renderBar(screen tcell.Screen, x, y, current, maximum, totalWidth int)`: fill `current*totalWidth/maximum` cells with `colorBarFilled`, the rest with `colorBarEmpty`, and write `HP: cur/max` on top from the second cell.
>>> 2. In `engine.go`, replace the HP text with `renderBar(screen, 0, 45, hp, maxHP, 20)` and drop the now unused `fmt` import.

!!! A green bar reading `HP: 30/30` under the map. Take damage and watch it shrink into dark red.

--- reveal

{{file render_functions.go}}

- Integer arithmetic decides the fill: 24 of 30 on a 20-wide bar fills 16 cells.

{{diff engine.go}}

- `engine.go` no longer needs `fmt`, so its import shrinks back to one line. Go refuses to compile with an unused import.

--- end

%%% Compute `filled` as `current / maximum * totalWidth`. With integers, `24 / 30` is 0 and the bar is always empty. Order of operations in integer arithmetic matters; multiply before you divide.
