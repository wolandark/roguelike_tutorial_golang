# Step 56 · Saving the game
## Chapter: Saving and loading

### In this chapter

Quitting should save, the next run should continue, and dying should delete the save. Go's `encoding/gob` does the serialisation in one call, but two properties of our data need care and both were decided earlier for exactly this reason: no cycles between components and entities, and a player that is one of the map's entities.

### The problem

Everything about a game is reachable from the engine: the map, its tiles and entities, the log. gob can write all of that, including interface values, with two catches. First, gob writes what a pointer points to, so two pointers to the same entity come back as two separate entities; `Engine.Player` is also inside `GameMap.Entities`, and after loading they would be different objects. The fix is to save the player's *index* and re-link. Second, gob must know every concrete type that can hide behind `AI` and `Consumable` before it sees one, so each is registered once at start-up.

>>> 1. Create `saveload.go` with a constant `saveFile = "savegame.sav"` and a struct `saveData` (`GameMap`, `MessageLog`, `PlayerIndex int`).
>>> 2. Add a function `func init()` that calls `gob.Register` for `HostileEnemy{}`, `&ConfusedEnemy{}` and the four consumable types.
>>> 3. Add a method `func (e *Engine) SaveAs(path string) error` that finds the player's index and encodes `saveData` to the file with `gob`.
>>> 4. Add a function `func LoadGame(path string) (*Engine, error)` that decodes the file, rebuilds the engine, re-links `Player` from the index and calls `UpdateFOV`.
>>> 5. In `input.go`, make Escape in the main game call `SaveAs` before quitting, and Escape on game over call `os.Remove(saveFile)`.
>>> 6. In `main.go`, try `LoadGame(saveFile)` first and fall back to building a new game.

!!! Play a little, press Escape, run again: everything is where you left it, log included. Die and press Escape: next run is a fresh dungeon.

--- reveal

{{file saveload.go}}

- `init` runs automatically before `main`. Register a value for types stored as values, a pointer for `ConfusedEnemy`.
- gob only writes **exported** (capitalised) fields, which every field in the game already is.
- `defer f.Close()` right after the error check, the same shape as `Fini` in step 3.

{{diff input.go}}

{{diff main.go}}

- `engine, err := LoadGame(saveFile)` followed by `engine = &Engine{...}` with plain `=`: an assignment to an existing variable, not a declaration.

--- end

%%% Comment out `gob.Register(&ConfusedEnemy{})`, confuse an orc, and quit. Saving fails with `gob: type not registered`, and thanks to the check in the Escape handler you stay in the game with an error message instead of losing it.

%%% Delete the line in `LoadGame` that sets `engine.Player` from the saved index. Loading now crashes on the first key press with a nil pointer, because after decoding, the player is only reachable through that index.
