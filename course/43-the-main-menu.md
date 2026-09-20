# Step 43 · The main menu

A title screen with new game, continue and quit. It is one more event handler, and a **popup** is another: it draws its parent dimmed, writes one line, and returns the parent on any key. The setup code moves out of `main.go` into a `NewGame` function next to the menu.

Create `setup_game.go`:

{{file setup_game.go}}

- `NewGame` is the setup that was in `main.go`.
- `MainMenu` is a handler with no state, so it is declared as an empty struct and used as a value (`MainMenu{}`), and its methods have value receivers.
- `OnRender` draws a block-letter title (a slice of strings) and three options. `HandleEvent` starts a new game on `n`, loads on `c`, quits on `q` or Escape.
- Loading can fail two ways. `errors.Is(err, os.ErrNotExist)` asks whether the error *is* "file not found" (also seeing through wrapped errors); that gets "No saved game to load.". Any other error gets its own message. Both show a `PopupMessage`.
- `PopupMessage` renders its parent, then `dimScreen` walks every cell with `screen.Size()` and re-sets it with `style.Dim(true)`, then writes one centred line. Any key returns the parent.

{{diff main.go}}

- The loop starts from `MainMenu{}` now, and `main.go` is nothing but the loop. It will not change again.

{{diff colors.go}}

!!! Run it: the title screen. Press `c` before ever saving: the screen dims and "No saved game to load." appears; any key returns to the menu. `n` starts a game; Escape saves and brings you back to the shell; `c` next time continues it.
