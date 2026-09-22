# Step 13 · A floor

### The problem

The `@` stands on nothing. A dungeon floor is, at its simplest, a rectangle of cells drawn before the entities so they appear on top of it. No new types yet: two nested loops and `SetContent`, the same call that draws the `@`.

>>> Before the entities are drawn, fill a rectangle from column 30 to 49 and row 20 to 29 with `.` cells.

!!! A block of dots below the `@`. Walk onto it: the `@` is drawn over the dots, because entities are drawn after the floor.

--- reveal

{{diff main.go}}

- The outer loop is rows, the inner one columns: `y` then `x`. Every cell is one `SetContent`.
- The order of drawing is the order of layering: whatever is drawn later covers what was drawn earlier in the same cell. That is why the floor goes first and the entities after.

--- end

%%% Move the floor loops *after* the entity loop. Walk onto the dots: the `@` disappears under them. Layering is nothing more than call order.
