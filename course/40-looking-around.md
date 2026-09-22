# Step 40 · Looking around

### The problem

Confusion and fireball need the player to choose a cell. That is a mode: a highlight moves over the map with the keys or the mouse, Enter or a click confirms, anything else cancels. What happens with the chosen cell differs per use (look at it, throw a scroll at it), so the handler takes a **callback**, a function value that receives the cell and returns the next handler. Built once, the mode gives a `/` look command for free, which is the version this step adds; the scrolls come next.

>>> Add a `SelectIndexHandler` with `Engine` and `OnSelect func(x, y int) EventHandler`. The cursor is `Engine.MouseX/MouseY`, starting on the player. `OnRender` draws the game and flips the cursor cell to reverse video (read it back with `GetContent`). `HandleEvent`: mouse moves the cursor and a left click selects; movement keys move it (times 5, 10, 20 with Shift, Ctrl, Alt), clamped to the map; Enter selects; anything else returns to the main game. Add `NewLookHandler` whose callback just returns to the game, bound to `/`.

!!! Press `/`, move the highlight with the arrows or `hjkl` (try Shift-arrow), and read the names line as it passes over monsters. Enter or Escape returns.

--- reveal

{{diff input.go}}

- `screen.GetContent` returns the rune, combining runes, style and width of a cell; we keep the rune and style and rewrite the cell with `style.Reverse(true)`.
- `e.Buttons()&tcell.Button1 != 0` and `e.Modifiers()&tcell.ModShift != 0` are bit tests on bit sets.
- `NewLookHandler` passes a **function literal** as the callback; the names line from step 33 does the rest.

--- end

%%% Make the callback in `NewLookHandler` return `h` (the select handler itself) instead of the main game. Enter no longer leaves look mode; only a non-movement key does. The callback decides what "select" means, including "keep selecting".
