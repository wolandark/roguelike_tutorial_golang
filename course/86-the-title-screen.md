# Step 86 · The title screen

### The problem

The game should open on a title screen with three choices: `n` for a new game, `c` to continue the saved one, `q` to quit. A screen that draws something and reacts to keys is an event handler, so the title screen is one more handler, and the main loop starts with it instead of with the game.

The menu has no state of its own, so it can be an empty struct, used as a value.

>>> 1. In `setup_game.go`, declare `type MainMenu struct{}` and a package-level variable `menuTitle`, a `[]string` with the lines of the title.
>>> 2. Give `MainMenu` an `OnRender` that draws the title lines centred in `colorMenuTitle`, and the three options `[N] Play a new game`, `[C] Continue last game` and `[Q] Quit` below it in `colorMenuText`.
>>> 3. Give it a `HandleEvent` that returns `nil` for Escape, Ctrl-C and `q`. For `n` it returns `&MainGameEventHandler{Engine: NewGame()}`. For `c` it calls `LoadGame(saveFile)` and returns a main game handler on success, or `m` on an error.
>>> 4. In `main.go`, remove the loading code and start the loop with `var handler EventHandler = MainMenu{}`. In `colors.go`, add `colorMenuTitle` (yellow) and `colorMenuText`.

!!! The title screen appears. `n` starts a game, Escape saves and quits. Run again and press `c`: the saved game continues. Delete `savegame.sav` and press `c`: nothing happens at all.

--- reveal

{{diff setup_game.go}}

- `(screenWidth-len([]rune(line)))/2` is the x that centres a line: half of the space that is left over.
- `case 'q', 'Q':` accepts both lower and upper case.
- The methods have value receivers, `func (MainMenu)` and `func (m MainMenu)`. A value receiver without a name is fine when the method does not use it.

{{diff main.go}}

{{diff colors.go}}

--- end

%%% Change `var handler EventHandler = MainMenu{}` to `var handler EventHandler = &MainMenu{}`. It still works: the methods of a value type can be called through a pointer too, so `*MainMenu` also satisfies `EventHandler`.
