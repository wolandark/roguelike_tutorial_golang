# Step 26 · Render order

A monster standing on a corpse should be drawn on top of it. Whatever is drawn last wins, so the entities are sorted before drawing: corpses first, later items, then living actors.

{{diff entity.go}}

- `type RenderOrder int` is a new named type based on `int`, and the `const` block with **`iota`** numbers its values 0, 1, 2. `iota` counts up within a `const` block; only the first line needs the type and value.

{{diff fighter.go}}

{{diff entity_factories.go}}

{{diff gamemap.go}}

- `append([]*Entity(nil), m.Entities...)` makes a *copy* of the slice (`...` spreads a slice into variadic arguments). We sort the copy so the game's own order, which decides who acts first, stays untouched.
- `sort.SliceStable` takes the slice and a **function literal** (an anonymous function, also called a closure) that says whether element `i` should come before element `j`. *Stable* keeps equal elements in their original order.

!!! Run it: kill an orc, then step onto its corpse and let a troll follow you: living things are always drawn over the `%`.
