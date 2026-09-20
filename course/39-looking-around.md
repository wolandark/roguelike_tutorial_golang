# Step 39 · Looking around

The next two scrolls need the player to pick a target. That needs a cursor mode: move a highlight over the map with the keyboard or the mouse, confirm with Enter or a click. Built once, the same mode gives us a `/` **look** command for free, so that is what this step adds; the scrolls come next.

{{diff input.go}}

- `SelectIndexHandler` holds the engine and an `OnSelect` callback: a **function value** that receives the chosen tile and returns the next handler. What happens with the tile is entirely up to whoever created the handler.
- The cursor *is* `Engine.MouseX/MouseY`; `NewSelectIndexHandler` starts it on the player.
- `OnRender` draws the game, then reads the cell under the cursor back with `screen.GetContent` and rewrites it with `style.Reverse(true)`: a highlight without knowing what is there.
- `HandleEvent` uses a type switch on the event. Mouse movement moves the cursor; a click (`e.Buttons()&tcell.Button1 != 0`, a bit test) selects. Movement keys move it, and holding Shift, Ctrl or Alt multiplies the step by 5, 10 or 20 (`e.Modifiers()` is a bit set too). `max` and `min` clamp the cursor to the map. Enter selects; any other key cancels back to the main game.
- `NewLookHandler` is the first use: its callback simply returns to the game. The names line from step 32 does the rest.

!!! Run it: press `/`, move the highlight with the arrows or `hjkl` (try Shift-arrow), and read the names line as it passes over monsters. Enter or Escape returns to the game.
