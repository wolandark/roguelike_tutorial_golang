# Step 94 · Deeper is harder
## Chapter: Increasing difficulty

### In this chapter

Floor 1 and floor 9 spawn the same monsters, so going deeper has no bite. Two steps replace the fixed numbers with small tables keyed by floor: first how *many* things spawn per room, then *which* things. The Go detail to notice is that maps iterate in random order, and what to do when the order matters.

### The problem

Rooms should get more crowded as you go down. We want "at most 2 monsters per room from floor 1, 3 from floor 4, 5 from floor 6", and "at most 1 item per room, 2 from floor 4". Each of those is a short list of (floor, value) pairs, sorted by floor. To look up a floor, walk the list and keep the value of the last entry whose floor is not deeper than the one asked for.

For that, the generator has to know which floor it builds. `GameWorld` has the floor counter, so it passes `CurrentFloor` down instead of the two per-room maximums, which the tables now replace.

>>> 1. In `procgen.go`, declare `type floorValue struct{ Floor, Value int }` and two package-level variables: `maxItemsByFloor = []floorValue{{1, 1}, {4, 2}}` and `maxMonstersByFloor = []floorValue{{1, 2}, {4, 3}, {6, 5}}`.
>>> 2. Add a function `func maxValueForFloor(values []floorValue, floor int) int` that returns the value of the last entry whose `Floor` is not greater than `floor`, and 0 when there is none.
>>> 3. Change `placeEntities` to take `floor int` instead of the two maximums and look them up with `maxValueForFloor`. Change `GenerateDungeon` to take `currentFloor int` instead of the two maximums and pass it on.
>>> 4. In `gamemap.go`, remove `MaxMonstersPerRoom` and `MaxItemsPerRoom` from `GameWorld` and pass `w.CurrentFloor` to `GenerateDungeon`. Remove those two fields from `NewGame` in `setup_game.go` and the two constants from `main.go`.

!!! Floor 1 has fewer potions than before, at most one per room. After a few staircases, rooms hold more monsters.

--- reveal

{{diff procgen.go}}

- `struct{ Floor, Value int }` declares two fields of the same type in one line. `{1, 2}` inside the slice literal is a `floorValue` with its fields in order: the element type can be left out inside a slice literal.
- `break` leaves the loop early: the list is sorted, so once an entry is too deep, all the later ones are too.

{{diff gamemap.go}}

{{diff setup_game.go}}

{{diff main.go}}

- A save made before this step still loads: `gob` skips the fields that no longer exist.

--- end

%%% Put the entries of `maxMonstersByFloor` in the wrong order, `{{6, 5}, {1, 2}, {4, 3}}`. On floor 1 the loop stops at the first entry, which is too deep, and returns 0: no monsters at all on the first floors. The function relies on the list being sorted.
