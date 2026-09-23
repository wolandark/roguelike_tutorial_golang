# Step 79 · Mouse and big steps

### The problem

Moving the cursor one cell per key press is slow across an 80-column map. Two improvements. The mouse should move the cursor, and a left click should select, just like Enter. And holding a modifier should move the cursor further: Shift 5 cells, Ctrl 10, Alt 20, multiplied when several are held.

Handling two event types in one function calls for a **type switch**: `switch e := ev.(type)` runs the case that matches the event's concrete type, and inside each case `e` has that type.

>>> 1. In `SelectIndexHandler.HandleEvent`, replace the key check at the top with `switch e := ev.(type)`, move the existing key handling into a `case *tcell.EventKey:`, and return `h` after the switch.
>>> 2. Add a `case *tcell.EventMouse:` that moves the cursor to `e.Position()` when it is inside the map, and returns `h.OnSelect(x, y)` when `e.Buttons()&tcell.Button1 != 0`.
>>> 3. In the key case, compute a `step` that starts at 1 and is multiplied by 5, 10 and 20 when `e.Modifiers()` contains `tcell.ModShift`, `tcell.ModCtrl` and `tcell.ModAlt`. Move the cursor by `d[0]*step` and `d[1]*step`.

!!! Press `/` and move the mouse: the highlight follows it. Shift-arrow jumps 5 cells. A left click returns to the game, like Enter.

--- reveal

{{diff input.go}}

- `e.Buttons()&tcell.Button1 != 0` tests one bit: `Buttons()` returns a set of bits, one per button, and `&` keeps only the bit for the left button. `Modifiers()` works the same way.
- Some terminals catch Ctrl- or Alt-arrow for themselves, so those may not reach the game. Shift usually works.

--- end

%%% Remove the `InBounds` check in the mouse case and move the mouse over the message log below the map while looking. The highlight follows it off the map, and a click there selects a cell that does not exist. Every coordinate from outside must be checked, as in step 18.
