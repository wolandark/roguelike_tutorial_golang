# Step 87 · A popup

### The problem

Pressing `c` with no save file does nothing, and the player cannot tell why. The menu should say "No saved game to load." For any other loading error, a damaged file for example, it should show that error.

A message that covers the screen until a key is pressed is a small modal handler: a **popup**. It draws its parent dimmed, writes its text on top, and returns to the parent on any key. To tell "no file" apart from other errors, `errors.Is(err, os.ErrNotExist)` checks whether the error means "file does not exist".

>>> 1. In `setup_game.go`, add a function `func dimScreen(screen tcell.Screen)` that re-sets every cell of the screen with `style.Dim(true)`.
>>> 2. Declare a struct `PopupMessage` with the fields `Parent EventHandler` and `Text string`. Its `OnRender` draws the parent, dims the screen and draws the text centred. Its `HandleEvent` returns `p.Parent` for any key and `p` otherwise.
>>> 3. In `MainMenu.HandleEvent`, when `LoadGame` fails with `os.ErrNotExist`, return a `PopupMessage` with "No saved game to load.". For any other error, return one with "Failed to load save: " and the error.

!!! Delete `savegame.sav` and press `c` on the title screen: the menu dims and "No saved game to load." appears. Any key returns to the menu.

--- reveal

{{diff setup_game.go}}

- `errors.Is(err, os.ErrNotExist)` compares against a known error value, also through wrapped errors. `os.Open` wraps it in an error that also names the file, so `err == os.ErrNotExist` would be false.
- `dimScreen` reads each cell with `GetContent`, as in step 78, and writes it back dimmed.

--- end

%%% Make `PopupMessage.HandleEvent` return `p` for every event. The popup can never be closed, and Ctrl-C does not help either, because the popup ignores it. Close the terminal window to get out. Every modal handler needs a way out.
