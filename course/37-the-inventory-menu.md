# Step 37 · The inventory menu

`i` opens a list of what you carry, each item with a letter; pressing the letter uses it. `d` opens the same list to drop something. "Use" and "drop" are the **same handler** with a different title and a different function for what happens to the chosen item.

{{diff input.go}}

- `OnSelect func(engine *Engine, item *Entity) (Action, EventHandler)` is a field of **function type**. `useItem` and `dropItem` are two plain functions with that signature; the main handler passes one or the other when it creates the menu. This is Go's way to avoid one handler type per menu.
- `useItem` asks the consumable via `GetAction`; `dropItem` returns a `DropItem` action.
- `OnRender` draws the window on whichever side of the screen the player is *not* on, so the menu never covers the `@`. `'a'+i` turns an index into a letter; `%c` prints a rune.
- `HandleEvent` maps a letter back to an index (`key.Rune() - 'a'`), complains about letters past the end of the list, and otherwise either switches to the handler the item returned or runs the action through `runAction`. Any non-letter key returns `Parent` and closes the menu.
- `input.go` now formats strings, so `fmt` joins its imports.

!!! Run it: pick up a potion, take some damage, press `i` then `a`: you recover 4 HP and the potion is gone. Try it at full health: grey message, potion kept, no turn. `d` then a letter drops an item where you stand.
