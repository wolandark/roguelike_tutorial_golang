# Step 24 · Entities live in the map

### The problem

Two things are wrong at once. Monsters are not solid: `MovementAction` checks tiles but not entities. And the entity list lives in the engine, while everything else about a level lives in the map; in chapter 11 levels will be created and thrown away, and the inhabitants must go with them. The fix for the second decides the shape of the first: once the map owns its entities, "is something standing here?" is a question the map can answer, and `Spawn` can put a new entity straight into it.

>>> Move the entity slice from `Engine` into `GameMap`. Give `Spawn` the map as a parameter and append the clone to it. Add `GetBlockingEntityAt(x, y) *Entity` to the map, make `MovementAction` refuse a destination with a blocking entity, and draw entities in `GameMap.Render` (only visible ones, on the lit floor colour). `GenerateDungeon` should receive the map instead of creating it, so the player can be spawned into it first.

!!! Bump into the orc; you stop. The `@` and the monsters are drawn on the yellow lit floor.

--- reveal

{{diff entity.go}}

- `clone` is a local variable whose address escapes into the map; Go allocates it on the heap for you.

{{diff gamemap.go}}

- `GetBlockingEntityAt` is a linear scan. A level has a few dozen entities; that is plenty fast, and a spatial index would be premature.

{{diff engine.go}}

{{diff actions.go}}

{{diff procgen.go}}

{{diff main.go}}

--- end

%%% Set `BlocksMovement: false` on the orc template. You walk through orcs again but not trolls. The flag is per template, which is how ghosts or items will work: same entity type, different flag.
