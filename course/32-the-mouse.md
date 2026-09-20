# Step 32 · The mouse

Hovering over a `%` and reading "remains of Orc" costs almost nothing: ask tcell for mouse events, remember the last position, and look up what is there when drawing.

{{diff render_functions.go}}

- `getNamesAt` collects the names of every entity on a visible tile into a slice and joins them with `strings.Join`, so a monster on a corpse reads "Orc, remains of Troll". `var names []string` declares an empty slice; `append` grows it.

{{diff engine.go}}

- The engine remembers `MouseX, MouseY` and draws the names on row 44, just above the log.

{{diff input.go}}

- `trackMouse` runs for every event that is *not* a key. `ev.(*tcell.EventMouse)` asserts the mouse type; `Position()` returns the cell under the pointer; positions outside the map are ignored.

{{diff main.go}}

- `screen.EnableMouse()` is required, or the terminal never sends mouse events at all.

!!! Run it: move the mouse over monsters and corpses; their names appear above the log. The in-page terminal supports the mouse too.
