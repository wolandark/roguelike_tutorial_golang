# Step 48 · Your hit points

### The problem

Monsters will hit back soon, and the player needs to see their health. For now a line of text under the map is enough: `HP: 30/30`. The messages move to the right to make room for it.

>>> 1. In `engine.go`, in `Render`, draw `HP: <hp>/<max hp>` of the player at column 1, row 46.
>>> 2. Draw the messages starting at column 21 instead of column 0.

!!! `HP: 30/30` at the bottom left, the messages to its right.

--- reveal

{{diff engine.go}}

- `fmt.Sprintf("HP: %d/%d", ...)` builds the text; it is then drawn rune by rune like the messages.

--- end

%%% Draw the HP text at row 45 instead of 46. It collides with the first message line and the two overwrite each other character by character: whatever is drawn last wins, just as with entities over tiles.
