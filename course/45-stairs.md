# Step 45 · Stairs
## Chapter: Delving deeper

### In this chapter

One floor is not a dungeon. Stairs lead to new levels, kills give experience, and experience buys stat increases through a menu you cannot dismiss. Nothing new in Go; the interesting parts are what survives a floor change (only the player) and how a component can serve both monsters and the player with different fields set.

### The problem

Pressing `>` on a staircase should generate a fresh level one floor down, with the player in it and everything else left behind. Two design questions. Where is the staircase: the last room the generator carved, whose centre we already have. And who owns "how deep are we" and "how to make a floor": a small `GameWorld` with the generator settings and the floor counter, which the engine holds and the save file includes. Because the player must move into the new map's entity list, `Spawn` gains the ability to copy a template without placing it.

>>> Add a `downStairs` tile and `DownstairsX/Y` on the map; place the stairs in the last room. Add `GameWorld` (generator settings, `CurrentFloor`) with `GenerateFloor(engine)` that makes a new map, appends the player, and generates. Add `TakeStairsAction` (`Impossible` off the stairs), bound to `>`. Let `Spawn` accept a `nil` map; `NewGame` spawns the player that way and calls `GenerateFloor`. Save the world, and draw "Dungeon level: N" under the bar.

!!! "Dungeon level: 1" under the bar. Explore until you find the `>`, stand on it, press `>`: "You descend the staircase." and a new level 2. Press `>` elsewhere: "There are no stairs here."

--- reveal

{{diff tiles.go}}

{{diff gamemap.go}}

{{diff procgen.go}}

{{diff actions.go}}

{{diff entity.go}}

{{diff setup_game.go}}

{{diff engine.go}}

{{diff saveload.go}}

{{diff input.go}}

{{diff colors.go}}

--- end

%%% Remove the line that appends the player to the new map's entities in `GenerateFloor`. You descend, the map is generated around the player's position, but the `@` is never drawn and monsters ignore you: you exist in the engine but not in the world. The two lists must agree.
