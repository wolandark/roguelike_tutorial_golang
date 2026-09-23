# Step 46 · Event handlers

### The problem

When the player is dead, movement keys should do nothing and only Escape should work. That is a second *mode* of the game, and there will be more: an inventory menu, an aiming cursor, a title screen. Adding `if dead` checks all over `HandleEvent` scales badly. The alternative is one type per mode, each deciding for itself what a key means and what to draw, and a main loop that only knows "the current mode". The trick that keeps it simple: a mode's `HandleEvent` **returns the mode that should be active next**, usually itself, another to switch, `nil` to quit. Menus later become trivial: they return their parent to close.

>>> 1. In `input.go`, declare an interface `type EventHandler interface { HandleEvent(ev tcell.Event) EventHandler; OnRender(screen tcell.Screen) }`.
>>> 2. Declare a struct `MainGameEventHandler` with a field `Engine *Engine`, and give it `OnRender` (calls `Engine.Render`) and `HandleEvent` (what `Engine.HandleEvent` did; returns a `GameOverEventHandler` when the player is dead, `nil` on Escape, itself otherwise).
>>> 3. Declare a struct `GameOverEventHandler` with a field `Engine *Engine`, and give it `OnRender` and a `HandleEvent` that returns `nil` on Escape or Ctrl-C and itself for everything else.
>>> 4. In `engine.go`, delete `HandleEvent`, and remove `screen.Clear()` and `screen.Show()` from `Render`.
>>> 5. In `main.go`, replace the loop with `var handler EventHandler = &MainGameEventHandler{Engine: engine}` and `for handler != nil { clear; handler.OnRender(screen); show; handler = handler.HandleEvent(screen.PollEvent()) }`.

!!! Plays as before. Let a troll kill you: "You died!", the `@` becomes a `%`, movement keys do nothing, Escape quits.

--- reveal

{{diff input.go}}

- `func (h *MainGameEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }`: a whole method on one line.

{{diff engine.go}}

{{diff main.go}}

- `var handler EventHandler = &MainGameEventHandler{Engine: engine}` declares an interface variable holding a pointer. `for handler != nil` runs until some handler returns `nil`. This exact loop runs the finished game.

--- end

%%% Make `GameOverEventHandler.HandleEvent` return `&MainGameEventHandler{Engine: h.Engine}` on `r` (for "resurrect"). Nothing else changes, and you have a cheat key: switching modes is just returning a different value.
