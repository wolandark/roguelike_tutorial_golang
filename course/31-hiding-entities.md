# Step 31 · Hiding entities

### The problem

The map hides tiles, but the engine still draws every entity, including ones standing in the dark. An entity should be drawn only if its cell is visible. The map already knows that; the engine just has to ask. To test it we need something to hide, so `main.go` puts a yellow `@` twelve cells to your right for this step only.

>>> 1. In `engine.go`, inside `Render`, draw an entity only if `e.GameMap.IsVisible(ent.X, ent.Y)`.
>>> 2. In `main.go`, add a yellow NPC at `player.X+12, player.Y` to test it.

!!! The yellow `@` is invisible until your field of view reaches it, then it appears and disappears as you move.

--- reveal

{{diff engine.go}}

{{diff main.go}}

--- end

%%% Move the check so that entities are drawn when their cell is *explored* rather than visible. Now the yellow `@` stays on screen once seen, even when you cannot see it: that is how some roguelikes show remembered items but never remembered monsters. Put it back to `IsVisible`.
