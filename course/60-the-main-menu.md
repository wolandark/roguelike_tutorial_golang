# Step 60 · The main menu

### The problem

Loading automatically on start is fine for testing and wrong for a game: the player should choose between a new game and continuing, and be told when there is nothing to continue. That is a title screen, one more handler, and a **popup** for the message, another handler that draws its parent dimmed and closes on any key. It is also the moment to move set-up code out of `main.go`, which from now on is only the loop.

>>> 1. Create `setup_game.go` and move the new-game code from `main.go` into a function `func NewGame() *Engine`.
>>> 2. In the same file, declare `type MainMenu struct{}` with an `OnRender` (title and three options) and a `HandleEvent` (`n` new game, `c` load, `q` or Escape quit).
>>> 3. Declare a struct `PopupMessage` with fields `Parent EventHandler` and `Text string`, an `OnRender` that dims the parent and writes the text, and a `HandleEvent` that returns `Parent` on any key. Add a helper `func dimScreen(screen tcell.Screen)`.
>>> 4. In `main.go`, start the handler loop from `MainMenu{}`. In `colors.go`, add `colorMenuTitle` and `colorMenuText`.

!!! The title screen. Press `c` before ever saving: the screen dims and "No saved game to load." appears; any key returns to the menu. `n` starts a game; Escape saves; `c` next time continues it.

--- reveal

{{file setup_game.go}}

- `MainMenu` has no state, so it is an empty struct used as a value with value receivers.
- `errors.Is(err, os.ErrNotExist)` asks whether the error *is* "file not found", seeing through wrapping.
- `dimScreen` walks every cell with `screen.Size()` and re-sets it with `style.Dim(true)`.

{{diff main.go}}

{{diff colors.go}}

--- end

%%% Make `PopupMessage.HandleEvent` return `p` for every event. The popup can never be closed and the only way out is Ctrl-C, which the popup also ignores, so you must kill the terminal. Every modal handler needs an exit; check yours before you run it.
