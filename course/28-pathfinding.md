# Step 28 · Pathfinding

A monster that wants to reach you needs a route around walls. **Dijkstra's algorithm** finds the cheapest one: flood outwards from the monster, always expanding the cheapest tile reached so far, until the player's tile is reached. Go's standard library has the priority queue this needs in `container/heap`; we only supply the five small methods it asks for.

Create `pathfinding.go`:

{{file pathfinding.go}}

- `cost` and `prev` are one entry per tile, like the map's own slices. `-1` means "not reached yet".
- `pq := &pathQueue{{idx: start, cost: 0}}` is a pointer to a slice with one node in it. `heap.Push` and `heap.Pop` keep it ordered by cost through our `Less` method, so `heap.Pop(pq).(pathNode)` is always the cheapest tile; the type assertion is needed because the heap works with `any`.
- The nested loops visit the eight neighbours. Straight steps cost 2, diagonals 3 (so paths prefer straight lines), and a tile with a blocking entity costs 10 extra, so monsters flow *around* each other instead of queueing in a corridor.
- `if cur.cost > cost[cur.idx] { continue }` drops stale queue entries: a tile can be pushed twice if a cheaper route is found later.
- After the loop, walking the `prev` links from the goal back to the start gives the path in reverse; the final loop reverses it in place with a parallel assignment `path[i], path[j] = path[j], path[i]`.
- `pathQueue` implements `heap.Interface`: `Len`, `Less`, `Swap` (from `sort.Interface`) plus `Push` and `Pop`. `Push` and `Pop` need pointer receivers because they change the slice's length. `x.(pathNode)` converts the `any` back to our type.

The AI follows the path:

{{diff ai.go}}

- `path[0]` is the first step; the movement is the difference between it and the monster's position.

!!! Run it: monsters round corners and pour through doorways to get to you. Retreat into a corridor and watch them line up.
