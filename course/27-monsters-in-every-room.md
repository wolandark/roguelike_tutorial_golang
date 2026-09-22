# Step 27 · Monsters in every room

### The problem

Two monsters next to the player is a demo, not a dungeon. Each room should get a random handful, mostly orcs, some trolls, on random free cells. The generator already visits every room as it carves, so that is the moment to populate it. Two small decisions: how many per room (0 to a maximum), and what to do when a random cell is already taken (skip it, do not stack).

>>> Add `placeEntities(room, dungeon, maxMonsters)` that rolls a count, picks random floor cells inside the room, skips occupied ones, and spawns an orc 80% of the time, a troll otherwise. Call it from `GenerateDungeon` for every room, with a new constant `maxMonstersPerRoom = 2`.

!!! `o`s and `T`s scattered through the dungeon, appearing as your field of view reaches them.

--- reveal

{{diff procgen.go}}

- `rand.IntN(maxMonsters + 1)` is 0 to `maxMonsters`. `X1 + 1 + rand.IntN(X2-X1-1)` stays off the wall ring. `rand.Float64()` is uniform in `[0, 1)`, so `< 0.8` is an 80% chance.
- The first room is included, so you may wake up next to an orc; the chapter exercise fixes that.

{{diff main.go}}

--- end

%%% Remove the `GetBlockingEntityAt` check in `placeEntities`. Two monsters can now share a cell; you will see one glyph and fight the other. Later, `GetActorAt` returns the first one it finds, and the second is invisible until the first dies.
