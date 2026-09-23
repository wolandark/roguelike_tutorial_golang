# Step 84 · Loading the game

### The problem

On start, the game should continue from the save file when there is one. Loading is saving in reverse: open the file, decode a `saveData`, rebuild the engine. The player is then re-linked from the saved index, so that `Engine.Player` and the entity in the map are the same object again.

If the file does not exist, `os.Open` returns an error, and the game builds a new dungeon as before. And when the player dies, the save must go, or the next run would continue a dead game.

>>> 1. In `saveload.go`, add a function `func LoadGame(path string) (*Engine, error)` that opens the file, decodes it into a `saveData`, builds an `Engine` from its map and log, sets `engine.Player = data.GameMap.Entities[data.PlayerIndex]` and calls `UpdateFOV`.
>>> 2. In `main.go`, start with `engine, err := LoadGame(saveFile)`, and move the code that builds a new game into an `if err != nil { ... }` block. Inside it, assign with `engine = &Engine{...}`, not `:=`.
>>> 3. In `input.go`, in `GameOverEventHandler.HandleEvent`, call `os.Remove(saveFile)` before quitting.

!!! Play a little, press Escape, run again: everything is where you left it, message log included. Die and press Escape: the next run is a fresh dungeon.

--- reveal

{{diff saveload.go}}

- `Decode(&data)` needs a pointer so that it can fill `data`.

{{diff main.go}}

- `engine = &Engine{...}` assigns to the `engine` declared by `engine, err := ...`. With `:=` inside the block, it would declare a second `engine` that exists only inside the `if`, and the outer one would stay `nil`.

{{diff input.go}}

- `os.Remove` returns an error, which we ignore: if there is no file to remove, there is nothing to do.

--- end

%%% Delete the line in `LoadGame` that sets `engine.Player`. Save a game and run again: it panics at once with `invalid memory address or nil pointer dereference`, because `UpdateFOV` reads the player's position and `Player` is `nil`.
