# Step 38 · Solid monsters

### The problem

You can still walk straight through the orc. Movement checks the tile at the destination but not what stands on it. Now that the map owns its entities, "is something solid standing here?" is a question the map can answer.

>>> 1. In `gamemap.go`, add a method `func (m *GameMap) GetBlockingEntityAt(x, y int) *Entity` that returns the first entity on that cell whose `BlocksMovement` is true, or `nil` if there is none.
>>> 2. In `actions.go`, in `MovementAction.Perform`, return without moving when `GetBlockingEntityAt` finds something on the destination.

!!! Walk into the orc or the troll; you stop.

--- reveal

{{diff gamemap.go}}

- `GetBlockingEntityAt` is a linear scan. A level has a few dozen entities; that is plenty fast, and a spatial index would be premature.

{{diff actions.go}}

--- end

%%% Set `BlocksMovement: false` on the orc template. You walk through orcs again but not trolls. The flag is per template, which is how ghosts or items will work: same entity type, different flag.
