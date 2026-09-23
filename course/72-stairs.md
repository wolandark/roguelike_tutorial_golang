# Step 72 · Stairs
## Chapter: Delving deeper

### In this chapter

One floor is not a dungeon. Stairs lead to new levels, kills give experience, and experience buys stat increases through a menu you cannot dismiss. Nothing new in Go; the interesting parts are what survives a floor change (only the player) and how a component can serve both monsters and the player with different fields set.

### The problem

Pressing `>` on a staircase should generate a fresh level one floor down, with the player in it and everything else left behind. Two design questions. Where is the staircase: the last room the generator carved, whose centre we already have. And who owns "how deep are we" and "how to make a floor": a small `GameWorld` with the generator settings and the floor counter, which the engine holds and the save file includes. Because the player must move into the new map's entity list, `Spawn` gains the ability to copy a template without placing it.

>>> 1. In `tiles.go`, add a package variable `downStairs`, a walkable tile drawn as `>`.
>>> 2. In `gamemap.go`, add the fields `DownstairsX, DownstairsY int` to `GameMap`, and declare a struct `GameWorld` (the room settings and `CurrentFloor int`) with a method `func (w *GameWorld) GenerateFloor(engine *Engine)` that increments the floor, makes a new map, adds the player to it and generates.
>>> 3. In `procgen.go`, remember the centre of the last room and put `downStairs` there.
>>> 4. In `actions.go`, declare `type TakeStairsAction struct{}` with a `Perform` that is `Impossible` off the stairs and calls `GenerateFloor` on them. Bind it to `>` in `input.go`.
>>> 5. In `entity.go`, let `Spawn` accept a `nil` map. In `setup_game.go`, spawn the player with `Spawn(nil, 0, 0)` and build the world with `GenerateFloor`.
>>> 6. In `engine.go`, add a field `GameWorld *GameWorld` and draw "Dungeon level: N" on row 47. In `saveload.go`, add `GameWorld` to `saveData` and restore it. In `colors.go`, add `colorDescend`.

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
