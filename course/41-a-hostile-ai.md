# Step 41 · A hostile AI

### The problem

Monsters need a brain. Different monsters will behave differently, and in chapter 9 a monster's behaviour will be *swapped* while it is confused, so behaviour cannot be a method on `Entity`. It is a component, but one whose *type* varies: an interface with a single method, `Perform`. The hostile brain needs to know whether it can see you; rather than a second field of view per monster, use a trick: sight is symmetric, so *if the player can see the monster, the monster can see the player*. When adjacent it attacks; when it sees you but is far, it steps towards you (in a straight line for now, which gets it stuck on walls); otherwise it waits.

>>> 1. Create `ai.go` with an interface `type AI interface { Perform(engine *Engine, entity *Entity) }`.
>>> 2. In the same file, declare `type HostileEnemy struct{}` and a method `func (HostileEnemy) Perform(engine *Engine, entity *Entity)`: attack when adjacent to the player, step towards the player when visible, wait otherwise.
>>> 3. In `entity.go`, add a field `AI AI`, a method `func (e *Entity) Distance(x, y int) int` (the larger of the two differences) and a helper `func abs(n int) int`.
>>> 4. In `entity_factories.go`, give `orc` and `troll` `AI: HostileEnemy{}`.
>>> 5. In `engine.go`, rewrite `HandleEnemyTurns` to call `ent.AI.Perform(e, ent)` for every living non-player entity with an AI, looping over a copy of the list.

!!! Monsters come at you as soon as you see each other and hit back: orcs for 1, trolls for 2. They get stuck on walls. When your HP hits 0, "You died!"; you can still walk around as a corpse, which the next chapter fixes.

--- reveal

{{file ai.go}}

- `HostileEnemy` has no state, so it is an empty struct; every orc can share the same value.
- `sign` from `procgen.go`: one package, every file sees every function.

{{diff entity.go}}

- `AI AI` is a field named `AI` of type `AI`; Go allows that. Go has no integer `abs`, hence the helper.

{{diff entity_factories.go}}

{{diff engine.go}}

- The loop runs over a *copy* of the entity list (an action may add or remove entities while we iterate) and `ent.AI != nil` skips entities without a brain.

--- end

%%% Write a second AI, `Coward`, that steps *away* from the player when it sees them, and give it to the orc template. Nothing outside `ai.go` and the template changes: that is what the interface buys.
