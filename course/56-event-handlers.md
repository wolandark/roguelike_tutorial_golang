# Step 56 · Event handlers

### The problem

When the player is dead, movement keys should do nothing and only Escape should work. That is a different *mode* of the game, and more will follow: an inventory menu, an aiming cursor, a title screen. Adding `if dead` checks all over the input code scales badly. The alternative is one type per mode, each deciding for itself what a key means and what to draw, and a main loop that only knows "the current mode".

The idea that keeps this simple: a mode's `HandleEvent` **returns the mode that should be active next**. Usually that is itself; returning another mode switches to it; returning `nil` quits. This step introduces the idea with the one mode we have. The next step adds game over.

>>> 1. In `input.go`, declare an interface `type EventHandler interface { HandleEvent(ev tcell.Event) EventHandler; OnRender(screen tcell.Screen) }`.
>>> 2. Declare a struct `MainGameEventHandler` with a field `Engine *Engine`. Give it `OnRender`, which calls `h.Engine.Render(screen)`, and `HandleEvent`, which does what `Engine.HandleEvent` did but returns `nil` for Escape and `h` otherwise.
>>> 3. In `engine.go`, delete `HandleEvent`, and remove `screen.Clear()` and `screen.Show()` from `Render`.
>>> 4. In `main.go`, replace the loop with: `var handler EventHandler = &MainGameEventHandler{Engine: engine}`, then `for handler != nil` clear the screen, call `handler.OnRender(screen)`, show it, and set `handler = handler.HandleEvent(screen.PollEvent())`. Return `nil` after the loop.

!!! Plays exactly as before.

--- reveal

{{diff input.go}}

- `func (h *MainGameEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }` is a whole method on one line.
- `HandleEvent` returns `h`, the handler itself, to stay in this mode.

{{diff engine.go}}

- Clearing and showing move to the loop in `main.go`, so that a mode can draw on top of another mode's picture later.

{{diff main.go}}

- `var handler EventHandler = ...` declares a variable of an interface type holding a pointer. `for handler != nil` runs until some handler returns `nil`. This loop runs the finished game.

--- end

%%% Make `HandleEvent` return `nil` for every key, not only Escape. The first key press ends the game: returning `nil` means "no next mode".
