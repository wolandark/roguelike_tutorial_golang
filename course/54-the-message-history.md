# Step 54 · The message history

### The problem

Five lines of log scroll away fast. Pressing `v` should open the whole history in a window you can scroll, and closing it should return to the game exactly as it was. This is the first *window*, and the event-handler design from step 50 makes it a handler with a parent: it draws the parent first (so the game stays visible behind it), then its own box, and returns the parent on any key it does not use. Every menu from now on follows this shape. Windows need a frame, which is our first box-drawing characters, and clearing the inside is kept separate from drawing the border, because chapter 9 will want a border *over* the map.

>>> 1. In `render_functions.go`, add a function `func drawFrame(screen tcell.Screen, x, y, width, height int, title string, style tcell.Style)` that draws a box border with a centred title, and `func clearRect(screen tcell.Screen, x, y, width, height int, style tcell.Style)`.
>>> 2. In `input.go`, declare a struct `HistoryViewer` with fields `Engine *Engine`, `Parent EventHandler`, `LogLength int` and `Cursor int`, and a constructor function `func NewHistoryViewer(engine *Engine, parent EventHandler) *HistoryViewer`.
>>> 3. Give `HistoryViewer` an `OnRender` that renders the parent, clears a rectangle, draws a frame and the messages up to the cursor.
>>> 4. Give it a `HandleEvent` that moves the cursor with the arrows, PgUp/PgDn, Home/End (wrapping at both ends) and returns `Parent` for any other key.
>>> 5. In `MainGameEventHandler.HandleEvent`, return `NewHistoryViewer(h.Engine, h)` when the rune is `v`.

!!! Fight a bit, press `v`, scroll with the arrows, press any other key to close.

--- reveal

{{diff render_functions.go}}

- The corners come from a small map keyed by `[2]bool{top, left}` (a map with an array key). `len([]rune(title))` counts characters, not bytes, because box characters are three bytes each.

{{diff input.go}}

- The main handler looks at printable characters (`key.Key() == tcell.KeyRune`) before turning keys into actions.
- `Messages[:h.Cursor+1]` is a **slice expression**: everything up to and including the cursor.
- `NewHistoryViewer` is a *constructor function*: Go has no constructors, so a function that builds and returns the value is the convention.

--- end

%%% Remove the `clearRect` call in `OnRender`. The frame is drawn but the map shows through it, with log lines written over the map's colours. That "see-through frame" is exactly what the fireball reticle wants in chapter 9; here it is a bug.
