# Step 20 · Monsters in every room

Instead of two monsters next to the player, every room gets a random handful: 80% orcs, 20% trolls.

{{diff procgen.go}}

- `placeEntities` rolls how many monsters this room gets (`rand.IntN(maxMonsters + 1)` is 0 to `maxMonsters`), then for each one picks a random floor cell inside the room. The `X1 + 1 + rand.IntN(X2-X1-1)` arithmetic stays off the wall ring. If the spot is taken, `continue` skips that monster rather than stacking two.
- `rand.Float64()` is uniform in `[0, 1)`, so `< 0.8` is an 80% chance.
- `GenerateDungeon` takes the maximum per room as a parameter and calls `placeEntities` for every room, including the first one, so you may wake up next to an orc.

{{diff main.go}}

!!! Run it: `o`s and `T`s scattered through the dungeon, appearing as your field of view reaches them.
