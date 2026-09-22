# Step 36 · The health bar

### The problem

`HP: 24/30` works, but a bar you can read at a glance is a roguelike tradition. There is no widget library; a bar is a row of cells with two background colours, filled up to the current fraction, with the numbers written on top so they stay readable across the boundary.

>>> Create `render_functions.go` with `renderBar(screen, x, y, current, maximum, totalWidth)`: `filled = current*totalWidth/maximum` cells with the filled background, the rest with the empty one, and the label `HP: cur/max` written from the second cell using each cell's own background. Replace the HP text with a 20-wide bar at (0, 45).

!!! A green bar reading `HP: 30/30` under the map. Take damage and watch it shrink into dark red.

--- reveal

{{file render_functions.go}}

- Integer arithmetic decides the fill: 24 of 30 on a 20-wide bar fills 16 cells.

{{diff engine.go}}

- `engine.go` no longer needs `fmt`, so its import shrinks back to one line. Go refuses to compile with an unused import.

--- end

%%% Compute `filled` as `current / maximum * totalWidth`. With integers, `24 / 30` is 0 and the bar is always empty. Order of operations in integer arithmetic matters; multiply before you divide.
