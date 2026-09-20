# Step 24 · Hit points and death
## Chapter: Doing (and taking) damage

Not every entity fights: a potion has no hit points, a corpse has no strength. Rather than a class hierarchy, an `Entity` gets an optional **component**: a pointer that is `nil` when the entity does not have that ability. This step adds the `Fighter` component and what happens when its HP reaches zero.

Create `fighter.go`:

{{file fighter.go}}

- `SetHP` clamps the new value to `[0, MaxHP]` with the built-in `min` and `max`, then triggers death once. The component does not know which entity it belongs to (no pointer back: that would create reference cycles that make saving in chapter 10 painful), so the entity is passed in.
- `Die` is a makeover, not a removal: the same entity becomes a `%`, dark red, walkable, not alive, "remains of Orc". Corpses stay on the floor.
- `e == engine.Player` compares pointers: is this entity *the* player?

The entity gets the component, an `Alive` flag, and `Spawn` learns to copy the component:

{{diff entity.go}}

- `Fighter *Fighter` is a pointer field. Copying an entity copies the **pointer**, not the struct behind it, so without the extra lines in `Spawn` every orc would share one HP counter with the template. `f := *e.Fighter` reads the struct behind the pointer into a fresh variable and `&f` points at the copy.

Templates get stats:

{{diff entity_factories.go}}

- `Fighter: &Fighter{...}` inside a struct literal: a pointer to a new `Fighter` value.

The engine's monster loop skips corpses:

{{diff engine.go}}

!!! Run it: nothing visible changes yet; kicking still only annoys. Next step, the numbers get used.
