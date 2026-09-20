# Step 29 · Event handlers

### The problem

When the player is dead, movement keys should do nothing and only Escape should work. That is a second *mode* of the game, and there will be more: an inventory menu, an aiming cursor, a title screen. Adding `if dead` checks all over `HandleEvent` scales badly. The alternative is one type per mode, each deciding for itself what a key means and what to draw, and a main loop that only knows "the current mode". The trick that keeps it simple: a mode's `HandleEvent` **returns the mode that should be active next**, usually itself, another to switch, `nil` to quit. Menus later become trivial: they return their parent to close.

>>> Define an `EventHandler` interface with `HandleEvent(ev tcell.Event) EventHandler` and `OnRender(screen)`. Move the engine's event handling into a `MainGameEventHandler` (with the engine as a field) that returns a `GameOverEventHandler` when the player is dead; the game-over handler ignores everything but Escape. `Engine.Render` should no longer clear or show. The loop in `main.go` becomes: clear, current handler renders, show, handler = handler.HandleEvent(next event), until nil.

!!! Plays as before. Let a troll kill you: "You died!", the `@` becomes a `%`, movement keys do nothing, Escape quits.

--- reveal

{{diff input.go}}

- `func (h *MainGameEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }`: a whole method on one line.

{{diff engine.go}}

{{diff main.go}}

- `var handler EventHandler = &MainGameEventHandler{Engine: engine}` declares an interface variable holding a pointer. `for handler != nil` runs until some handler returns `nil`. This exact loop runs the finished game.

--- end

%%% Make `GameOverEventHandler.HandleEvent` return `&MainGameEventHandler{Engine: h.Engine}` on `r` (for "resurrect"). Nothing else changes, and you have a cheat key: switching modes is just returning a different value.
