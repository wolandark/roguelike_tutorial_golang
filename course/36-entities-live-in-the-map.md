# Step 36 · Entities live in the map

### The problem

The engine keeps the list of entities, while everything else about a level, its tiles, what is visible, what is explored, lives in the map. Later the game will create new levels and throw old ones away, and the monsters of a level must go with it. So the entity list moves into the map, and the map draws its entities along with its tiles.

>>> 1. In `gamemap.go`, add a field `Entities []*Entity` to `GameMap`.
>>> 2. At the end of `GameMap.Render`, after the tiles, draw every entity whose cell is visible, using the lit floor colour of its cell as the background.
>>> 3. In `engine.go`, remove the `Entities` field from `Engine` and the entity loop from `Render`.
>>> 4. In `main.go`, set `gameMap.Entities = []*Entity{monster, bigMonster, player}` and build the engine with only `Player` and `GameMap`.

!!! The same as before, except that the `@`, the `o` and the `T` now sit on the yellow lit floor instead of on a black background.

--- reveal

{{diff gamemap.go}}

- `m.TileAt(e.X, e.Y).Light.BG` is the background colour of the lit tile under the entity. Using it as the entity's background makes the glyph look like it stands on the floor.

{{diff engine.go}}

{{diff main.go}}

--- end

%%% Leave out `.Background(bg)` in `GameMap.Render`. Each entity is drawn on a black square, because a style without a background uses the terminal's default. The previous steps looked like that all along.
