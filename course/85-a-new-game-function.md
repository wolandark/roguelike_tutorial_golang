# Step 85 · A new-game function

### The problem

Loading automatically on start is handy for testing, but a real game asks: new game or continue? That will be a title screen, and its "new game" option needs the code that builds a fresh game. Right now that code sits in the middle of `run` in `main.go`. Before anything else can call it, it has to become a function.

>>> 1. Create `setup_game.go` with a function `func NewGame() *Engine`. Move the new-game code from the `if err != nil` block in `main.go` into it, building the engine first and filling `GameMap` and `Player` on it, and return the engine.
>>> 2. In `main.go`, replace the block's contents with `engine = NewGame()`.

!!! Nothing changes: with a save file the game continues, without one it starts fresh. This is a **refactoring**, a change in structure that keeps the behaviour the same.

--- reveal

{{file setup_game.go}}

{{diff main.go}}

- `setup_game.go` needs no imports: every name it uses is declared in the package.

--- end

%%% In `NewGame`, delete the `engine.UpdateFOV()` line and start a new game (delete `savegame.sav` first). The screen is black apart from the log and the health bar until you take the first step. The field of view is only computed after an action, so a fresh game must compute it once itself.
