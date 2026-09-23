# Step 83 · Saving the game
## Chapter: Saving and loading

### In this chapter

Quitting should save, the next run should continue, and dying should delete the save. Then a title screen lets the player choose between a new game and the saved one. Go's `encoding/gob` package turns the whole game into bytes and back in one call each way.

### The problem

Everything about a game is reachable from the engine: the map with its tiles and entities, and the message log. `gob` can write all of that to a file with one `Encode` call. Two details need care.

First, `gob` writes what a pointer points to. `Engine.Player` points at an entity that is also in `GameMap.Entities`. Written twice, it would come back as two separate entities. So we save the map and the log, plus the player's *index* in the entity list.

Second, `AI` and `Consumable` are interfaces. `gob` writes the concrete type's name next to the value, and it must know every such type in advance. `gob.Register` tells it, once, before anything is saved.

>>> 1. Create `saveload.go` with a constant `saveFile = "savegame.sav"` and a struct `saveData` with the fields `GameMap *GameMap`, `MessageLog *MessageLog` and `PlayerIndex int`.
>>> 2. Add a function `func init()` that calls `gob.Register` for `HostileEnemy{}`, `&ConfusedEnemy{}` and the four consumable types.
>>> 3. Add a method `func (e *Engine) SaveAs(path string) error` that finds the player's index in `e.GameMap.Entities`, creates the file with `os.Create`, and encodes a `saveData` into it with `gob.NewEncoder(f).Encode(data)`.
>>> 4. In `input.go`, when the main game gets Escape, call `h.Engine.SaveAs(saveFile)` before returning `nil`. If it fails, log "Failed to save: " and the error in `colorError` and stay in the game.

!!! Play a little and press Escape. A file `savegame.sav` now sits in the folder you ran the game from. Running the game again still starts a new dungeon: nothing loads the file yet.

--- reveal

{{file saveload.go}}

- A function named `init` runs automatically before `main`. A package can have several; you never call them yourself.
- Register a value for types stored as values, and a pointer for `ConfusedEnemy`, which is stored as a pointer.
- `gob` only writes **exported** fields, the ones whose names start with a capital letter. Every field in the game already does.
- `defer f.Close()` right after the error check, the same shape as `defer screen.Fini()` in step 3.

{{diff input.go}}

--- end

%%% Comment out `gob.Register(&ConfusedEnemy{})`, confuse an orc and press Escape. The log shows "Failed to save: gob: type not registered for interface: main.ConfusedEnemy" and you stay in the game. Put the line back.
