# Step 27 · A hostile AI

### The problem

Monsters need a brain. Different monsters will behave differently, and in chapter 9 a monster's behaviour will be *swapped* while it is confused, so behaviour cannot be a method on `Entity`. It is a component, but one whose *type* varies: an interface with a single method, `Perform`. The hostile brain needs to know whether it can see you; rather than a second field of view per monster, use a trick: sight is symmetric, so *if the player can see the monster, the monster can see the player*. When adjacent it attacks; when it sees you but is far, it steps towards you (in a straight line for now, which gets it stuck on walls); otherwise it waits.

>>> Create `ai.go` with an `AI` interface (`Perform(engine, entity)`) and a `HostileEnemy` empty struct implementing it: attack when adjacent, step towards the player when visible, wait otherwise. Add an `AI AI` field to `Entity`, a `Distance(x, y)` method (king's move: the larger of the two differences), give the monster templates `HostileEnemy{}`, and make `HandleEnemyTurns` run each living monster's AI (over a copy of the list).

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
