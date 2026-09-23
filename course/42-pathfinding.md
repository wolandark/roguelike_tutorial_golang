# Step 42 · Pathfinding

### The problem

A straight step towards you stops at the first wall. A monster needs a route, and the classic tool is **Dijkstra's algorithm**: flood outwards from the monster, always expanding the cheapest cell reached so far, until you reach the player's cell; then walk back along the recorded "came from" links. "Cheapest" needs a priority queue, and Go's standard library has one in `container/heap`, which asks you to supply five small methods on your own slice type. Two tuning choices make monsters behave well: diagonal steps cost a little more than straight ones (paths prefer straight lines), and a cell with another monster on it costs a lot more (monsters flow around each other instead of queueing).

>>> 1. Create `pathfinding.go` with two types for the priority queue: `type pathNode struct{ x, y, cost int }` and `type pathQueue []pathNode`.
>>> 2. Give `pathQueue` the five methods `container/heap` needs: `Len`, `Less`, `Swap`, and `Push` and `Pop` with pointer receivers.
>>> 3. Add a function `func FindPath(m *GameMap, x1, y1, x2, y2 int) [][2]int` implementing Dijkstra over the eight neighbours, with two grids shaped like the map: `cost [][]int` (-1 for unreached) and `prev [][][2]int` (where each cell was reached from). Straight steps cost 2, diagonal 3, cells with a blocking entity +10. Return the cells from the start (exclusive) to the goal.
>>> 4. In `ai.go`, replace the straight step in `HostileEnemy.Perform` with the first cell of `FindPath`.

!!! Monsters round corners and pour through doorways to get to you. Retreat into a corridor and watch them line up.

--- reveal

{{file pathfinding.go}}

- `cost` and `prev` are grids shaped like the map, built the same way: rows first, then each row. `cost[y][x] == -1` means "not reached yet"; `prev[y][x]` is the `[x, y]` of the cell we came from.
- `heap.Pop(pq).(pathNode)` is always the cheapest cell, thanks to our `Less`; the type assertion is needed because the heap works with `any`. Stale queue entries (a cell pushed twice, the second time cheaper) are skipped by `if cur.cost > cost[cur.y][cur.x]`.
- The path is rebuilt backwards: start at the goal, follow `prev` until the start, then reverse in place with a parallel assignment. `p != [2]int{x1, y1}` compares two arrays; Go compares arrays element by element.
- `pathQueue` implements `heap.Interface`: `Len`, `Less`, `Swap`, `Push`, `Pop`. `Push` and `Pop` need pointer receivers because they change the slice's length.

{{diff ai.go}}

--- end

%%% Set `blockedCost` to 0. Monsters now path *through* each other's cells and jam in corridors, each waiting for the one in front. Then set diagonal cost to 2 (same as straight): paths become jagged staircases, because diagonals are never worse than straight steps.
