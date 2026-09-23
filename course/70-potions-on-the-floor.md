# Step 70 · Potions on the floor

### The problem

An item is just an entity lying on the floor: a glyph, a colour, a name, no `Fighter`, not blocking. So a health potion starts as one more template, spawned by the dungeon generator next to the monsters.

Two small problems come with it. The generator picks a random cell inside a room twice now, once for monsters and once for items, so that calculation becomes a method on the room. And the generator must not put a potion under a monster or two potions on one cell. `GetBlockingEntityAt` only sees blocking entities, which potions are not, so the map needs a lookup for *any* entity on a cell.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) EntityAt(x, y int) *Entity` that returns the first entity at `x, y`, or `nil`.
>>> 2. In `entity_factories.go`, add a `healthPotion` template: `'!'`, purple, named "Health Potion", with `RenderOrder: RenderItem`.
>>> 3. In `procgen.go`, add a method `func (r RectangularRoom) randomTile() (int, int)` that returns a random cell inside the room, the same calculation `placeEntities` does now. Add a `maxItems` parameter to `placeEntities`: use `randomTile` and `EntityAt` in the monster loop, and add a second loop that spawns up to `maxItems` potions the same way.
>>> 4. Add a `maxItemsPerRoom` parameter to `GenerateDungeon` and pass it on to `placeEntities`. In `main.go`, add the constant `maxItemsPerRoom = 2` and pass it to `GenerateDungeon`.

!!! Purple `!` in some rooms. You can walk over them, and the `@` is drawn on top. Nothing else happens yet.

--- reveal

{{diff gamemap.go}}

{{diff entity_factories.go}}

- The template has no `Fighter` and no `AI`, so `HandleEnemyTurns` skips it and `GetActorAt` never returns it. `BlocksMovement` is left at its zero value, `false`.

{{diff procgen.go}}

- `randomTile` returns two values, and `x, y := room.randomTile()` receives both.

{{diff main.go}}

--- end

%%% In the item loop in `procgen.go`, remove the `EntityAt` check. Run the game a few times and look for a monster standing on a potion: the monster is drawn and the potion under it is invisible. Render order hides the overlap; only the check prevents it.
