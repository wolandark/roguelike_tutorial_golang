# Step 21 · Actions perform themselves

### The problem

The movement rule now sits inside `Engine.HandleEvent`. Monsters will move too, from a different place, and they must obey the same rule. So the rule belongs to the *action*: an action knows how to `Perform` itself, given the engine and the entity doing it. Then the engine only says "perform this", whoever it is for.

>>> 1. In `actions.go`, give the interface a method: `type Action interface { Perform(engine *Engine, entity *Entity) }`.
>>> 2. Add an empty method `func (EscapeAction) Perform(*Engine, *Entity) {}` so `EscapeAction` still satisfies the interface.
>>> 3. Add a method `func (a MovementAction) Perform(engine *Engine, entity *Entity)` containing the bounds-and-walkable check from `HandleEvent`, moving `entity` instead of the player.
>>> 4. In `engine.go`, replace the type switch in `HandleEvent`: return false for a `nil` action, true for `EscapeAction`, and otherwise call `action.Perform(e, e.Player)`.

!!! Plays exactly like step 18. The rule has moved into the action.

--- reveal

{{diff actions.go}}

- `Action` now requires `Perform`. Any type with that method *is* an `Action`; nothing declares "I implement Action", the compiler checks the method set. `EscapeAction` gets an empty `Perform` so that it still qualifies.
- Bounds first, because `TileAt` on a coordinate outside the slice would panic.

{{diff engine.go}}

- `action.(EscapeAction)` is a type assertion on our own interface, in the same `_, ok` form as before.

--- end

%%% Remove the empty `Perform` on `EscapeAction`. The compiler tells you exactly which method is missing for `EscapeAction` to be used as an `Action`. That message is how you find out what an interface wants.
