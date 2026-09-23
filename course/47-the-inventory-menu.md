# Step 47 · The inventory menu

### The problem

`i` should show what you carry, each item with a letter, and pressing the letter should use it; `d` should show the same list to drop something. Two menus that differ only in title and in what happens to the chosen item. Writing two handler types would duplicate the drawing and the letter handling. Go's alternative to a subclass-per-menu is a field of **function type**: the menu calls `OnSelect(engine, item)` and does not care what it does. The item's own `Consumable.GetAction` decides whether using it is immediate or needs a target.

>>> 1. In `input.go`, declare a struct `InventoryHandler` with fields `Engine *Engine`, `Parent EventHandler`, `Title string` and `OnSelect func(engine *Engine, item *Entity) (Action, EventHandler)`, and a constructor `NewInventoryHandler`.
>>> 2. Add two plain functions with that signature: `useItem` (asks the item's consumable for its action) and `dropItem` (returns `DropItem{Item: item}`).
>>> 3. Give `InventoryHandler` an `OnRender` that draws the parent and a framed list `(a) Name`, `(b) ...` on the side away from the player.
>>> 4. Give it a `HandleEvent` that maps a letter to an item and calls `OnSelect`, logs "Invalid entry." past the end, and returns `Parent` for any other key.
>>> 5. In `MainGameEventHandler.HandleEvent`, open it for `i` with `useItem` and for `d` with `dropItem`.

!!! Pick up a potion, take some damage, press `i` then `a`: you recover 4 HP and the potion is gone. At full health: grey message, potion kept, no turn. `d` then a letter drops an item where you stand.

--- reveal

{{diff input.go}}

- `'a'+i` turns an index into a letter; `key.Rune() - 'a'` turns it back. `%c` prints a rune.
- `input.go` now formats strings, so `fmt` joins its imports.

--- end

%%% Add a third menu bound to `x` with the title "Select an item to examine" and an `OnSelect` that logs the item's name and returns `nil, nil`. Three lines, no new types: that is what the function field buys.
