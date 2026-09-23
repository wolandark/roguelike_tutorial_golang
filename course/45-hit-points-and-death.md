# Step 45 · Hit points and death
## Chapter: Doing (and taking) damage

### In this chapter

Combat, in five steps: hit points and death, actual damage, drawing order for corpses, a hostile brain, and pathfinding so monsters can find you around corners. The Go theme is optional behaviour: not every entity fights, not every entity thinks, and Go has no class hierarchy to express "a fighting entity". Components, pointer fields that may be nil, are the answer, and they come with one trap around copying.

### The problem

A potion has no hit points; a corpse has no strength. In a class hierarchy you would subclass `Entity` into `Actor`. Go does not have inheritance, and its alternative is better here anyway: give `Entity` an optional **component**, a pointer to a `Fighter` struct that is `nil` when the entity cannot fight. "Can it fight?" becomes `e.Fighter != nil`. Death is then not a removal but a makeover: the same entity becomes a `%` that no longer blocks or acts.

One subtle decision: the `Fighter` will not keep a pointer back to its entity. That would create a cycle (entity to fighter to entity), and cycles make saving in chapter 10 painful. Methods that need the entity take it as a parameter instead.

>>> 1. Create `fighter.go` with a struct `Fighter` holding `HP, MaxHP int`, `Defense int` and `Power int`.
>>> 2. Add a method `func (f *Fighter) SetHP(engine *Engine, entity *Entity, hp int)` that clamps `hp` to `0..MaxHP` and calls `entity.Die(engine)` when it reaches 0 on a living entity.
>>> 3. Add a method `func (e *Entity) Die(engine *Engine)` that logs the death and turns the entity into a red `%` named "remains of ...", non-blocking and not alive.
>>> 4. In `entity.go`, add the fields `Alive bool` and `Fighter *Fighter`, and make `Spawn` copy the `Fighter` struct into a new pointer.
>>> 5. In `entity_factories.go`, give the three templates `Alive: true` and a `Fighter` with their stats.
>>> 6. In `engine.go`, skip dead entities in `HandleEnemyTurns`.

!!! Nothing visible changes yet; kicking still only annoys. The numbers get used next step.

--- reveal

{{file fighter.go}}

- `SetHP` uses the built-in `min` and `max` to clamp. `e == engine.Player` compares pointers: is this *the* player?

{{diff entity.go}}

- `Fighter *Fighter` is a pointer field. Copying an entity copies the **pointer**, not the struct behind it, so without the extra lines in `Spawn` every orc would share one HP counter with the template. `f := *e.Fighter` reads the struct behind the pointer into a fresh variable and `&f` points at the copy.

{{diff entity_factories.go}}

- `Fighter: &Fighter{...}` inside a struct literal: a pointer to a new `Fighter` value.

{{diff engine.go}}

--- end

%%% Remove the three lines that copy the Fighter in `Spawn`, then (after the next step) hit one orc. Every orc on the level loses HP, and so does the template. This is the copy trap; you have now seen it once and will recognise it forever.
