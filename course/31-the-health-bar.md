# Step 31 · The health bar

`HP: 24/30` as plain text works, but a bar is a roguelike tradition. It is a row of cells with two background colours, filled up to the current fraction, with the numbers written on top.

Create `render_functions.go`:

{{file render_functions.go}}

- Integer arithmetic decides how many cells are filled: 24 of 30 HP on a 20-wide bar fills 16.
- The first loop paints the backgrounds; the second writes the label one cell in, giving each letter the background of the cell it sits on so it stays readable across the boundary.

{{diff engine.go}}

- The bar replaces the text at the bottom left; `engine.go` no longer needs `fmt`, so its import shrinks back to one line. Go refuses to compile with an unused import, which is why the block changes.

!!! Run it: a green bar reading `HP: 30/30` under the map. Take damage and watch it shrink into dark red.
