# Step 19 · Entities live in the map

Two things at once, because they belong together. First, the entity list moves from the engine into the map: a level owns its inhabitants, and later levels will come and go. Second, monsters become solid: movement checks for a blocking entity, and `Spawn` puts the new entity straight into the map.

{{diff entity.go}}

- `Spawn` takes the map and appends the clone to `m.Entities`. Because `clone` is a local variable whose address escapes, Go allocates it on the heap; you never think about that.

{{diff gamemap.go}}

- `GetBlockingEntityAt` is a linear scan over the entities. A level has a few dozen; that is plenty fast.
- `Render` draws the entities after the tiles, only when visible, using the *lit* floor colour as background so the glyph sits on its tile.

{{diff engine.go}}

- The engine keeps just the player and the map.

{{diff actions.go}}

- A third check in `MovementAction`: the destination must be free of blocking entities.

The generator no longer creates the map; it receives it, so the player can be spawned into it first:

{{diff procgen.go}}

{{diff main.go}}

!!! Run it: bump into the orc; you stop. The `@` and the monsters are drawn on the yellow lit floor.
