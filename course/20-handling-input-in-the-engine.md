# Step 20 · Handling input in the engine

### The problem

The loop in `run` still does the input half of the turn: poll, turn the key into an action, apply it. That belongs next to the world it changes, so it moves into the engine as a `HandleEvent` method. The loop only has to know when to stop, so the method reports that as a boolean.

>>> 1. In `engine.go`, add a method `func (e *Engine) HandleEvent(ev tcell.Event) (quit bool)`. Return false for non-key events. Type-switch on `handleKey(key)`: return true for `EscapeAction`; for `MovementAction`, do the same bounds-and-walkable check as in step 18, using `e.Player` and `e.GameMap`.
>>> 2. In `main.go`, replace everything after `engine.Render(screen)` in the loop with `if quit := engine.HandleEvent(screen.PollEvent()); quit { return nil }`.

!!! Plays exactly like step 18. The loop in `run` is now two lines.

--- reveal

{{diff engine.go}}

- `(quit bool)` is a **named result**. It documents what the boolean means; the function still returns values explicitly.
- There is no `case nil:` here: a type switch with no matching case simply does nothing, and the function falls through to `return false`.

{{diff main.go}}

--- end

%%% Remove the `if !ok { return false }` lines and resize the window: `key` is `nil`, `handleKey` calls `ev.Key()` on it and the program panics. Non-key events must be filtered before the key is used.
