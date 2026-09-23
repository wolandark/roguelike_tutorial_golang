# Step 74 · An inventory menu

### The problem

You can pick potions up but not see them. `i` should open a window that lists what you carry, each item labelled with a letter: `(a) Health Potion`, `(b) ...`. The window is a new mode, so it is a new event handler, built like the history viewer from step 65. It remembers its parent handler, draws the parent first and the window on top, and returns to the parent when you press a key.

The window should not cover the player. It goes on the right half of the screen when the player is in the left part of the map, and on the left otherwise.

>>> 1. In `input.go`, declare a struct `InventoryHandler` with the fields `Engine *Engine`, `Parent EventHandler` and `Title string`, and a constructor `func NewInventoryHandler(engine *Engine, parent EventHandler, title string) *InventoryHandler`.
>>> 2. Give it an `OnRender` that calls `h.Parent.OnRender(screen)`, then clears a rectangle and draws a frame titled `h.Title`, at x 0, or at x 40 when the player's `X` is 30 or less. Inside, draw one line per item, `(a) Name`, or `(Empty)` when there are none.
>>> 3. Give it a `HandleEvent` that returns `h` for events that are not keys and `h.Parent` for any key.
>>> 4. In `MainGameEventHandler.HandleEvent`, turn the `if` for `v` into a `switch` on `key.Rune()` and add a `case 'i'` that returns `NewInventoryHandler(h.Engine, h, "Inventory")`.

!!! Press `i`: an empty window titled "Inventory" with "(Empty)". Close it with any key, pick up two potions, press `i` again: `(a) Health Potion` and `(b) Health Potion`.

--- reveal

{{diff input.go}}

- `'a'+i` adds an index to a rune, which gives the `i`-th letter. `%c` in `fmt.Sprintf` prints a rune as a character. `input.go` now formats strings, so `fmt` joins its imports.
- `max(len(items)+2, 3)` is the height: one line per item plus the two frame lines, and at least one line inside for `(Empty)`.
- `len([]rune(h.Title))` counts characters, not bytes, as in step 66.

--- end

%%% Change the `x = 40` to `x = 0`. Walk to the left edge of the map and press `i`: the window now covers the `@` and the monsters around it. Put it back.
