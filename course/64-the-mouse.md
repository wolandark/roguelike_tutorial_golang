# Step 64 · The mouse

### The problem

A `%` on the floor could be an orc or a troll; a `o` might be standing on a potion. Hovering the mouse over a cell and reading what is there is cheap to add: tcell reports mouse events if asked, the engine remembers the last position, and drawing looks up the names on that cell. The same names line will later serve a keyboard "look" command.

>>> 1. In `main.go`, call `screen.EnableMouse()` right after `defer screen.Fini()`.
>>> 2. In `engine.go`, add the fields `MouseX, MouseY int` to `Engine`. In `input.go`, add a function `func trackMouse(engine *Engine, ev tcell.Event)` that stores the position of a `*tcell.EventMouse` when it is inside the map, and call it for every non-key event in `MainGameEventHandler.HandleEvent`.
>>> 3. In `render_functions.go`, add a function `func getNamesAt(m *GameMap, x, y int) string` that joins the names of every entity on a visible cell with `, `, and a function `func renderNamesAtMouse(screen tcell.Screen, x, y int, engine *Engine)` that draws them.
>>> 4. In `engine.go`, call `renderNamesAtMouse(screen, 21, 44, e)` in `Render`.

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
