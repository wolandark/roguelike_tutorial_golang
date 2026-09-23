# Step 88 · A staircase
## Chapter: Delving deeper

### In this chapter

One floor is not a dungeon. Stairs lead to new floors, kills give experience, and experience buys better stats through a menu you cannot dismiss. The interesting parts are what survives a floor change (only the player) and how one component can serve both monsters and the player with different fields set.

### The problem

Each floor needs a way down: a staircase tile, placed where the player has to explore to find it. The generator already computes the centre of every room it carves, so the centre of the last room is a good spot. The map also has to remember where the stairs are, so that an action can later check whether the player stands on them.

>>> 1. In `tiles.go`, add a package-level variable `downStairs`, a walkable, transparent `Tile` drawn as `>` in both glyphs.
>>> 2. In `gamemap.go`, add the fields `DownstairsX, DownstairsY int` to `GameMap`.
>>> 3. In `procgen.go`, in `GenerateDungeon`, declare `var centerOfLastRoom [2]int` before the loop and set it to the centre of each room that gets a tunnel.
>>> 4. After the loop, set that cell to `downStairs` with `SetTile` and store its position in `DownstairsX` and `DownstairsY`.

!!! Start a new game (`n`) and explore: somewhere in the dungeon a `>` lies on the floor. You can walk over it; nothing happens yet.

--- reveal

{{diff tiles.go}}

{{diff gamemap.go}}

{{diff procgen.go}}

- `[2]int{cx, cy}` is an array literal: two values, fixed size, the same type as the entries of `directions` in step 80.

--- end

%%% Move the two lines that place the stairs to *before* the loop. `centerOfLastRoom` still holds its zero value there, so the stairs land at `(0, 0)`, in the rock of the top-left corner. No room ever reaches that cell, so you never find them. Put the lines back after the loop.
