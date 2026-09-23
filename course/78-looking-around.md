# Step 78 · Looking around

### The problem

Confusion and fireball scrolls need the player to choose a cell. That is a mode: a highlighted cursor moves over the map with the movement keys, Enter confirms, any other key cancels. What happens with the chosen cell differs per use: look at it, or throw a scroll at it. So the handler takes a **callback**, a function value that receives the chosen cell and returns the next handler.

The cursor can reuse the mouse position from step 64, `Engine.MouseX` and `Engine.MouseY`: the names line already shows what is under it. Built once, this mode gives a `/` look command for free, which is what this step adds. The scrolls come later in the chapter.

>>> 1. In `input.go`, declare a struct `SelectIndexHandler` with the fields `Engine *Engine` and `OnSelect func(x, y int) EventHandler`, and a constructor `NewSelectIndexHandler(engine *Engine, onSelect func(x, y int) EventHandler) *SelectIndexHandler` that first moves `engine.MouseX, engine.MouseY` to the player.
>>> 2. Give it an `OnRender` that renders the game, reads the cursor cell back with `screen.GetContent` and sets it again with `style.Reverse(true)`.
>>> 3. Give it a `HandleEvent` that ignores non-key events, moves the cursor one cell for a key found in `moveKeys` or `moveRunes` (kept inside the map), calls `h.OnSelect` with the cursor on Enter, and returns `&MainGameEventHandler{Engine: h.Engine}` for any other key.
>>> 4. Add a function `func NewLookHandler(engine *Engine) *SelectIndexHandler` whose callback returns `&MainGameEventHandler{Engine: engine}`, and a `case '/'` in `MainGameEventHandler.HandleEvent` that returns it.

!!! Press `/`: the `@` cell is highlighted. Move the highlight with the arrows or `hjkl` and read the names line as it passes over monsters and potions. Enter or Escape returns to the game.

--- reveal

{{diff input.go}}

- `screen.GetContent(x, y)` returns what is drawn in a cell: the rune, extra combining runes, the style and the width. We keep the rune and the style, and draw the cell again with the colours swapped.
- `NewLookHandler` passes a **function literal**, `func(x, y int) EventHandler { ... }`, as the callback. It uses `engine` from the surrounding function, which is allowed: a function literal can read the variables around it.
- `max(0, min(..., Width-1))` keeps the cursor inside the map.

--- end

%%% In `NewLookHandler`, make the callback return `nil` instead of the main game handler. Press `/` and then Enter: the game quits, because a `nil` handler ends the main loop (step 56). The callback decides what comes next.
