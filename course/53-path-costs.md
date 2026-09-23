# Step 53 · Path costs

### The problem

Dijkstra's algorithm starts at the monster with cost 0 and repeatedly takes the cheapest cell out of the queue, looks at its eight neighbours, and for each one it can walk on works out the cost of reaching it through this cell. If that is cheaper than any way found before, it records the new cost and **where it came from**, and puts the neighbour in the queue. When the player's cell comes out of the queue, the cheapest route is known.

Two tuning choices make monsters move naturally. A diagonal step costs a little more than a straight one (3 against 2), so routes prefer straight lines. A cell with a monster on it costs 10 more, so monsters walk around each other instead of queueing behind one another.

This step fills in the costs and the "came from" records. Turning them into a path is the next step.

>>> 1. In `pathfinding.go`, add a function `func FindPath(m *GameMap, x1, y1, x2, y2 int) [][2]int`. Import `container/heap`.
>>> 2. Inside it, make two grids shaped like the map: `cost [][]int`, every cell set to -1 (not reached), and `prev [][][2]int` for where each cell was reached from. Set the start cell's cost to 0 and create the queue with the start in it.
>>> 3. Loop while the queue is not empty: pop the cheapest node; stop if it is the goal; skip it if a cheaper cost was already recorded; otherwise, for each walkable in-bounds neighbour, compute the step cost (2 straight, 3 diagonal, +10 if blocked) and record the neighbour's cost, `prev` and a queue entry when the new cost is lower.
>>> 4. For now, `return nil` at the end.

!!! Nothing changes on screen; nothing calls `FindPath` yet.

--- reveal

{{diff pathfinding.go}}

- `heap.Pop(pq).(pathNode)` always returns the cheapest node, thanks to `Less`. The type assertion is needed because the heap works with `any`.
- A cell can be pushed twice if a cheaper route to it is found later. The older entry is **stale**; `if cur.cost > cost[cur.y][cur.x]` skips it when it comes out.
- `prev[ny][nx] = [2]int{cur.x, cur.y}` records that the neighbour was reached from the current cell.

--- end

%%% Change the diagonal cost to 2, the same as a straight step. Nothing visible yet, but once monsters follow these paths they zigzag, because a diagonal is now never worse than going straight.
