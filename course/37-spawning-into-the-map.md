# Step 37 · Spawning into the map

### The problem

Every entity has to be added to `gameMap.Entities` by hand after it is created, and that list in `main.go` will be forgotten the first time an entity is created somewhere else, such as by the dungeon generator placing monsters. The creator should not have to remember: `Spawn` can take the map and put the new entity into it itself.

That changes the order of things in `main.go`: the map must exist before the player can be spawned into it, so the generator can no longer create the map. It receives one instead.

>>> 1. In `entity.go`, change `Spawn` to `func (e Entity) Spawn(m *GameMap, x, y int) *Entity` and append the new copy to `m.Entities` before returning it.
>>> 2. In `procgen.go`, change `GenerateDungeon` to take the map as its first parameter, `dungeon *GameMap`, instead of creating one, drop its `mapWidth` and `mapHeight` parameters, and remove the `return`.
>>> 3. In `main.go`, create the map with `NewGameMap` first, spawn the player into it, call `GenerateDungeon` with the map, spawn the orc and the troll into it, and delete the line that set `gameMap.Entities`.

!!! Exactly the same game as step 36.

--- reveal

{{diff entity.go}}

- `clone` is a local variable, and its address is stored in the map. Go notices that it outlives the function and keeps it alive; there is nothing to manage by hand.

{{diff procgen.go}}

{{diff main.go}}

- The player is spawned at `0, 0`; `GenerateDungeon` moves it into the first room, as before.

--- end

%%% Spawn the orc *before* calling `GenerateDungeon`, still at `player.X+3`. The player is at `0, 0` at that moment, so the orc lands at `3, 0`, inside the rock of the top row. Order of calls decides where things are.
