# Step 94 · Deeper is harder
## Chapter: Increasing difficulty

### In this chapter

Floor 1 and floor 9 spawn the same monsters, so going deeper has no bite. One step replaces the fixed chances with two small tables keyed by floor. The Go to notice is that maps iterate in random order, and what to do about it when you need reproducible draws.

### The problem

We want "at most N monsters per room, more as you go down", and "orcs from the start, trolls from floor 3 and more common later, fireballs from floor 6". Two kinds of table express that. A list of (floor, value) pairs where the last applicable entry wins, for the maximums. And a map from floor to weighted entries for what spawns, where entries accumulate as you go down and a later entry for the same entity replaces its weight. Drawing from a weighted table is roulette-wheel selection: roll below the total weight, walk the entries subtracting weights until the roll goes negative.

>>> 1. In `procgen.go`, declare `type floorValue struct{ Floor, Value int }`, the tables `maxItemsByFloor` and `maxMonstersByFloor`, and a function `func maxValueForFloor(values []floorValue, floor int) int`.
>>> 2. Declare a struct `spawnChance` (`Template *Entity`, `Weight int`) and the tables `itemChances` and `enemyChances` of type `map[int][]spawnChance`.
>>> 3. Add a function `func entitiesAtRandom(chances map[int][]spawnChance, n, floor int) []*Entity` that merges the floors up to `floor` and draws `n` templates by weight.
>>> 4. Rewrite `placeEntities` as `func placeEntities(room RectangularRoom, dungeon *GameMap, floor int)` using the tables, and change `GenerateDungeon` to take `currentFloor int` instead of the per-room maximums.
>>> 5. In `gamemap.go`, remove the per-room fields from `GameWorld` and pass `w.CurrentFloor`. Update `setup_game.go` and remove the two constants from `main.go`.

!!! Floor 1 is gentle: only potions and at most two orcs per room. Descend: trolls from floor 3, lightning from floor 4, crowded rooms with fireballs from floor 6.

--- reveal

{{diff procgen.go}}

- Templates are referenced with `&orc`, so the table points at the package-level template; `Spawn` still copies when it places one.
- Go maps iterate in **random order**, so `entitiesAtRandom` keeps a separate `order` slice; without it the same roll could pick different entities on different runs.
- `append(monsters, items...)` joins two slices.

{{diff gamemap.go}}

{{diff setup_game.go}}

{{diff main.go}}

--- end

%%% Change `enemyChances[7]` to give trolls weight 0. On floor 7 and below only orcs spawn: a later floor can *remove* an entity by zeroing its weight, which is how a low-level monster can stop appearing deep down.
