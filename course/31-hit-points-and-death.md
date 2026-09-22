# Step 31 · Hit points and death
## Chapter: Doing (and taking) damage

### In this chapter

Combat, in five steps: hit points and death, actual damage, drawing order for corpses, a hostile brain, and pathfinding so monsters can find you around corners. The Go theme is optional behaviour: not every entity fights, not every entity thinks, and Go has no class hierarchy to express "a fighting entity". Components, pointer fields that may be nil, are the answer, and they come with one trap around copying.

### The problem

A potion has no hit points; a corpse has no strength. In a class hierarchy you would subclass `Entity` into `Actor`. Go does not have inheritance, and its alternative is better here anyway: give `Entity` an optional **component**, a pointer to a `Fighter` struct that is `nil` when the entity cannot fight. "Can it fight?" becomes `e.Fighter != nil`. Death is then not a removal but a makeover: the same entity becomes a `%` that no longer blocks or acts.

One subtle decision: the `Fighter` will not keep a pointer back to its entity. That would create a cycle (entity to fighter to entity), and cycles make saving in chapter 10 painful. Methods that need the entity take it as a parameter instead.

>>> Create `fighter.go` with a `Fighter` struct (HP, MaxHP, Defense, Power), a `SetHP(engine, entity, hp)` that clamps to `[0, MaxHP]` and triggers death once, and an `Entity.Die(engine)` that logs, turns the entity into a red `%`, stops it blocking, clears `Alive` and renames it "remains of ...". Add `Alive bool` and `Fighter *Fighter` to `Entity`, make `Spawn` copy the Fighter struct, give the templates stats, and make the engine's monster loop skip dead entities.

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
