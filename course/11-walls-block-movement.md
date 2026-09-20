# Step 11 · Walls block movement

To stop at a wall, the movement action needs the map. So actions grow a method, `Perform`, and the pieces that have been living in `run` (the entities, the player, the map, the event handling, the drawing) move into an **Engine**. After this step `main.go` is tiny and stays that way.

`actions.go` changes the most:

{{diff actions.go}}

- `Action` is no longer an empty interface: it now requires a method `Perform(engine *Engine, entity *Entity)`. Any type with that method *is* an `Action`; nothing declares "I implement Action", the compiler checks the method set. `EscapeAction` gets an empty `Perform` so that it still qualifies (the engine handles it specially and never calls it).
- `MovementAction.Perform` computes the destination and refuses to move if it is off the map or not walkable. Bounds first, because `TileAt` on an out-of-range coordinate would panic.

Create `engine.go`:

{{file engine.go}}

- `HandleEvent` is the update phase from `run`, moved here: key to action, escape check, `Perform`. Its named result `(quit bool)` documents what the boolean means.
- `action.(EscapeAction)` is a type assertion on our own interface, with the same `_, ok` form as before.
- `Render` is the draw phase: clear, map, entities on top, show.

And `main.go` shrinks:

{{diff main.go}}

- The engine is built with a struct literal; the loop becomes render, then hand the next event to the engine, and stop when it says so.

!!! Run it: walk into the three-cell wall; the `@` stops. The map edge stops you too.
