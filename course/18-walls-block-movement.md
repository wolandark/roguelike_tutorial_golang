# Step 18 · Walls block movement

### The problem

Walk into the wall in step 13: you pass through. The movement code changes x and y without asking the map, because it cannot: the action has no access to it. Who should know about the map? Three options. The key handler could check, but input should not know about walls. The loop in `run` could check, but then every rule ends up in the loop. Or the action itself could perform itself, given what it needs. The third is the roguelike answer: an action knows how to `Perform`, and receives the world to do it in.

"The world" is about to be several things: entities, the player, the map. Bundling them into an **Engine** struct gives actions one parameter, and gives the loop one object to ask for drawing and input handling. After this step `main.go` is tiny and stays that way.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) InBounds(x, y int) bool` that returns true when `0 <= x < m.Width` and `0 <= y < m.Height`.
>>> 2. In `actions.go`, give the interface a method: `type Action interface { Perform(engine *Engine, entity *Entity) }`.
>>> 3. Add an empty method `func (EscapeAction) Perform(*Engine, *Entity) {}` so `EscapeAction` still satisfies it.
>>> 4. Add a method `func (a MovementAction) Perform(engine *Engine, entity *Entity)`: compute the destination, return if it is not `InBounds`, return if `TileAt(...).Walkable` is false, otherwise call `entity.Move(a.DX, a.DY)`.
>>> 5. Create `engine.go` with a struct `Engine` holding `Entities []*Entity`, `Player *Entity` and `GameMap *GameMap`.
>>> 6. Add a method `func (e *Engine) HandleEvent(ev tcell.Event) (quit bool)`: ignore non-key events, turn the key into an action with `handleKey`, return true for `EscapeAction`, otherwise call `action.Perform(e, e.Player)`.
>>> 7. Add a method `func (e *Engine) Render(screen tcell.Screen)` that clears, renders the map, draws every entity and shows.
>>> 8. In `main.go`, build `engine := &Engine{...}` and replace the whole loop body with `engine.Render(screen)` and `if engine.HandleEvent(screen.PollEvent()) { return nil }`.

!!! Walk into the three-cell wall; the `@` stops. The map edge stops you too.

--- reveal

{{diff gamemap.go}}

- `InBounds` is four comparisons: the cell is inside when `x` is at least 0 and below `Width`, and the same for `y` and `Height`. It exists because `TileAt` on a cell outside the map either panics or, for some coordinates, quietly returns a cell from a neighbouring row (see the experiment in step 17). Every caller that computes a coordinate checks it first.

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
