# Step 66 · Frames

### The problem

A black box on top of the map reads as a glitch rather than a window. A window has a border and a title. Terminals have box-drawing characters for this, `┌ ─ ┐ │ └ ┘`, so a frame is a loop over the rectangle's edge that picks the right character for each cell. Drawing the border is kept separate from clearing the inside, because a border over the map without clearing it will be useful for aiming later.

>>> 1. In `render_functions.go`, add a function `func drawFrame(screen tcell.Screen, x, y, width, height int, title string, style tcell.Style)` that draws `┌ ┐ └ ┘` in the corners, `─` along the top and bottom and `│` along the sides, leaves the inside untouched, and writes `title` centred on the top edge.
>>> 2. In `HistoryViewer.OnRender`, after clearing, draw a frame titled `┤Message history├` around the rectangle, and draw the messages one cell inside it.

!!! Press `v`: the messages sit inside a white frame with a title.

--- reveal

{{diff render_functions.go}}

- The four corners come from a map keyed by `[2]bool{top, left}`: a map whose key is an array. Arrays can be map keys because they can be compared with `==`.
- `len([]rune(title))` counts characters, not bytes: `┤` is three bytes but one cell.

{{diff input.go}}

- The messages start at `4, 4` and are `w-2` wide, so they stay inside the frame.

--- end

%%% Call `drawFrame` *before* `clearRect`. The frame disappears: the clear wipes it. Drawing order is layering, again.
