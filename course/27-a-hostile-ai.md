# Step 27 · A hostile AI

Monsters act on their own turn. Behaviour is a component too, but an **interface** rather than a struct, so different monsters can behave differently and a monster's behaviour can even be swapped at run time (chapter 9 confuses one).

Create `ai.go`:

{{file ai.go}}

- `AI` is an interface with one method. `HostileEnemy` is an empty struct that implements it; it has no state, so every orc can share the same value.
- A trick: field of view is symmetric, so *if the player can see the monster, the monster can see the player*. `IsVisible(entity.X, entity.Y)` is that check. When adjacent (distance 1) it attacks; when visible but far it steps straight towards you (a crude route; the next step fixes that); otherwise it waits.
- The move uses `sign(dx), sign(dy)` from `procgen.go`: one package, every file sees every function.

{{diff entity.go}}

- `AI AI` is a field named `AI` of type `AI`; Go allows that. `Distance` is the "king's move" distance: the larger of the horizontal and vertical differences. `abs` is a small helper; Go has no integer `abs` in the standard library.

{{diff entity_factories.go}}

{{diff engine.go}}

- `HandleEnemyTurns` now loops over a *copy* of the entity list (an action may add or remove entities while we iterate) and runs each living monster's AI. `ent.AI != nil` skips entities without a brain.

!!! Run it: monsters come at you as soon as you see each other and hit back: orcs for 1, trolls for 2. They get stuck on walls, because they walk in a straight line. When your HP hits 0, "You died!"; you can still walk around as a corpse, which the next chapter fixes.
