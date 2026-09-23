# Step 55 · Following the path

### The problem

The hostile AI still steps straight at the player. It should take the first step of the path `FindPath` returns instead.

>>> 1. In `ai.go`, in `HostileEnemy.Perform`, replace the straight step with: call `FindPath` from the monster to the player, and if the path is not empty, perform a `MovementAction` towards its first cell.

!!! Monsters round corners and pour through doorways to get to you. Retreat into a corridor and watch them line up.

--- reveal

{{diff ai.go}}

- `path[0]` is the first cell to walk to; the step is the difference between it and the monster's position.

--- end

%%% Set `blockedCost` to 0 in `FindPath`. Monsters now plan routes *through* each other's cells and jam in corridors, each waiting for the one in front to move.
