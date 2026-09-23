# Step 51 · A hostile AI

### The problem

Monsters need a brain. Different monsters will behave differently, and later a monster's behaviour will be *swapped* while it is confused, so behaviour cannot be a fixed method on `Entity`. It is a component whose *type* varies: an interface with one method, `Perform`, that each kind of behaviour implements.

The hostile brain needs to know whether it can see you. Rather than computing a field of view for every monster, it uses a trick: sight is symmetric, so *if the player can see the monster, the monster can see the player*. When it can see you and is next to you, it attacks; when it can see you but is further away, it steps towards you (in a straight line for now); otherwise it waits.

>>> 1. Create `ai.go` with an interface `type AI interface { Perform(engine *Engine, entity *Entity) }`, and a struct `type HostileEnemy struct{}` with a method `func (HostileEnemy) Perform(engine *Engine, entity *Entity)` that attacks, steps or waits as described above.
>>> 2. In `entity.go`, add a field `AI AI` to `Entity`.
>>> 3. In `entity_factories.go`, give `orc` and `troll` the field `AI: HostileEnemy{}`.
>>> 4. In `engine.go`, rewrite `HandleEnemyTurns` to call `ent.AI.Perform(e, ent)` for every living entity other than the player that has an AI, looping over a copy of the entity list.

!!! Monsters come at you as soon as you see each other and hit back: orcs for 1, trolls for 2. They get stuck on walls. When your HP reaches 0, "You died!"; you can still walk around as a corpse, which is fixed at the end of this chapter.

--- reveal

{{file ai.go}}

- `HostileEnemy` has no state, so it is an empty struct; every orc can share the same value.
- `sign` from `procgen.go`: one package, every file sees every function.

{{diff entity.go}}

- `AI AI` is a field named `AI` of type `AI`; Go allows that.

{{diff entity_factories.go}}

{{diff engine.go}}

- The loop runs over a *copy* of the entity list (an action may add or remove entities while we iterate) and `ent.AI != nil` skips entities without a brain.

--- end

%%% Write a second AI, `Coward`, that steps *away* from the player when it sees them, and give it to the orc template. Nothing outside `ai.go` and the template changes: that is what the interface buys.
