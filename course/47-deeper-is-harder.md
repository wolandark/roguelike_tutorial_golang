# Step 47 · Deeper is harder
## Chapter: Increasing difficulty

Floor 1 and floor 9 spawn the same monsters, so going deeper has no bite. Two small tables fix that: "from this floor on, at most N per room", and "from this floor on, this entity with this weight".

{{diff procgen.go}}

- `floorValue` lists say "from this floor on, use this value". `maxValueForFloor` walks the list and keeps the last entry whose floor is not deeper than the current one: floors 1 to 3 allow 2 monsters per room, 4 and 5 allow 3, 6 and beyond 5.
- `spawnChance` tables are weighted lists keyed by floor in a `map[int][]spawnChance`. Entries **accumulate** as you go down; listing an entity again on a deeper floor *replaces* its weight, which is how trolls go from rare (15 on floor 3) to common (60 on floor 7). Templates are referenced with `&orc` so the table points at the package-level template; `Spawn` still copies when it places one.
- `entitiesAtRandom` merges every floor up to the current one into `weights`. Go maps iterate in **random order**, so a separate `order` slice keeps the draw reproducible. Then, `n` times, it rolls a number below the total weight and walks the ordered list subtracting weights until the roll goes negative: that entry wins. Roulette-wheel selection, Go's replacement for Python's `random.choices`.
- `placeEntities` shrinks to: pick the counts, draw the templates, spawn each on a free tile. `append(monsters, items...)` joins two slices.

The world no longer needs per-room maximums:

{{diff gamemap.go}}

{{diff setup_game.go}}

{{diff main.go}}

!!! Run it: floor 1 is gentle, only potions and at most two orcs per room. Descend: trolls from floor 3, lightning from floor 4, and crowded rooms with fireballs from floor 6.
