# Step 37 · The inventory menu

### The problem

`i` should show what you carry, each item with a letter, and pressing the letter should use it; `d` should show the same list to drop something. Two menus that differ only in title and in what happens to the chosen item. Writing two handler types would duplicate the drawing and the letter handling. Go's alternative to a subclass-per-menu is a field of **function type**: the menu calls `OnSelect(engine, item)` and does not care what it does. The item's own `Consumable.GetAction` decides whether using it is immediate or needs a target.

>>> Add an `InventoryHandler` with `Engine`, `Parent`, `Title` and `OnSelect func(*Engine, *Entity) (Action, EventHandler)`. `OnRender` draws the parent, then a window on the side away from the player listing `(a) Name`. `HandleEvent` maps a letter to an item, logs "Invalid entry." past the end, otherwise switches to the returned handler or runs the action through `runAction`; any non-letter closes. Two functions, `useItem` (asks the consumable) and `dropItem` (returns `DropItem`). Bind `i` and `d`.

!!! Pick up a potion, take some damage, press `i` then `a`: you recover 4 HP and the potion is gone. At full health: grey message, potion kept, no turn. `d` then a letter drops an item where you stand.

--- reveal

{{diff input.go}}

- `'a'+i` turns an index into a letter; `key.Rune() - 'a'` turns it back. `%c` prints a rune.
- `input.go` now formats strings, so `fmt` joins its imports.

--- end

%%% Add a third menu bound to `x` with the title "Select an item to examine" and an `OnSelect` that logs the item's name and returns `nil, nil`. Three lines, no new types: that is what the function field buys.
