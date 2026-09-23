# Step 46 · Death

### The problem

Hit points going to zero should kill. In a roguelike, death is not a removal but a makeover: the same entity becomes a corpse, a `%` that no longer blocks and no longer acts, and it stays on the floor. The entity needs a flag saying whether it is alive, and something has to watch the hit points and trigger the makeover exactly once, when they reach zero.

One design decision: the `Fighter` does not keep a pointer back to its entity. That would create a cycle, entity to fighter to entity, and cycles make saving the game painful later. Methods that need the entity take it as a parameter instead.

>>> 1. In `entity.go`, add a field `Alive bool`, and in `entity_factories.go` set `Alive: true` on the three templates.
>>> 2. In `fighter.go`, add a method `func (f *Fighter) SetHP(engine *Engine, entity *Entity, hp int)` that stores `hp`, clamped so it is at least 0 and at most `MaxHP`, and calls `entity.Die(engine)` when the result is 0 and the entity is still alive.
>>> 3. Add a method `func (e *Entity) Die(engine *Engine)` that logs "You died!" for the player or "<name> is dead!" for others, and turns the entity into a red `%` called "remains of <name>" that neither blocks nor is alive.
>>> 4. In `engine.go`, skip dead entities in `HandleEnemyTurns`.

!!! Nothing visible changes yet; nothing deals damage. The next step does.

--- reveal

{{diff fighter.go}}

- `max(0, min(hp, f.MaxHP))` clamps with the built-in `min` and `max` functions: the result is never below 0 and never above `MaxHP`.
- `e == engine.Player` compares two pointers: is this entity *the* player?
- `Die` checks nothing itself; `SetHP` makes sure it is called only once, by checking `entity.Alive` first.

{{diff entity.go}}

{{diff entity_factories.go}}

{{diff engine.go}}

--- end

%%% Delete the line `e.Alive = false` from `Die`. After the next step, kill an orc and bump into its corpse: you attack it again, because it still counts as a living fighter, and it dies a second time as "remains of remains of Orc". The `Alive` flag is what makes death happen once.
