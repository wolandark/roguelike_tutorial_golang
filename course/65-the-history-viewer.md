# Step 65 · The history viewer

### The problem

Five lines of log scroll away fast. Pressing `v` should show the whole history over the game, and any key should return to the game exactly as it was. That is a new mode, so it is a new event handler, and it is the first handler that draws *on top of* another one: it keeps a pointer to the handler it was opened from, its **parent**, draws the parent first so the game stays visible, then clears a rectangle and draws the messages into it. Closing it is just returning the parent.

>>> 1. In `render_functions.go`, add a function `func clearRect(screen tcell.Screen, x, y, width, height int, style tcell.Style)` that sets every cell of the rectangle to a space.
>>> 2. In `input.go`, declare a struct `HistoryViewer` with fields `Engine *Engine` and `Parent EventHandler`, and a function `func NewHistoryViewer(engine *Engine, parent EventHandler) *HistoryViewer` that returns one.
>>> 3. Give `HistoryViewer` an `OnRender` that calls `h.Parent.OnRender(screen)`, clears a rectangle 3 cells in from every edge of the screen, and draws all messages into it with `renderMessages`; and a `HandleEvent` that returns `h.Parent` for any key.
>>> 4. In `MainGameEventHandler.HandleEvent`, return `NewHistoryViewer(h.Engine, h)` when the key is the rune `v`.

!!! Press `v`: the game is covered by a black box full of messages, the newest at the bottom. Any key brings the game back.

--- reveal

{{diff render_functions.go}}

{{diff input.go}}

- `NewHistoryViewer` is a **constructor function**. Go has no constructors; a function that builds and returns the value is the convention.
- `h.Parent.OnRender(screen)` draws the game first, then the viewer draws over it. That works because clearing and showing the screen happen in the main loop, once per frame, since step 56.
- The main handler checks for `v` before turning keys into actions, because `v` is not an action in the game.

--- end

%%% Remove `h.Parent.OnRender(screen)`. The box is drawn on an otherwise black screen: drawing the parent is what makes the viewer look like a window over the game.
