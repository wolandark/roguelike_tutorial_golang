# Step 36 · Entities live in the map

### The problem

Two things are wrong at once. Monsters are not solid: `MovementAction` checks tiles but not entities. And the entity list lives in the engine, while everything else about a level lives in the map; in chapter 11 levels will be created and thrown away, and the inhabitants must go with them. The fix for the second decides the shape of the first: once the map owns its entities, "is something standing here?" is a question the map can answer, and `Spawn` can put a new entity straight into it.

>>> 1. In `gamemap.go`, add a field `Entities []*Entity` to `GameMap`.
>>> 2. Add a method `func (m *GameMap) GetBlockingEntityAt(x, y int) *Entity` that returns the first blocking entity on that cell, or nil.
>>> 3. Move entity drawing into `GameMap.Render`: after the tiles, draw every visible entity on the lit floor colour.
>>> 4. In `entity.go`, change `Spawn` to `func (e Entity) Spawn(m *GameMap, x, y int) *Entity` and append the copy to `m.Entities`.
>>> 5. In `engine.go`, remove the `Entities` field and the entity loop from `Render`.
>>> 6. In `actions.go`, in `MovementAction.Perform`, also return when `GetBlockingEntityAt` finds something on the destination.
>>> 7. In `procgen.go`, change `GenerateDungeon` to take the map as its first parameter instead of creating it, and drop the width and height parameters.
>>> 8. In `main.go`, create the map first, spawn the player into it, generate, spawn the two monsters into it, and build the engine with `Player` and `GameMap` only.

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
