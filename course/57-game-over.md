# Step 57 · Game over

### The problem

Now the second mode. When the player dies, the game should switch to a mode in which only Escape and Ctrl-C do anything. With event handlers that is a new type, and one `if` in the main handler that returns it.

>>> 1. In `input.go`, declare a struct `GameOverEventHandler` with a field `Engine *Engine`. Give it an `OnRender` that calls `h.Engine.Render(screen)`, and a `HandleEvent` that returns `nil` for Escape or Ctrl-C and `h` for everything else.
>>> 2. In `MainGameEventHandler.HandleEvent`, after the enemy turns and the field of view, return `&GameOverEventHandler{Engine: h.Engine}` when the player is no longer alive.

!!! Let a troll kill you: "You died!", the `@` becomes a `%`, movement keys do nothing, Escape quits.

--- reveal

{{diff input.go}}

- The game-over mode draws the same picture as the game, so its `OnRender` also calls `Engine.Render`. Only its `HandleEvent` differs.

--- end

%%% Make `GameOverEventHandler.HandleEvent` return `&MainGameEventHandler{Engine: h.Engine}` on `r` (for "resurrect"). Nothing else changes, and you have a cheat key: switching modes is just returning a different value.
