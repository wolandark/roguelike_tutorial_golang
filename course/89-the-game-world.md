# Step 89 · The game world

### The problem

Going down means building a new map. Someone has to know how: the room settings, and how deep the player is. Those are constants in `main.go` now, passed to `GenerateDungeon` once. A small struct, `GameWorld`, will hold the settings and a floor counter, and a method `GenerateFloor` builds the next floor. The engine keeps the world, and the save file includes it.

`GenerateFloor` makes a new map and puts the player into it. Everything else is left behind. That means the player must exist before any map does, so `Spawn` has to accept a `nil` map: copy the template, but do not add it to a list.

>>> 1. In `gamemap.go`, declare a struct `GameWorld` with the fields `MaxRooms`, `RoomMinSize`, `RoomMaxSize`, `MaxMonstersPerRoom`, `MaxItemsPerRoom` and `CurrentFloor`, all `int`. Add a method `func (w *GameWorld) GenerateFloor(engine *Engine)` that increments `CurrentFloor`, sets `engine.GameMap` to a new map, appends the player to its entities, and calls `GenerateDungeon` with the world's settings.
>>> 2. In `entity.go`, make `Spawn` append the clone to `m.Entities` only when `m != nil`.
>>> 3. In `setup_game.go`, in `NewGame`, spawn the player with `playerTemplate.Spawn(nil, 0, 0)`, set `engine.GameWorld` to a `GameWorld` built from the constants, and call `engine.GameWorld.GenerateFloor(engine)`.
>>> 4. In `engine.go`, add a field `GameWorld *GameWorld` and draw "Dungeon level: N" at row 47. In `saveload.go`, add `GameWorld *GameWorld` to `saveData`, save it, and set it again in `LoadGame`.

!!! Delete `savegame.sav` first: a save from before this step has no world, and loading it would crash. Start a new game: "Dungeon level: 1" appears under the health bar.

--- reveal

{{diff gamemap.go}}

- `GenerateFloor` has a pointer receiver because it changes `CurrentFloor`.

{{diff entity.go}}

{{diff setup_game.go}}

{{diff engine.go}}

{{diff saveload.go}}

--- end

%%% In `GenerateFloor`, remove the line that appends the player to the new map's entities. The map is generated around the player's position, but the `@` is never drawn and monsters ignore you: you exist in the engine but not in the map. The two must agree.
