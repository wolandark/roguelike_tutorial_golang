# Step 95 · Weighted spawn tables

### The problem

*What* spawns should depend on the floor too: orcs from the start, trolls from floor 3 and more of them deeper down, confusion scrolls from floor 2, lightning from floor 4, fireballs from floor 6.

Each floor gets a list of (template, weight) entries. Entries accumulate as you go down, and a later entry for the same template replaces its weight. To pick one entry, draw a random number below the total weight, then walk the entries subtracting each weight until the number drops below zero. An entry with weight 30 is picked twice as often as one with weight 15.

The tables are Go maps from floor to entries. A map is the natural fit, but Go iterates maps in a **random order**. The merged weights also live in a map, so the function keeps a separate slice with the templates in the order it first met them, and walks that.

>>> 1. In `procgen.go`, declare a struct `spawnChance` with the fields `Template *Entity` and `Weight int`, and two package-level variables `itemChances` and `enemyChances` of type `map[int][]spawnChance`.
>>> 2. Add a function `func entitiesAtRandom(chances map[int][]spawnChance, n, floor int) []*Entity`. It merges floors 0 to `floor` into a `map[*Entity]int` of weights plus an `order` slice, then draws `n` templates.
>>> 3. In `placeEntities`, replace the two loops with: draw the monsters and the items with `entitiesAtRandom`, then loop once over `append(monsters, items...)` and spawn each template on a free random tile.

!!! Floor 1: only orcs and health potions. Descend: trolls appear from floor 3, lightning scrolls from floor 4, fireballs from floor 6.

--- reveal

{{diff procgen.go}}

- `&orc` points at the package-level template. `Spawn` still copies it when it places one.
- `if _, seen := weights[c.Template]; !seen` uses the two-result map lookup: `seen` is `false` the first time a template appears.
- `append(monsters, items...)` joins two slices into one.

--- end

%%% In `enemyChances`, give trolls weight 0 on floor 7. From floor 7 down only orcs spawn: a later floor can *remove* an entity by setting its weight to zero.
