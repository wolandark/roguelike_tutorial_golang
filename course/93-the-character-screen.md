# Step 93 · The character screen

### The problem

The player cannot see their level, XP or stats anywhere. `c` should open a small window with all of them. It is the same pattern as the history viewer and the inventory: a handler with a `Parent` that draws the parent first, a framed window on top, and returns to the parent on any key.

>>> 1. In `input.go`, declare a struct `CharacterScreenHandler` with the fields `Engine *Engine` and `Parent EventHandler`.
>>> 2. Give it an `OnRender` that calls `h.Parent.OnRender(screen)` and draws a window titled "Character Information", on the side away from the player, with the level, the XP, the XP for the next level, attack and defense.
>>> 3. Give it a `HandleEvent` that returns `h.Parent` for any key and `h` for other events.
>>> 4. In `MainGameEventHandler.HandleEvent`, add a `case 'c'` that returns `&CharacterScreenHandler{Engine: h.Engine, Parent: h}`.

!!! Press `c`: level 1, XP 0, "XP for next Level: 350", attack 5, defense 2. Any key closes it.

--- reveal

{{diff input.go}}

--- end

%%% Level up with `b` (Strength), then open the sheet: attack is 6. Kill one more orc and look again: the XP counter restarted from what was left over, and the next level needs 500. The formula from step 91 at work.
