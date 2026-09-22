# Step 33 · Pathfinding

### The problem

A straight step towards you stops at the first wall. A monster needs a route, and the classic tool is **Dijkstra's algorithm**: flood outwards from the monster, always expanding the cheapest cell reached so far, until you reach the player's cell; then walk back along the recorded "came from" links. "Cheapest" needs a priority queue, and Go's standard library has one in `container/heap`, which asks you to supply five small methods on your own slice type. Two tuning choices make monsters behave well: diagonal steps cost a little more than straight ones (paths prefer straight lines), and a cell with another monster on it costs a lot more (monsters flow around each other instead of queueing).

>>> Create `pathfinding.go` with `FindPath(m, x1, y1, x2, y2) [][2]int` implementing Dijkstra over the eight neighbours, using `container/heap` on a `pathQueue` of `pathNode{idx, cost}`; straight steps cost 2, diagonal 3, occupied cells +10. Make `HostileEnemy` take the first step of the path instead of the straight step.

!!! Monsters round corners and pour through doorways to get to you. Retreat into a corridor and watch them line up.

--- reveal

{{file pathfinding.go}}

- `cost` and `prev` are one entry per cell, like the map's own slices; `-1` means "not reached yet".
- `heap.Pop(pq).(pathNode)` is always the cheapest cell, thanks to our `Less`; the type assertion is needed because the heap works with `any`. Stale queue entries (a cell pushed twice, the second time cheaper) are skipped by `if cur.cost > cost[cur.idx]`.
- The path is rebuilt backwards from the goal and reversed in place with a parallel assignment.
- `pathQueue` implements `heap.Interface`: `Len`, `Less`, `Swap`, `Push`, `Pop`. `Push` and `Pop` need pointer receivers because they change the slice's length.

{{diff ai.go}}

--- end

%%% Set `blockedCost` to 0. Monsters now path *through* each other's cells and jam in corridors, each waiting for the one in front. Then set diagonal cost to 2 (same as straight): paths become jagged staircases, because diagonals are never worse than straight steps.
