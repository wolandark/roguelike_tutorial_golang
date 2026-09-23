# Step 90 · Taking the stairs

### The problem

Pressing `>` on the staircase should take the player one floor down. That is an action, so it follows the usual rule: off the stairs it is `Impossible` and costs nothing. On the stairs, it asks the world for the next floor. The field of view is computed by `runAction` after every successful action, so the new floor is visible at once.

>>> 1. In `colors.go`, add `colorDescend` (purple).
>>> 2. In `actions.go`, declare `type TakeStairsAction struct{}` with a `Perform` that returns `Impossible{"There are no stairs here."}` unless the entity stands on `DownstairsX, DownstairsY`.
>>> 3. On the stairs, `Perform` calls `engine.GameWorld.GenerateFloor(engine)`, logs "You descend the staircase." and returns `nil`.
>>> 4. In `input.go`, add a `case '>'` to `MainGameEventHandler.HandleEvent` that returns `runAction(h.Engine, h, TakeStairsAction{})`.

!!! Find the `>`, stand on it, press `>`: "You descend the staircase." and "Dungeon level: 2". Press `>` anywhere else: grey "There are no stairs here."

--- reveal

{{diff colors.go}}

{{diff actions.go}}

{{diff input.go}}

- `>` is a character like any other, so it goes in the rune `switch`, not in `handleKey`.

--- end

%%% In `TakeStairsAction.Perform`, remove the position check. `>` now works anywhere: a free escape from every fight. The rule lives in the action, not in the key binding.
