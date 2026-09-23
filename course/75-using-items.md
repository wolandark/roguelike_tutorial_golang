# Step 75 · Using items

### The problem

Pressing a letter in the menu should use that item. Soon `d` will open the same menu to drop an item instead. The two menus differ only in their title and in what happens to the chosen item. Writing a second handler type would duplicate all the drawing and letter handling.

Go's answer is a field of **function type**. The menu stores a function `OnSelect` and calls it with the chosen item; it does not care what the function does. For `i` that function asks the item's consumable for its action (step 72). It returns an action to perform, or a handler to switch to.

Performing the action from the menu must follow the same rules as a move: an `Impossible` error costs no turn, anything else runs the enemy turns and updates the field of view. That sequence lives inside `MainGameEventHandler.HandleEvent` now. Instead of copying it, move it into a function that both handlers call.

>>> 1. In `input.go`, move the part of `MainGameEventHandler.HandleEvent` from `if err := action.Perform(...)` to the end into a new function `func runAction(engine *Engine, self EventHandler, action Action) EventHandler`. It returns `self` for a `nil` action or an error, and `&MainGameEventHandler{Engine: engine}` after a successful turn. End `HandleEvent` with `return runAction(h.Engine, h, action)`.
>>> 2. Add a field `OnSelect func(engine *Engine, item *Entity) (Action, EventHandler)` to `InventoryHandler` and an `onSelect` parameter to `NewInventoryHandler`. Add a function `useItem` with that signature that returns `item.Consumable.GetAction(engine, engine.Player, item)`.
>>> 3. In `InventoryHandler.HandleEvent`, when the key is a letter from `a` to `z`, turn it into an index. Past the end of the list, log "Invalid entry." in `colorInvalid` and stay. Otherwise call `OnSelect`: if it returns a handler, switch to it, else return `runAction(h.Engine, h, action)`. Any other key still returns `h.Parent`.
>>> 4. In `colors.go`, add `colorInvalid` (yellow). In `MainGameEventHandler.HandleEvent`, open the menu for `i` with the title "Select an item to use" and `useItem`.

!!! Pick up a potion and let a monster hit you. Press `i` then `a`: "You consume the Health Potion, and recover 4 HP!" and the potion is gone. At full health: grey "Your health is already full.", the potion stays and no monster moves. `i` then `z`: "Invalid entry."

--- reveal

{{diff input.go}}

- `runAction` returns a fresh `MainGameEventHandler` after a turn, not `self`. Used from the menu, `self` would be the menu, and the menu should close once the item is used.
- `useItem` is a plain function, and `NewInventoryHandler(..., useItem)` passes it as a value, without calling it: no parentheses.
- `key.Rune() - 'a'` turns a letter back into an index, the reverse of `'a'+i`.

{{diff colors.go}}

--- end

%%% In `runAction`, return `self` instead of `&MainGameEventHandler{Engine: engine}` at the end. Drink a potion from the menu: it works, but the menu stays open afterwards.
