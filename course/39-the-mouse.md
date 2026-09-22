# Step 39 · The mouse

### The problem

A `%` on the floor could be an orc or a troll; a `o` might be standing on a potion. Hovering the mouse over a cell and reading what is there is cheap to add: tcell reports mouse events if asked, the engine remembers the last position, and drawing looks up the names on that cell. The names line also becomes the "look" command in chapter 9.

>>> Call `screen.EnableMouse()` after `Init`. Add `MouseX, MouseY` to the engine and a `trackMouse` that records the position of any `*tcell.EventMouse` inside the map (call it for every non-key event in the main handler). Add `getNamesAt(m, x, y)` that joins the names of every entity on a visible cell with `, `, and draw the result on row 44.

!!! Move the mouse over monsters and corpses; their names appear above the log. The in-page terminal supports the mouse too.

--- reveal

{{diff render_functions.go}}

- `var names []string` declares an empty slice; `strings.Join` joins with a separator.

{{diff engine.go}}

{{diff input.go}}

- `ev.(*tcell.EventMouse)` asserts the mouse type; `Position()` returns the cell under the pointer.

{{diff main.go}}

--- end

%%% Remove the `IsVisible` check in `getNamesAt` and hover over the dark: you can identify monsters you cannot see. Information leaks through UI as easily as through rendering.
