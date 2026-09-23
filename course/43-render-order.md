# Step 43 · Render order

### The problem

Kill an orc, step onto its corpse, and let a troll follow you: sometimes the `%` is drawn over the troll. Entities are drawn in list order, and whatever is drawn last wins. We want corpses at the bottom, items (chapter 8) above them, living actors on top. That is a sort key per entity and a sort before drawing, but the sort must not disturb the list itself, because list order also decides who acts first.

>>> 1. In `entity.go`, declare `type RenderOrder int` with the constants `RenderCorpse`, `RenderItem`, `RenderActor` using `iota`, and add a field `RenderOrder RenderOrder` to `Entity`.
>>> 2. In `entity_factories.go`, set `RenderOrder: RenderActor` on the three templates.
>>> 3. In `fighter.go`, set `e.RenderOrder = RenderCorpse` in `Die`.
>>> 4. In `gamemap.go`, in `Render`, copy `m.Entities`, sort the copy with `sort.SliceStable` by `RenderOrder`, and draw the sorted copy.

!!! Kill an orc, step onto its corpse, let a troll follow you: the living are always drawn over the `%`.

--- reveal

{{diff entity.go}}

- `type RenderOrder int` is a named type based on `int`, and the `const` block with **`iota`** numbers its values 0, 1, 2. `iota` counts up within a `const` block; only the first line needs the type and value.

{{diff fighter.go}}

{{diff entity_factories.go}}

{{diff gamemap.go}}

- `append([]*Entity(nil), m.Entities...)` copies the slice (`...` spreads a slice into variadic arguments).
- `sort.SliceStable` takes the slice and a **function literal** (a closure) that says whether element `i` sorts before element `j`. *Stable* keeps equal elements in their original order, so two orcs on one cell do not flicker.

--- end

%%% Sort `m.Entities` itself instead of a copy. Nothing looks different, but the player is now drawn last *and acts last* among equals; in chapter 8, dropping an item and picking it up again reorders the list under you. Sorting a copy keeps drawing and logic separate.
