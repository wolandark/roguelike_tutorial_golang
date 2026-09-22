# Step 48 · The main menu

### The problem

Loading automatically on start is fine for testing and wrong for a game: the player should choose between a new game and continuing, and be told when there is nothing to continue. That is a title screen, one more handler, and a **popup** for the message, another handler that draws its parent dimmed and closes on any key. It is also the moment to move set-up code out of `main.go`, which from now on is only the loop.

>>> Create `setup_game.go` with `NewGame()` (the set-up from `main.go`), a `MainMenu` handler with `n`/`c`/`q` (and Escape), and a `PopupMessage` handler with a parent and a line of text. Continuing with no save (`errors.Is(err, os.ErrNotExist)`) or a bad save shows a popup. `main.go` starts from `MainMenu{}`.

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
