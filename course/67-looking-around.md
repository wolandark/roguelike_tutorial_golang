# Step 67 · Looking around

### The problem

Confusion and fireball need the player to choose a cell. That is a mode: a highlight moves over the map with the keys or the mouse, Enter or a click confirms, anything else cancels. What happens with the chosen cell differs per use (look at it, throw a scroll at it), so the handler takes a **callback**, a function value that receives the cell and returns the next handler. Built once, the mode gives a `/` look command for free, which is the version this step adds; the scrolls come next.

>>> 1. In `input.go`, declare a struct `SelectIndexHandler` with fields `Engine *Engine` and `OnSelect func(x, y int) EventHandler`, and a constructor `NewSelectIndexHandler` that puts the cursor (`Engine.MouseX/MouseY`) on the player.
>>> 2. Give it an `OnRender` that renders the game and redraws the cursor cell with `style.Reverse(true)`.
>>> 3. Give it a `HandleEvent` that moves the cursor with the mouse and the movement keys (x5, x10, x20 with Shift, Ctrl, Alt), calls `OnSelect` on Enter or a left click, and returns to the main game on any other key.
>>> 4. Add a function `func NewLookHandler(engine *Engine) *SelectIndexHandler` whose callback returns to the main game, and open it for `/`.

!!! Press `/`, move the highlight with the arrows or `hjkl` (try Shift-arrow), and read the names line as it passes over monsters. Enter or Escape returns.

--- reveal

{{diff input.go}}

- `screen.GetContent` returns the rune, combining runes, style and width of a cell; we keep the rune and style and rewrite the cell with `style.Reverse(true)`.
- `e.Buttons()&tcell.Button1 != 0` and `e.Modifiers()&tcell.ModShift != 0` are bit tests on bit sets.
- `NewLookHandler` passes a **function literal** as the callback; the names line from step 60 does the rest.

--- end

%%% Make the callback in `NewLookHandler` return `h` (the select handler itself) instead of the main game. Enter no longer leaves look mode; only a non-movement key does. The callback decides what "select" means, including "keep selecting".
