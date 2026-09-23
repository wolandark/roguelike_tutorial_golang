# Step 13 · A floor

### The problem

The `@` stands on nothing. A dungeon floor is, at its simplest, a rectangle of cells drawn before the entities so they appear on top of it. No new types yet: two nested loops and `SetContent`, the same call that draws the `@`.

>>> 1. In the loop in `main.go`, right after `screen.Clear()` and before the entities are drawn, add two nested loops: `y` from 20 to 29, `x` from 30 to 49.
>>> 2. Inside them, draw a floor cell: `screen.SetContent(x, y, '.', nil, tcell.StyleDefault)`.

!!! A block of dots below the `@`. Walk onto it: the `@` is drawn over the dots, because entities are drawn after the floor.

--- reveal

{{diff main.go}}

- The outer loop is rows, the inner one columns: `y` then `x`. Every cell is one `SetContent`.
- The order of drawing is the order of layering: whatever is drawn later covers what was drawn earlier in the same cell. That is why the floor goes first and the entities after.

--- end

%%% Move the floor loops *after* the entity loop. Walk onto the dots: the `@` disappears under them. Layering is nothing more than call order.
