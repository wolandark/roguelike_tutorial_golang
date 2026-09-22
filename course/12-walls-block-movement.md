# Step 12 · Walls block movement

### The problem

Walk into the wall in step 11: you pass through. The movement code changes x and y without asking the map, because it cannot: the action has no access to it. Who should know about the map? Three options. The key handler could check, but input should not know about walls. The loop in `run` could check, but then every rule ends up in the loop. Or the action itself could perform itself, given what it needs. The third is the roguelike answer: an action knows how to `Perform`, and receives the world to do it in.

"The world" is about to be several things: entities, the player, the map. Bundling them into an **Engine** struct gives actions one parameter, and gives the loop one object to ask for drawing and input handling. After this step `main.go` is tiny and stays that way.

>>> Give `Action` a method, `Perform(engine *Engine, entity *Entity)`. Make `MovementAction.Perform` refuse to move off the map or onto a non-walkable tile. Create `engine.go` with an `Engine` holding the entities, the player and the map, a `HandleEvent(ev) (quit bool)` that turns keys into actions and performs them, and a `Render(screen)` that clears, draws map and entities, and shows. Shrink `run` to build the engine and loop.

!!! Walk into the three-cell wall; the `@` stops. The map edge stops you too.

--- reveal

{{diff actions.go}}

- `Action` now requires `Perform`. Any type with that method *is* an `Action`; nothing declares "I implement Action", the compiler checks the method set. `EscapeAction` gets an empty `Perform` so that it still qualifies.
- Bounds first, because `TileAt` on a coordinate outside the slice would panic.

{{file engine.go}}

- `HandleEvent` is the update phase from `run`: key to action, escape check, `Perform`. Its named result `(quit bool)` documents what the boolean means.
- `action.(EscapeAction)` is a type assertion on our own interface, in the same `_, ok` form as before.

{{diff main.go}}

--- end

%%% Swap the two checks in `MovementAction.Perform` so `Walkable` is tested first, then walk off the left edge: panic, index out of range. Order of checks is part of correctness, not style.

%%% Remove the empty `Perform` on `EscapeAction`. The compiler tells you exactly which method is missing for `EscapeAction` to be used as an `Action`. That message is how you find out what an interface wants.
