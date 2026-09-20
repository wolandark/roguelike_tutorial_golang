# Step 44 · Stairs
## Chapter: Delving deeper

One floor is not a dungeon. The last room of every level gets a `>` staircase; pressing `>` on it generates a fresh level one floor deeper. A `GameWorld` remembers the generator settings and the depth.

{{diff tiles.go}}

{{diff gamemap.go}}

- The map remembers where its stairs are, so the stairs check is a comparison, not a search.
- `GameWorld.GenerateFloor` throws the old map away, makes a new one, puts the player into its entity list and runs the generator, which moves the player into the first room. Everything left on the old floor is gone; there is no going back up, as in the original tutorial.

{{diff procgen.go}}

- The generator remembers the centre of the last room it carved and turns it into stairs.

{{diff actions.go}}

- `TakeStairsAction` is `Impossible` unless the player stands exactly on the stairs.

{{diff entity.go}}

- `Spawn` now accepts a `nil` map: "copy the template, place it later". `NewGame` spawns the player that way and lets `GenerateFloor` do the placing.

{{diff setup_game.go}}

{{diff engine.go}}

- The engine holds the world and prints "Dungeon level: N" under the health bar.

{{diff saveload.go}}

- The save gains the `GameWorld`, otherwise a loaded game would forget its depth.

{{diff input.go}}

{{diff colors.go}}

!!! Run it: "Dungeon level: 1" under the bar. Explore until you find the `>`, stand on it, press `>`: "You descend the staircase." and a brand new level 2. Press `>` elsewhere: "There are no stairs here."
