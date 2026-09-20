# Step 29 · Event handlers

When the player is dead, only Escape should work. That is a second *mode* of the game, and later there will be more: menus, aiming, a title screen. So input handling moves out of the engine into **event handlers**, one type per mode, and the main loop becomes generic: it asks the current handler to draw, gives it the next event, and replaces it with whatever the handler returns.

`input.go` is rewritten around this idea:

{{diff input.go}}

- `EventHandler` is an interface with two methods. `HandleEvent` processes one event and **returns the handler that should be active next**: itself to stay, another to switch modes, `nil` to quit. `OnRender` draws that mode. This one idea carries the rest of the course.
- `MainGameEventHandler` holds a pointer to the engine and does what `Engine.HandleEvent` did: key to action, escape check, perform, enemy turns, field of view. If the player died it returns a `GameOverEventHandler`.
- `func (h *MainGameEventHandler) OnRender(screen tcell.Screen) { h.Engine.Render(screen) }` is a one-line method; Go allows the whole body on one line.
- `GameOverEventHandler` ignores everything except Escape and Ctrl-C.

The engine loses `HandleEvent`, and `Render` no longer clears or shows (the loop does):

{{diff engine.go}}

And the loop in `main.go`:

{{diff main.go}}

- `var handler EventHandler = &MainGameEventHandler{Engine: engine}` declares an interface variable holding a pointer to a handler. `for handler != nil` runs until some handler returns `nil`. This exact loop runs the finished game.

!!! Run it: plays as before. Let a troll kill you: the `@` becomes a `%`, "You died!" appears, movement keys do nothing, Escape quits.
