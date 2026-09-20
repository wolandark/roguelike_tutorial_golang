# Step 33 · The message history

Five lines of log scroll away fast. Pressing `v` opens the whole history in a window you can scroll. Windows need frames, and frames are our first box-drawing characters.

{{diff render_functions.go}}

- `drawFrame` draws only the **border**: corners from a small map keyed by `[2]bool{top, left}` (a map with an array key), `─` on the horizontal edges, `│` on the vertical ones, and a centred title on the top edge. `len([]rune(title))` counts characters, not bytes, because box characters are three bytes each.
- `clearRect` blanks a rectangle. It is separate from the frame so that chapter 9 can draw a frame *over* the map without hiding what is inside it.

{{diff input.go}}

- The main handler now looks at printable characters (`key.Key() == tcell.KeyRune`) before turning keys into actions, and `v` switches to a `HistoryViewer`.
- `HistoryViewer` keeps a `Parent` handler and a cursor. Its `OnRender` first calls `Parent.OnRender`, so the game stays visible behind the window, then clears a rectangle, draws the frame and renders the log up to the cursor with a **slice expression** `Messages[:h.Cursor+1]` (everything up to and including the cursor).
- `HandleEvent` moves the cursor: arrows by one, PgUp/PgDn by ten, Home/End jump, and the cursor wraps around at both ends. **Any other key returns `Parent`**, which is how the window closes. Every menu from now on follows this shape.
- `NewHistoryViewer` is a *constructor function*: Go has no constructors, so a function that builds and returns the value is the convention.

!!! Run it: fight a bit, press `v`, scroll with the arrows, press any other key to close.
