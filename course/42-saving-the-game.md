# Step 42 · Saving the game
## Chapter: Saving and loading

Escape should save, and the next run should continue where you left off. Go's `encoding/gob` serialises our whole graph of structs, slices and interface values in one call, with two catches that this file handles.

Create `saveload.go`:

{{file saveload.go}}

- **Catch 1: pointer identity.** `Engine.Player` points at an entity that is *also* in `GameMap.Entities`. gob does not know they are the same object: it would write the player twice and read back two separate copies, and the map's player would no longer be the one you control. So `saveData` stores the player's **index** in the entity list, and `LoadGame` re-links `Player` after decoding. This is also why components never point back at their entity: no cycles, nothing duplicated.
- **Catch 2: interfaces.** `Entity.AI` and `Entity.Consumable` are interface fields. To decode them gob must know every concrete type that can appear, so the `init` function registers each one: a value for types we store as values, a pointer for `ConfusedEnemy`. `init` runs automatically before `main`.
- gob only writes **exported** (capitalised) fields, which every field in the game already is.
- `os.Create` and `os.Open` return files that must be closed; `defer f.Close()` right after the error check, the same shape as `Fini` in step 3.

Escape in the main game saves before quitting; Escape on the game-over screen deletes the file instead (dead is dead):

{{diff input.go}}

And for now `main.go` loads the save if there is one:

{{diff main.go}}

- `engine, err := LoadGame(saveFile)` followed by `engine = &Engine{...}` with plain `=`: the variable already exists, so this is an assignment, not a declaration.

!!! Run it: play a little, press Escape, run it again: everything is where you left it, message log included. Die and press Escape: next run is a fresh dungeon.
